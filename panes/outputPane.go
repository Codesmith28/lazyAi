package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/getlantern/systray"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/api"
	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/internal/clipboard"
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

func HandlePromptChange(
	query *internal.Query,
	clipboard *clipboard.Clipboard,
	app *tview.Application,
) {
	content, err := api.SendPrompt(
		query.PromptString,
		query.SelectedModel,
		query.InputString,
		nil,
	)

	if err != nil {
		panic(err)
	}

	app.QueueUpdateDraw(func() {
		OutputPane.SetText(content)
	})

	clipboard.Mu.Lock()
	OutputText.OutputString = content
	clipboard.OutputText = content
	err = clipboard.SetClipboardText(content)

	if err != nil {
		panic(err)
	}

	clipboard.Mu.Unlock()

	systray.SetTooltip("Ready!!")
}
