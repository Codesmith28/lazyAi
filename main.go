package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/api"
	"github.com/Codesmith28/cheatScript/internal"
	"github.com/Codesmith28/cheatScript/panes"
)

var (
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

	detachedMode := flag.Bool("d", false, "Run in detached mode")
	defaultPrompt := flag.String("p", "", "Set the default prompt")
	flag.Parse()

	if *defaultPrompt != "" {
		PromptText.PromptString = *defaultPrompt
	}

	app := tview.NewApplication().EnableMouse(true)

	if !api.CheckCredentials(FileLocation, nil) {
		credentialModal := panes.CreateCredentialModal(app, func(apiInput string) {

			if ok := api.CheckCredentials("", &apiInput); !ok {
				app.Stop()
				log.Println("Invalid API key. Please try again.")
				return
			}

			err := os.WriteFile(FileLocation, []byte(apiInput), 0644)
			checkNilErr(err)

			if *detachedMode {
				runDetachedMode()
			} else {
				setupMainUI(app)
			}
		})

		app.SetRoot(credentialModal, true)

		err := app.Run()
		checkNilErr(err)
	} else {
		if *detachedMode {
			runDetachedMode()
		} else {
			setupMainUI(app)
		}
	}
}

func runDetachedMode() {
	panes.StartClipboardMonitoring(nil)
	panes.ApplySystemNavConfig(nil)

	select {}
}

func setupMainUI(app *tview.Application) {
	group2 := panes.CreateGroup2(panes.HistoryPane, ModelsPane)
	group4 := panes.CreateGroup4(InputPane, PromptPane)
	group3 := panes.CreateGroup3(group4, OutputPane)
	group1 := panes.CreateGroup1(group2, group3)
	mainFlex := panes.CreateMainFlex(group1, KeybindingsPane)

	panes.SetupGlobalKeybindings(app, HistoryLocation)
	panes.InitHistoryPane(HistoryLocation)

	app.SetRoot(mainFlex, true)

	panes.StartClipboardMonitoring(app)
	panes.ApplySystemNavConfig(app)

	err := app.Run()
	checkNilErr(err)
}
