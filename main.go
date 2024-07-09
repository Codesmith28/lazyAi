package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rivo/tview"

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

	filelocation    string
	historyLocation string
)

func init() {
	filelocation = internal.GetFileLocation()
	historyLocation = internal.GetHistoryLocation()
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCredentials() bool {
	apiKey, err := os.ReadFile(filelocation)
	if err != nil {
		return false
	}

	// Validate the API key by pinging an endpoint
	apiKeyStr := strings.TrimSpace(string(apiKey))
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s",
		apiKeyStr,
	)

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": "Ping, reply with a 200 status code."},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	checkNilErr(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	checkNilErr(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	checkNilErr(err)
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

func main() {
	app := tview.NewApplication().EnableMouse(true)

	// Check for credentials
	if !checkCredentials() {
		// Ensure the directory exists
		err := os.MkdirAll(filepath.Dir(filelocation), os.ModePerm)
		checkNilErr(err)

		credentialModal := panes.CreateCredentialModal(app, func(apiInput string) {
			log.Println("ApiInput:", apiInput)
			err := os.WriteFile(filelocation, []byte(apiInput), 0644)
			checkNilErr(err)

			log.Println("Starting clipboard monitoring after credential input.")
			panes.StartClipboardMonitoring(app)
			panes.ApplySystemNavConfig(app)
			setupMainUI(app)
		})

		app.SetRoot(credentialModal, true)

		err = app.Run()
		checkNilErr(err)
	} else {
		panes.StartClipboardMonitoring(app)
		panes.ApplySystemNavConfig(app)
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
	log.Println("Running app for main UI.")
	err := app.Run()
	checkNilErr(err)
}
