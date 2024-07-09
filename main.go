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
	OutputText = panes.OutputText
	Selected   = panes.Selected
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkCredentials() bool {
	return false // Placeholder
}

func main() {
	app := tview.NewApplication().EnableMouse(true)

	panes.StartClipboardMonitoring(app)

	// Check for credentials
	if !checkCredentials() {
		credentialModal := panes.CreateCredentialModal(app, func(username, password string) {
			println("Username:", username, "Password:", password)

			// After handling credentials, set up the main UI
			setupMainUI(app)
		})
		app.SetRoot(credentialModal, true)
		err := app.Run()
		checkNilErr(err)
	} else {
		setupMainUI(app)
	}
}

func setupMainUI(app *tview.Application) {
	group2 := panes.CreateGroup2(HistoryPane, ModelsPane)
	group4 := panes.CreateGroup4(InputPane, PromptPane)
	group3 := panes.CreateGroup3(group4, OutputPane)
	group1 := panes.CreateGroup1(group2, group3)
	mainFlex := panes.CreateMainFlex(group1, KeybindingsPane)

	panes.SetupGlobalKeybindings(app)

	app.SetRoot(mainFlex, true)
	err := app.Run()
	checkNilErr(err)
}
