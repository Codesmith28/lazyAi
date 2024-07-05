package panes

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/internal/clipboard"
	core "github.com/Codesmith28/cheatScript/internal/queue"
)

var (
	OutputPane = tview.NewTextView()
	OutputText = &internal.Output{}
)

func init() {
	OutputText = &internal.Output{
		OutputString: "",
	}

	OutputPane.SetText(OutputText.OutputString)

	OutputPane.
		SetWrap(true).
		SetScrollable(true).
		SetBorder(true).
		SetTitle(" Output ").
		SetBorderPadding(1, 1, 2, 2).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyUp:
				currRow, _ := OutputPane.GetScrollOffset()
				OutputPane.ScrollTo(currRow-1, 0)
			case tcell.KeyDown:
				currRow, _ := OutputPane.GetScrollOffset()
				OutputPane.ScrollTo(currRow+1, 0)
			}
			return event
		})
}

func StartOutputMonitoring(
	app *tview.Application,
	clipboard *clipboard.Clipboard,
	queue *core.Queue,
) {
	go func() {
		for {
			// Consume the response
			message, err := queue.Consume()
			if err != nil {
				fmt.Println("Error consuming message:", err)
				time.Sleep(time.Second) // Add a delay to avoid tight loop on errors
				continue
			}

			app.QueueUpdateDraw(func() {
				outputMessage := fmt.Sprintf(message)
				OutputText.OutputString = outputMessage
				OutputPane.SetText(outputMessage)
			})

			clipboard.Mu.Lock()
			clipboard.OutputText = message
			err = clipboard.SetClipboardText(message)
			checkNilErr(err)
			clipboard.Mu.Unlock()
		}
	}()
}
