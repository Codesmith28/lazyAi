package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
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

func UpdateInputPane() {
	InputPane.SetText(InputText.InputString)
}
