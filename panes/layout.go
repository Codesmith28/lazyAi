package panes

import "github.com/rivo/tview"

func CreateGroup2(historyPane *tview.List, modelsPane *tview.Flex) *tview.Flex {
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
