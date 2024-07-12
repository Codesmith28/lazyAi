package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
)

var (
	PromptPane = tview.NewTextArea()
	PromptText *internal.Prompt
)

func init() {
	PromptText = &internal.Prompt{}
	PromptPane.SetText(PromptText.PromptString, true).
		SetWrap(true).
		SetBorder(true).
		SetTitle(" Prompt ").
		SetBorderPadding(1, 1, 2, 2).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Key() == tcell.KeyRune && event.Rune() == '5' {
				return nil
			}
			return event
		})

	PromptPane.SetChangedFunc(func() {
		PromptText.PromptString = PromptPane.GetText()
	})
}
