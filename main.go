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
	homeDir, err := os.UserHomeDir()
	checkNilErr(err)

	FileLocation = filepath.Join(homeDir, ".local/share/lazyAi/lazy_ai_api")
	HistoryLocation = filepath.Join(homeDir, ".local/share/lazyAi/history.json")

	err = os.MkdirAll(filepath.Dir(FileLocation), os.ModePerm)
	checkNilErr(err)
	err = os.MkdirAll(filepath.Dir(HistoryLocation), os.ModePerm)
	checkNilErr(err)
}

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCredentials() bool {
	apiKey, err := os.ReadFile(FileLocation)
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
					{"text": "Explain how AI works"},
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

	if resp.StatusCode != http.StatusOK {
		return false
	}

	return true
}

func main() {
	app := tview.NewApplication().EnableMouse(true)

	if !checkCredentials() {
		credentialModal := panes.CreateCredentialModal(app, func(apiInput string) {
			log.Println("ApiInput:", apiInput)
			err := os.WriteFile(FileLocation, []byte(apiInput), 0644)
			checkNilErr(err)

			log.Println("Starting clipboard monitoring after credential input.")
			panes.StartClipboardMonitoring(app)
			setupMainUI(app)
		})

		app.SetRoot(credentialModal, true)
		log.Println("Running app for credential input.")

		err := app.Run()
		checkNilErr(err)
	} else {
		log.Println("Starting clipboard monitoring with existing credentials.")
		panes.StartClipboardMonitoring(app)
		setupMainUI(app)
	}
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
	log.Println("Running app for main UI.")
	err := app.Run()
	checkNilErr(err)
}
