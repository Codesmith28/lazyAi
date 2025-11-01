package clipboard

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	"github.com/atotto/clipboard"
)

type Clipboard struct {
	Prompt              chan string
	LastText            string
	Mu                  sync.RWMutex
	OutputText          string
	ignoreNextChange    bool
	ignoreChangeTimeout *time.Timer
	outputHashes        map[string]bool // Track hashes of output we've written
}

func NewClipboard() *Clipboard {
	return &Clipboard{
		Prompt:              make(chan string),
		LastText:            "",
		Mu:                  sync.RWMutex{},
		ignoreNextChange:    false,
		ignoreChangeTimeout: nil,
		outputHashes:        make(map[string]bool),
	}
}

func (c *Clipboard) StartMonitoring() {
	if os.Getenv("WAYLAND_DISPLAY") != "" {
		log.Println("Starting Wayland clipboard monitoring")

		if err := c.MonitorWaylandLinux(); err != nil {
			log.Println("Error starting Wayland clipboard monitoring:", err)
			log.Println("Falling back to polling method")
			go c.PollingClipboard()
			return
		}
		return
	}

	// Non-Wayland fallback
	log.Println("Starting polling clipboard monitoring (non-Wayland)")
	go c.PollingClipboard()
}

func (c *Clipboard) MonitorWaylandLinux() error {
	desktop := strings.ToLower(os.Getenv("XDG_CURRENT_DESKTOP"))
	log.Printf("Detected desktop environment: %s\n", desktop)

	switch {
	case strings.Contains(desktop, "kde"):
		return c.monitorKDEKlipper()
	default:
		return c.monitorPortalClipboard()
	}
}

/* =============================
   KDE (Klipper) Clipboard Monitor
   ============================= */

func (c *Clipboard) monitorKDEKlipper() error {
	log.Println("Using KDE Klipper clipboard monitoring")

	// Monitor for clipboard changes via dbus
	cmd := exec.Command("dbus-monitor", "type='signal',interface='org.kde.klipper.klipper'")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start KDE clipboard monitor: %w", err)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			// When we detect a clipboard history update, fetch the complete clipboard content
			if strings.Contains(line, "clipboardHistoryUpdated") {
				// Use wl-paste to get the complete clipboard content at once
				text, err := exec.Command("qdbus", "org.kde.klipper", "/klipper", "org.kde.klipper.klipper.getClipboardContents").Output()
				if err != nil {
					continue
				}
				// Process the complete clipboard content as a single unit
				c.processClipboardText(strings.TrimSpace(string(text)))
			}
		}
		if err := scanner.Err(); err != nil {
			log.Println("KDE Klipper monitor scanner error:", err)
		}
	}()

	log.Println("KDE clipboard monitoring started via Klipper")
	return nil
}

/* =============================
   GNOME / Portal Clipboard Monitor
   ============================= */

func (c *Clipboard) monitorPortalClipboard() error {
	log.Println("Using GNOME/Freedesktop clipboard monitoring")

	introspect := exec.Command("gdbus", "introspect", "--session",
		"--dest", "org.freedesktop.portal.Desktop",
		"--object-path", "/org/freedesktop/portal/desktop")

	if err := introspect.Run(); err != nil {
		return fmt.Errorf("no portal interface detected: %w", err)
	}

	cmd := exec.Command("dbus-monitor", "--session",
		"destination='org.freedesktop.portal.Desktop',interface='org.freedesktop.portal.Clipboard'")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to get stdout pipe: %w", err)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start portal clipboard monitor: %w", err)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "SelectionOwnerChanged") {
				text, err := exec.Command("wl-paste").Output()
				if err != nil {
					continue
				}
				c.processClipboardText(strings.TrimSpace(string(text)))
			}
		}
		if err := scanner.Err(); err != nil {
			log.Println("Portal monitor scanner error:", err)
		}
	}()

	log.Println("Wayland clipboard monitoring started via Freedesktop portal")
	return nil
}

/* =============================
   Shared Helpers
   ============================= */

func (c *Clipboard) readClipboardStream(stdout io.Reader, source string) {
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		c.processClipboardText(text)
	}
	if err := scanner.Err(); err != nil {
		log.Printf("[%s] stream read error: %v\n", source, err)
	}
}

func (c *Clipboard) processClipboardText(text string) {
	if text == "" {
		return
	}

	c.Mu.Lock()
	defer c.Mu.Unlock()

	// Compute hash of incoming clipboard text
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(text)))

	// Check if this text is one of our outputs
	if c.outputHashes[hash] {
		log.Printf("Ignoring clipboard change - detected our own output (hash: %s)", hash[:16])
		c.LastText = text
		return
	}

	// If we're ignoring the next change (because we just wrote output), skip it
	if c.ignoreNextChange {
		log.Println("Ignoring clipboard change (output feedback prevention)")
		c.ignoreNextChange = false
		if c.ignoreChangeTimeout != nil {
			c.ignoreChangeTimeout.Stop()
			c.ignoreChangeTimeout = nil
		}
		// Update LastText to prevent re-processing if the same text appears later
		c.LastText = text
		return
	}

	if text != c.LastText && text != c.OutputText {
		c.LastText = text
		c.Prompt <- text
	}
}

/* =============================
   Fallback Polling
   ============================= */

func (c *Clipboard) PollingClipboard() {
	log.Println("Starting polling clipboard monitoring")

	for {
		c.Mu.Lock()
		text, err := clipboard.ReadAll()
		if err != nil || strings.TrimSpace(text) == "" {
			c.Mu.Unlock()
			time.Sleep(250 * time.Millisecond)
			continue
		}

		// Compute hash of incoming clipboard text
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte(text)))

		// Check if this text is one of our outputs
		if c.outputHashes[hash] {
			log.Printf("Ignoring clipboard change - detected our own output (hash: %s)", hash[:16])
			c.LastText = text
			c.Mu.Unlock()
			time.Sleep(250 * time.Millisecond)
			continue
		}

		// If we're ignoring the next change (because we just wrote output), skip it
		if c.ignoreNextChange {
			log.Println("Ignoring clipboard change (output feedback prevention)")
			c.ignoreNextChange = false
			if c.ignoreChangeTimeout != nil {
				c.ignoreChangeTimeout.Stop()
				c.ignoreChangeTimeout = nil
			}
			// Update LastText to prevent re-processing if the same text appears later
			c.LastText = text
			c.Mu.Unlock()
			time.Sleep(250 * time.Millisecond)
			continue
		}

		if !isLikelyScreenshot(text) && text != c.LastText && text != c.OutputText {
			c.LastText = text
			c.Prompt <- text
		}

		c.Mu.Unlock()
		time.Sleep(250 * time.Millisecond)
	}
}

/* =============================
   Public API
   ============================= */

func (c *Clipboard) GetClipboardText() (string, error) {
	return <-c.Prompt, nil
}

func (c *Clipboard) SetClipboardText(text string) error {
	c.Mu.Lock()

	// Compute hash of the output we're writing
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
	c.outputHashes[hash] = true
	log.Printf("Recording output hash: %s (total tracked: %d)", hash[:16], len(c.outputHashes))

	// Set the flag to ignore the next clipboard change
	c.ignoreNextChange = true

	// Clear any existing timeout
	if c.ignoreChangeTimeout != nil {
		c.ignoreChangeTimeout.Stop()
	}

	// Set a timeout to clear the ignore flag (safety mechanism)
	// This prevents the flag from staying set indefinitely if something goes wrong
	c.ignoreChangeTimeout = time.AfterFunc(2*time.Second, func() {
		c.Mu.Lock()
		defer c.Mu.Unlock()
		if c.ignoreNextChange {
			log.Println("Ignore flag timeout - clearing flag")
			c.ignoreNextChange = false
		}
	})

	c.Mu.Unlock()

	return clipboard.WriteAll(text)
}

func Clear() error {
	return clipboard.WriteAll(" ")
}

/* =============================
   Screenshot Heuristic
   ============================= */

func isLikelyScreenshot(data string) bool {
	signatures := [][]byte{
		{0xFF, 0xD8, 0xFF},       // JPEG
		{0x89, 0x50, 0x4E, 0x47}, // PNG
		{0x47, 0x49, 0x46, 0x38}, // GIF
		{0x42, 0x4D},             // BMP
		{0x00, 0x00, 0x01, 0x00}, // ICO
	}

	dataBytes := []byte(data)
	if len(dataBytes) < 4 {
		return false
	}

	for _, sig := range signatures {
		if bytes.HasPrefix(dataBytes, sig) {
			return true
		}
	}

	if len(dataBytes) >= 8 {
		format := binary.LittleEndian.Uint32(dataBytes[:4])
		if format == 2 || format == 8 || format == 17 {
			return true
		}
	}

	return false
}
