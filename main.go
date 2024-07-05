package main

import (
	"github.com/rivo/tview"

	core "github.com/Codesmith28/cheatScript/internal/queue"
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
	OutputText = panes.OutputText
	Selected   = panes.Selected
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := tview.NewApplication().EnableMouse(true)

	queue := core.NewQueue()
	err := queue.Connect()
	checkNilErr(err)
	defer queue.Close()

	panes.StartClipboardMonitoring(app, queue)

	// Create layout groups
	group2 := panes.CreateGroup2(HistoryPane, ModelsPane)
	group4 := panes.CreateGroup4(InputPane, PromptPane)
	group3 := panes.CreateGroup3(group4, OutputPane)
	group1 := panes.CreateGroup1(group2, group3)
	mainFlex := panes.CreateMainFlex(group1, KeybindingsPane)

	// Set up global KeybindingsPane
	panes.SetupGlobalKeybindings(app)

	// Set up the application root
	err = app.SetRoot(mainFlex, true).Run()
	checkNilErr(err)
}
