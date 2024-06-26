package clipboard

import (
	"fmt"
	"time"

	"github.com/atotto/clipboard"
)

var prompt = make(chan string)
var lastText string

// start watching the clipboard for changes
func StartMonitoring() {
	for {
		text, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("Error reading clipboard:", err)
			continue
		}

		if text != lastText {
			prompt <- text
			lastText = text
		}

		time.Sleep(1 * time.Second)
	}
}

// return the current clipboard text
func GetClipboardText() (string, error) {
	return <-prompt, nil
}
