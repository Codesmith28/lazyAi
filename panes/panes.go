package panes

import (
	"log"

	"github.com/rivo/tview"
)

func SetupMainUI(app *tview.Application, HistoryLocation string) {
	group2 := CreateGroup2(HistoryPane, ModelsPane)
	group4 := CreateGroup4(InputPane, PromptPane)
	group3 := CreateGroup3(group4, OutputPane)
	group1 := CreateGroup1(group2, group3)
	mainFlex := CreateMainFlex(group1, KeybindingsPane)

	SetupGlobalKeybindings(app, HistoryLocation)
	InitHistoryPane(HistoryLocation)

	app.SetRoot(mainFlex, true)
	log.Println("Running app for main UI.")

	StartClipboardMonitoring(app)
	ApplySystemNavConfig(app)

	err := app.Run()
	checkNilErr(err)
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
