package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/getlantern/systray"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/internal/clipboard"
	querymaker "github.com/Codesmith28/cheatScript/internal/queryMaker"
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

func StartClipboardMonitoring(app *tview.Application) {
	clipboard.Clear()
	clipboard := clipboard.NewClipboard()
	var lastPublishedText string

	go clipboard.StartMonitoring()

	go func() {
		for {
			text, err := clipboard.GetClipboardText()
			checkNilErr(err)
			if text != lastPublishedText {
				app.QueueUpdateDraw(func() {
					InputText.InputString = text
					UpdateInputPane()
				})

				promptString := PromptText.PromptString
				selectedModel := Selected.SelectedModel

				localQuery := querymaker.MakeQuery(text, promptString, selectedModel)
				lastPublishedText = text

				systray.SetTooltip("Processing...")
				go HandlePromptChange(&localQuery, clipboard, app)
			}
		}
	}()
}
