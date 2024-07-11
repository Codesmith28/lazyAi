package panes

import (
	"log"

	"github.com/gdamore/tcell/v2"

	"github.com/Codesmith28/cheatScript/internal"
)

var Theme = internal.Theme{
	ActiveBorderCol:   tcell.ColorGreen.TrueColor(),
	InactiveBorderCol: tcell.ColorGray.TrueColor(),
	PrimaryTextCol:    tcell.ColorWhite.TrueColor(),
	SecondaryTextCol:  tcell.ColorBlue.TrueColor(),
	TertiaryTextCol:   tcell.ColorYellow.TrueColor(),
	TitleCol:          tcell.ColorRed.TrueColor(),
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
