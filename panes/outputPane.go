package panes

import (
	"fmt"

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

func StartOutputMonitoring(app *tview.Application, clipboard *clipboard.Clipboard) {
	consumerQueue := core.NewQueue()
	consumerQueue.Consume(clipboard)

	go func() {
		for {
			message, _ := consumerQueue.GetMessages()

			app.QueueUpdateDraw(func() {
				OutputText.OutputString = message
				OutputPane.SetText(message)
			})

			clipboard.Mu.Lock()

			clipboard.OutputText = message
			err := clipboard.SetClipboardText(message)
			if err != nil {
				fmt.Println("Error setting clipboard text: ", err)
			}
			clipboard.Mu.Unlock()
		}
	}()
}
