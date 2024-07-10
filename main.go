package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/api"
	"github.com/Codesmith28/cheatScript/internal"
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

	FileLocation    string
	HistoryLocation string
)

func init() {
	FileLocation = internal.GetFileLocation()
	HistoryLocation = internal.GetHistoryLocation()

	err := os.MkdirAll(filepath.Dir(FileLocation), os.ModePerm)
	checkNilErr(err)
	err = os.MkdirAll(filepath.Dir(HistoryLocation), os.ModePerm)
	checkNilErr(err)
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := tview.NewApplication().EnableMouse(true)

	if !api.CheckCredentials(FileLocation) {
		credentialModal := panes.CreateCredentialModal(app, func(apiInput string) {
			log.Println("ApiInput:", apiInput)
			err := os.WriteFile(FileLocation, []byte(apiInput), 0644)
			checkNilErr(err)

			log.Println("Starting clipboard monitoring after credential input.")
			setupMainUI(app)
		})

		app.SetRoot(credentialModal, true)

		err := app.Run()
		checkNilErr(err)
	} else {
		log.Println("Starting clipboard monitoring with existing credentials.")
		setupMainUI(app)
	}
}

func setupMainUI(app *tview.Application) {
	group2 := panes.CreateGroup2(HistoryPane, ModelsPane)
	group4 := panes.CreateGroup4(InputPane, PromptPane)
	group3 := panes.CreateGroup3(group4, OutputPane)
	group1 := panes.CreateGroup1(group2, group3)
	mainFlex := panes.CreateMainFlex(group1, KeybindingsPane)

	panes.SetupGlobalKeybindings(app, HistoryLocation)
	panes.InitHistoryPane(HistoryLocation)

	app.SetRoot(mainFlex, true)
	log.Println("Running app for main UI.")

	panes.StartClipboardMonitoring(app)
	panes.ApplySystemNavConfig(app)

	err := app.Run()
	checkNilErr(err)
}
