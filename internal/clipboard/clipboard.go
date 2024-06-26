package clipboard

import (
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
			continue // might need to find a better way to handle this
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

func Clear() error {
	return clipboard.WriteAll(" ")
}
