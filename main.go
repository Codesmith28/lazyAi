package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rivo/tview"

	"github.com/Codesmith28/lazyAi/api"
	"github.com/Codesmith28/lazyAi/internal"
	"github.com/Codesmith28/lazyAi/panes"
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

	logFile, _ := os.OpenFile("lazyai.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	log.SetOutput(logFile)
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
		panes.SetupMainUILayout(nil)
	} else {
		panes.SetupMainUILayout(app)
	}
}

func main() {
	detachedMode := flag.Bool("d", false, "Run in detached mode")
	defaultPrompt := flag.String("p", "", "Set the default prompt")
	helpCommand := flag.Bool("help", false, "Show help commands")
	flag.Parse()

	if (*defaultPrompt != "" || *detachedMode) && *helpCommand {
		fmt.Println("Not allowed to use -p or -d and -help together. Exiting...")
		os.Exit(1)
	}

	if *helpCommand {
		fmt.Println(panes.HelpCommands)
		os.Exit(0)
	}

	if *defaultPrompt != "" {
		PromptText.PromptString = *defaultPrompt
	}

	app := tview.NewApplication().EnableMouse(true)

	if !api.CheckCredentials(FileLocation, nil) {
		credentialModal := panes.CreateCredentialModal(app, func(apiInput string) {
			if ok := api.CheckCredentials("", &apiInput); !ok {
				app.Stop()
				fmt.Println("Invalid API key. Please try again.")
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
