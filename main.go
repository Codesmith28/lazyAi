package main

import (
	"time"

	"github.com/Codesmith28/cheatScript/internal/clipboard"
	"github.com/Codesmith28/cheatScript/panes"
	"github.com/rivo/tview"
)

var (
	historyPane     = panes.HistoryPane
	modelsPane      = panes.ModelsPane
	outputPane      = panes.OutputPane
	promptPane      = panes.PromptPane
	keybindingsPane = panes.KeybindingsPane
	textPane        = panes.TextView
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := tview.NewApplication()
	go clipboard.StartMonitoring()

	// Group 2: historyPane and promptPane in a vertical layout
	group2 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(historyPane, 0, 1, false).
		AddItem(promptPane, 0, 1, false)

	textView := textPane

	textView.SetWrap(true).SetScrollable(true).SetBorder(true).SetTitle(" Input Data: ").SetBorderPadding(1, 1, 2, 2)

	// Group 4: inputPane and modelsPane in a horizontal layout
	group4 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(textView, 0, 3, false).
		AddItem(modelsPane, 0, 1, false)

	go func() {
		for {
			text, _ := clipboard.GetClipboardText()
			app.QueueUpdateDraw(func() {
				textView.SetText("Prompt: " + text)
			})

			time.Sleep(1 * time.Second)
		}
	}()

	// Group 3: Group 4 and outputPane in a vertical layout
	group3 := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group4, 0, 9, false).
		AddItem(outputPane, 0, 10, false)

	// Group 1: Group 2 and Group 3 in a horizontal layout
	group1 := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(group2, 0, 1, false).
		AddItem(group3, 0, 2, false)

	// Main Flex: Group 1 and keybindingsPane in a vertical layout
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group1, 0, 4, false).
		AddItem(keybindingsPane, 3, 1, false) // Setting the weight to 1 for better visibility

	// Set up the application root
	err := app.SetRoot(mainFlex, true).Run()
	checkNilErr(err)
}
