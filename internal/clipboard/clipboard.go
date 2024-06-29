package clipboard

import (
	"strings"
	"sync"

	"github.com/atotto/clipboard"
)

type Clipboard struct {
	Prompt   chan string
	LastText string
	Mu       sync.RWMutex
}

// create a new clipboard object
func NewClipboard() *Clipboard {
	return &Clipboard{
		Prompt:   make(chan string),
		LastText: "",
		Mu:       sync.RWMutex{},
	}
}

// start watching the clipboard for changes
func (c *Clipboard) StartMonitoring() {
	for {
		text, err := clipboard.ReadAll()

		if err != nil || strings.TrimSpace(text) == "" {
			continue // might need to find a better way to handle this
		}

		c.Mu.Lock()

		if text != c.LastText {
			c.LastText = text
			c.Prompt <- text
		}

		c.Mu.Unlock()
	}
}

// return the current clipboard text
func (c *Clipboard) GetClipboardText() (string, error) {
	return <-c.Prompt, nil
}

// clear the contents of the clipboard
func Clear() error {
	clipboard.WriteAll(" ")
	return nil
}

// set the clipboard text to the provided text
func (c *Clipboard) SetClipboardText(text string) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return clipboard.WriteAll(text)
}
