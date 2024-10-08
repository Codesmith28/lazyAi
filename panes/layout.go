package panes

import (
	"log"

	"github.com/rivo/tview"
)

func CreateGroup2(historyPane *tview.List, modelsPane *tview.List) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(historyPane, 0, 3, true).
		AddItem(modelsPane, 0, 1, true)
}

func CreateGroup4(inputPane *tview.TextView, promptPane *tview.TextArea) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(promptPane, 0, 2, true).
		AddItem(inputPane, 0, 2, true)
}

func CreateGroup3(group4 *tview.Flex, outputPane *tview.TextView) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group4, 0, 1, true).
		AddItem(outputPane, 0, 1, true)
}

func CreateGroup1(group2, group3 *tview.Flex) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(group2, 0, 1, true).
		AddItem(group3, 0, 3, true)
}

func CreateMainFlex(group1 *tview.Flex, keybindingsPane *tview.TextView) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group1, 0, 20, true).
		AddItem(keybindingsPane, 1, 0, false)
}

func SetupMainUILayout(app *tview.Application) {
	if app == nil {
		StartClipboardMonitoring(nil)
		ApplySystemNavConfig(nil)

		select {}
	} else {
		group2 := CreateGroup2(HistoryPane, ModelList)
		group4 := CreateGroup4(InputPane, PromptPane)
		group3 := CreateGroup3(group4, OutputPane)
		group1 := CreateGroup1(group2, group3)
		mainFlex := CreateMainFlex(group1, KeybindingsPane)

		SetupGlobalKeybindings(app)
		InitHistoryPane()

		// list of all the panes:
		allPanes := []tview.Primitive{
			HistoryPane, ModelList, InputPane, PromptPane, OutputPane, KeybindingsPane,
		}

		// apply theme to all panes:
		for _, pane := range allPanes {
			ApplyTheme(pane)
			// if pane.HasFocus() {
			// 	ApplyFocusedStyle(pane)
			// } else {
			// 	ApplyUnfocusedStyle(pane)
			// }
		}

		app.SetRoot(mainFlex, true)
		log.Println("Running app for main UI.")

		StartClipboardMonitoring(app)
		ApplySystemNavConfig(app)

		err := app.Run()
		checkNilErr(err)
	}
}
