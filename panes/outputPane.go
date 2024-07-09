package panes

import (
	"fmt"
	"time"

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

func HandlePromptChange(query *internal.Query, clipboard *clipboard.Clipboard, app *tview.Application) {
	content, err := api.SendPrompt(query.PromptString, query.SelectedModel, query.InputString)

	if err != nil {
		OutputText.OutputString = fmt.Sprintf("Error: %s", err)
	} else {
		OutputText.OutputString = content
	}

	app.QueueUpdateDraw(func() {
		OutputPane.SetText(OutputText.OutputString)
	})

	if err == nil {
		clipboard.Mu.Lock()
		clipboard.OutputText = content
		err = clipboard.SetClipboardText(content)
		checkNilErr(err)
		clipboard.Mu.Unlock()
	}

	systray.SetTooltip("Ready!!")

	time.Sleep(1 * time.Second)

	systray.SetTooltip("Started!!")
}
