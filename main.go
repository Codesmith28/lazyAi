package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/panes"
)

var (
	historyPane     = panes.HistoryPane
	inputPane       = panes.InputPane
	modelsPane      = panes.ModelsPane
	outputPane      = panes.OutputPane
	promptPane      = panes.PromptPane
	keybindingsPane = panes.KeybindingsPane
	modelList       = panes.ModelList // Export ModelList to access it
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := tview.NewApplication()

	// Group 2: historyPane and promptPane in a vertical layout
	group2 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(historyPane, 0, 1, false).
		AddItem(promptPane, 0, 1, false)

	// Group 4: inputPane and modelsPane in a horizontal layout
	group4 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(inputPane, 0, 1, false).
		AddItem(modelsPane, 0, 1, true) // Set focusable to true

	// Group 3: Group 4 and outputPane in a vertical layout
	group3 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group4, 0, 1, false).
		AddItem(outputPane, 0, 2, false)

	// Group 1: Group 2 and Group 3 in a horizontal layout
	group1 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(group2, 0, 1, false).
		AddItem(group3, 0, 2, false)

	// Main Flex: Group 1 and keybindingsPane in a vertical layout
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group1, 0, 4, false).
		AddItem(keybindingsPane, 3, 1, false)

	// Set up global keybindings to focus on each pane
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case '1':
				app.SetFocus(historyPane)
			case '2':
				app.SetFocus(modelList)
			case '3':
				app.SetFocus(inputPane)
			case '4':
				app.SetFocus(outputPane)
			case '5':
				app.SetFocus(promptPane)
			case '6':
				app.SetFocus(keybindingsPane)
			}
		}
		return event
	})

	// Set up the application root
	err := app.SetRoot(mainFlex, true).Run()
	checkNilErr(err)
}
