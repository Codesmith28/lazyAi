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
	PrimaryTextCol:    tcell.ColorWhite,
	SecondaryTextCol:  tcell.ColorBlue,
	TitleCol:          tcell.ColorYellow,
}

func ApplyTheme(pane tview.Primitive) tview.Primitive {
	switch p := pane.(type) {
	case *tview.TextView:
		p.SetTextColor(Theme.PrimaryTextCol).
			SetBorderColor(Theme.InactiveBorderCol).
			SetTitleColor(Theme.TitleCol)
	case *tview.TextArea:
		p.SetBorderColor(Theme.InactiveBorderCol).
			SetTitleColor(Theme.TitleCol)
	case *tview.InputField:
		p.SetFieldTextColor(Theme.PrimaryTextCol).
			SetFieldBackgroundColor(Theme.InactiveBorderCol).
			SetTitleColor(Theme.TitleCol)
	case *tview.DropDown:
		p.SetFieldTextColor(Theme.PrimaryTextCol).
			SetFieldBackgroundColor(Theme.InactiveBorderCol).
			SetTitleColor(Theme.TitleCol)
	case *tview.List:
		p.SetMainTextColor(Theme.PrimaryTextCol).
			SetSelectedBackgroundColor(Theme.SecondaryTextCol).
			SetTitleColor(Theme.TitleCol).
			SetBorderColor(Theme.InactiveBorderCol)
	case *tview.Flex:
		p.SetBorderColor(Theme.InactiveBorderCol).
			SetTitleColor(Theme.TitleCol)
	default:
		log.Printf("Unknown pane type: %T", p)
	}
	return pane
}

func UpdateInputPane() {
	InputPane.SetText(InputText.InputString)
}

func UpdateOutputPane() {
	OutputPane.SetText(OutputText.OutputString)
}

func UpdatePromptPane() {
	PromptPane.SetText(PromptText.PromptString, true)
}

func SelectModel(model string) {
	Selected.SelectedModel = model
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
