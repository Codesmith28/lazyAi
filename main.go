package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal/clipboard"
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
	go clipboard.StartMonitoring()
	clipboard.Clear()

	go func() {
		for {
			text, _ := clipboard.GetClipboardText()
			app.QueueUpdateDraw(func() {
				InputText.InputString = text
				panes.UpdateInputPane()
			})
			time.Sleep(1 * time.Second)
		}
	}()

	// Group 2: historyPane and promptPane in a vertical layout
	group2 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(historyPane, 0, 1, true).
		AddItem(promptPane, 0, 1, true)

	// Group 4: inputPane and modelsPane in a horizontal layout
	group4 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(inputPane, 0, 1, true).
		AddItem(modelsPane, 0, 1, true) // Set focusable to true

	// Group 3: Group 4 and outputPane in a vertical layout
	group3 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group4, 0, 1, true).
		AddItem(outputPane, 0, 2, true)

	// Group 1: Group 2 and Group 3 in a horizontal layout
	group1 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(group2, 0, 1, true).
		AddItem(group3, 0, 2, true)

	// Main Flex: Group 1 and keybindingsPane in a vertical layout
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group1, 0, 4, true).
		AddItem(keybindingsPane, 3, 1, true)

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

	// Set up the application root
	err := app.SetRoot(mainFlex, true).Run()
	checkNilErr(err)
}
