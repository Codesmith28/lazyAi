package main

import (
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/panes"
)

var (
	HistoryPane     = panes.HistoryPane
	ModelsPane      = panes.ModelsPane
	OutputPane      = panes.OutputPane
	PromptPane      = panes.PromptPane
	KeybindingsPane = panes.KeybindingsPane
	InputPane       = panes.InputPane

	ModelList  = panes.ModelList
	PromptText = panes.PromptText
	InputText  = panes.InputText
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := tview.NewApplication().EnableMouse(true)
	panes.StartClipboardMonitoring(app, OutputPane)

	// Create layout groups
	group2 := panes.CreateGroup2(HistoryPane, PromptPane)
	group4 := panes.CreateGroup4(InputPane, ModelsPane)
	group3 := panes.CreateGroup3(group4, OutputPane)
	group1 := panes.CreateGroup1(group2, group3)
	mainFlex := panes.CreateMainFlex(group1, KeybindingsPane)

	// Set up global KeybindingsPane
	panes.SetupGlobalKeybindings(app)

	// Set up the application root
	err := app.SetRoot(mainFlex, true).Run()
	checkNilErr(err)
}
