package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/panes"
)

var (
	historyPane     = panes.HistoryPane
	modelsPane      = panes.ModelsPane
	outputPane      = panes.OutputPane
	promptPane      = panes.PromptPane
	keybindingsPane = panes.KeybindingsPane
	inputPane       = panes.InputPane

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
	panes.StartClipboardMonitoring(app)

	// Create layout groups
	group2 := panes.CreateGroup2(historyPane, promptPane)
	group4 := panes.CreateGroup4(inputPane, modelsPane)
	group3 := panes.CreateGroup3(group4, outputPane)
	group1 := panes.CreateGroup1(group2, group3)
	mainFlex := panes.CreateMainFlex(group1, keybindingsPane)

	// Set up global keybindings to focus on each pane
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				app.SetFocus(inputPane)
				return nil
			case '2':
				app.SetFocus(ModelList)
				return nil
			case '3':
				app.SetFocus(outputPane)
				return nil
			case '4':
				app.SetFocus(historyPane)
				return nil
			case '5':
				app.SetFocus(promptPane)
				return nil
			case '?':
				app.SetFocus(keybindingsPane)
				return nil
			}
		}
		return event
	})

	// Set the default focus to inputPane
	app.SetFocus(inputPane)

	// Set up the application root
	err := app.SetRoot(mainFlex, true).Run()
	checkNilErr(err)
}
