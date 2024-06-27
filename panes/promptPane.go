package panes

import (
	"github.com/rivo/tview"
)

// var PromptPane = tview.NewBox().SetBorder(true).SetTitle(" Prompt ")
var PromptPane = tview.NewTextArea()

func init() {
	PromptPane.
		SetWrap(true).
		SetBorder(true).
		SetTitle(" Prompt ").
		SetBorderPadding(1, 1, 2, 2)
}
