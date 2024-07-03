package panes

import (
	"encoding/json"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/internal/clipboard"
	core "github.com/Codesmith28/cheatScript/internal/queue"
)

var (
	InputPane = tview.NewTextView()
	InputText *internal.Input
)

func init() {
	InputText = &internal.Input{}

	InputPane.
		SetWrap(true).
		SetScrollable(true).
		SetBorder(true).
		SetTitle(" Input ").
		SetBorderPadding(1, 1, 2, 2).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyUp:
				currRow, _ := InputPane.GetScrollOffset()
				InputPane.ScrollTo(currRow-1, 0)
			case tcell.KeyDown:
				currRow, _ := InputPane.GetScrollOffset()
				InputPane.ScrollTo(currRow+1, 0)
			}
			return event
		})
}

var once sync.Once

func StartClipboardMonitoring(app *tview.Application) {
	clipboard.Clear()
	clipboard := clipboard.NewClipboard()
	var lastPublishedText string

	go clipboard.StartMonitoring()

	// initialize the queue
	queue := core.NewQueue()
	var err error = queue.Connect()

	checkNilErr(err)

	go func() {
		for {
			text, err := clipboard.GetClipboardText()
			checkNilErr(err)

			if text != lastPublishedText {
				app.QueueUpdateDraw(func() {
					InputText.InputString = text
					UpdateInputPane()
				})

				prompt := &internal.Prompt{
					PromptString: text,
					Model:        Selected.SelectedModel,
				}

				promptBytes, err := json.Marshal(prompt)
				checkNilErr(err)

				promptJson := string(promptBytes)

				err = queue.Publish(promptJson)
				checkNilErr(err)

				lastPublishedText = text

				once.Do(func() {
					StartOutputMonitoring(app, clipboard)
				})
			}
		}
	}()
}
