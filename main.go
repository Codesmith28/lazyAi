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
	FileLocation    string
	HistoryLocation string
	PromptText      = panes.PromptText
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

func setupUI(detachedMode *bool, app *tview.Application) {
	if *detachedMode {
		if app != nil {
			app.Stop()
		}
		panes.SetupMainUI(nil)
	} else {
		panes.SetupMainUI(app)
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

			log.Println("Starting clipboard monitoring after credential input.")
			setupUI(detachedMode, app)
		})

		app.SetRoot(credentialModal, true)

		err := app.Run()
		checkNilErr(err)
	} else {
		log.Println("Starting clipboard monitoring with existing credentials.")
		setupUI(detachedMode, app)
	}
}
