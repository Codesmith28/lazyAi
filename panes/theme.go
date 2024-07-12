package panes

import (
	"log"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal"
)

var Theme = internal.Theme{
	ActiveBorderCol:   tcell.ColorGreen,
	InactiveBorderCol: tcell.ColorWhite,
	SecondaryCol:      tcell.ColorBlue,
	TitleCol:          tcell.ColorYellow,
}

var lastFocused tview.Primitive

// applied when used in setting panes:
func ApplyTheme(pane tview.Primitive) tview.Primitive {
	switch p := pane.(type) {
	case *tview.TextView:
		p.SetTitleColor(Theme.TitleCol)
	case *tview.TextArea:
		p.SetTitleColor(Theme.TitleCol)
	case *tview.InputField:
		p.SetFieldTextColor(Theme.SecondaryCol).
			SetTitleColor(Theme.TitleCol)
	case *tview.List:
		p.SetSelectedBackgroundColor(Theme.SecondaryCol).
			SetTitleColor(Theme.TitleCol)
	case *tview.Flex:
		p.SetTitleColor(Theme.TitleCol)
	default:
		log.Printf("Unknown pane type: %T", p)
	}
	return pane
}

func ApplyFocusedStyle(p tview.Primitive) {
	switch v := p.(type) {
	case *tview.TextView:
		v.SetBorderColor(Theme.ActiveBorderCol)
	case *tview.TextArea:
		v.SetBorderColor(Theme.ActiveBorderCol)
	case *tview.InputField:
		v.SetBorderColor(Theme.ActiveBorderCol)
	case *tview.List:
		v.SetBorderColor(Theme.ActiveBorderCol)
	}
}

func ApplyUnfocusedStyle(p tview.Primitive) {
	switch v := p.(type) {
	case *tview.TextView:
		v.SetBorderColor(Theme.InactiveBorderCol)
	case *tview.TextArea:
		v.SetBorderColor(Theme.InactiveBorderCol)
	case *tview.InputField:
		v.SetBorderColor(Theme.InactiveBorderCol)
	case *tview.List:
		v.SetBorderColor(Theme.InactiveBorderCol)
	}
}
