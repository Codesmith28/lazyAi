package panes

import "github.com/rivo/tview"

var InputPane = tview.NewTextView()

func init() {
	InputPane.
		SetWrap(true).
		SetScrollable(true).
		SetBorder(true).
		SetTitle(" Input ").
		SetBorderPadding(1, 1, 2, 2)
}
