package clipboard

import (
	"bytes"
	"encoding/binary"
	"strings"
	"sync"
	"time"

	"github.com/atotto/clipboard"
)

type Clipboard struct {
	Prompt     chan string
	LastText   string
	Mu         sync.RWMutex
	OutputText string
}

func NewClipboard() *Clipboard {
	return &Clipboard{
		Prompt:   make(chan string),
		LastText: "",
		Mu:       sync.RWMutex{},
	}
}

func (c *Clipboard) StartMonitoring() {
	for {
		c.Mu.Lock()
		text, err := clipboard.ReadAll()
		if err != nil || strings.TrimSpace(text) == "" {
			c.Mu.Unlock()
			time.Sleep(100 * time.Millisecond) // Add a small delay to prevent tight looping
			continue
		}

		if !isLikelyScreenshot(text) && text != c.LastText && text != c.OutputText {
			c.LastText = text
			c.Prompt <- text
		}
		c.Mu.Unlock()
		time.Sleep(100 * time.Millisecond) // Add a small delay to prevent tight looping
	}
}

func (c *Clipboard) GetClipboardText() (string, error) {
	return <-c.Prompt, nil
}

func Clear() error {
	return clipboard.WriteAll(" ")
}

func (c *Clipboard) SetClipboardText(text string) error {
	return clipboard.WriteAll(text)
}

func isLikelyScreenshot(data string) bool {
	// Check for common image file signatures
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

	// Check for clipboard format headers (Windows-specific)
	if len(dataBytes) >= 8 {
		format := binary.LittleEndian.Uint32(dataBytes[:4])
		if format == 2 || format == 8 || format == 17 {
			return true
		}
	}

	return false
}
