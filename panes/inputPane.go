package panes

import (
	"sync"
	"time"

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
			case tcell.KeyRune:
				switch event.Rune() {
				case '1', '2', '3', '4', '5', '?':
					return nil
				}
			}
			return event
		})
}

var once sync.Once

func StartClipboardMonitoring(app *tview.Application, clipboard *clipboard.Clipboard, OutputPane *tview.TextView) {

	go clipboard.StartMonitoring()

	// intialize the queue
	queue := core.NewQueue()
	var err error = queue.Connect()

	if err != nil {
		panic(err)
	}

	consumerQueue := core.NewQueue()

	go func() {
		for {
			text, err := clipboard.GetClipboardText()

			if err != nil {
				panic(err)
			}

			app.QueueUpdateDraw(func() {
				InputText.InputString = text
				UpdateInputPane()
			})

			err = queue.Publish(text)

			if err != nil {
				panic(err)
			}

			once.Do(func() {
				go consumerQueue.Consume(clipboard)
			})

			message, _ := consumerQueue.GetMessages()

			app.QueueUpdateDraw(func() {
				OutputPane.SetText(message)
			})

			clipboard.LastText = message
			clipboard.SetClipboardText(message)

			time.Sleep(1 * time.Second)
		}
	}()
}
