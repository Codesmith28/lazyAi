package panes

import (
	"log"
)

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
