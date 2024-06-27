package panes

import "github.com/rivo/tview"

func CreateGroup2(historyPane *tview.Box, promptPane *tview.TextArea) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(historyPane, 0, 1, true).
		AddItem(promptPane, 0, 1, true)
}

func CreateGroup4(inputPane *tview.TextView, modelsPane *tview.Flex) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(inputPane, 0, 1, true).
		AddItem(modelsPane, 0, 1, true)
}

func CreateGroup3(group4 *tview.Flex, outputPane *tview.Box) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group4, 0, 1, true).
		AddItem(outputPane, 0, 2, true)
}

func CreateGroup1(group2, group3 *tview.Flex) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(group2, 0, 1, true).
		AddItem(group3, 0, 2, true)
}

func CreateMainFlex(group1 *tview.Flex, keybindingsPane *tview.Box) *tview.Flex {
	return tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(group1, 0, 4, true).
		AddItem(keybindingsPane, 3, 1, true)
}
