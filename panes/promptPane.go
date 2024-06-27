package panes

import (
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
)

var (
	PromptPane = tview.NewTextArea()
	PromptText internal.Prompt
)

func init() {
	PromptPane.SetText(PromptText.Prompt, true).
		SetWrap(true).
		SetBorder(true).
		SetTitle(" Prompt ").
		SetBorderPadding(1, 1, 2, 2)
}
