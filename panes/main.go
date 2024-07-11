package panes

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

var (
	ActiveBorderCol   = tcell.ColorGreen
	InactiveBorderCol = tcell.ColorDarkGray
)

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
