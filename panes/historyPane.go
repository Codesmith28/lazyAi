package panes

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/Codesmith28/cheatScript/internal/history"
)

var (
	HistoryPane = tview.NewList()
	History     = &history.History{}
)

func init() {
	loadHistory()

	HistoryPane.
		SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
			item := History.HistoryList[index]
			loadState(item)
		}).
		ShowSecondaryText(false).
		SetTitle(" History ").
		SetBorder(true).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyRune:
				switch event.Rune() {
				case 'd':
					deleteHistoryItem()
				}
			}
			return event
		})
}

// Load history from the JSON file and populate the history pane.
func loadHistory() {
	historyData, err := history.LoadHistory()
	checkNilErr(err)

	History = historyData
	updateHistoryPane()
}

// Update history pane with the current history items.
func updateHistoryPane() {
	HistoryPane.Clear()
	for _, item := range History.HistoryList {
		HistoryPane.AddItem(item.Query.PromptString+" - "+item.Date, "", 0, nil)
	}
}

// Save the current state as a history item.
func saveCurrentState() {
	query := history.Query{
		InputString:   InputText.InputString,
		PromptString:  PromptText.PromptString,
		SelectedModel: Selected.SelectedModel,
	}
	output := OutputText.OutputString
	err := history.AddHistoryItem(History, query, output)
	checkNilErr(err)
	updateHistoryPane()

	createNewState()
}

// Create a new state preserving only the prompt and selected model.
func createNewState() {
	InputText.InputString = ""
	OutputText.OutputString = ""
	updatePanes(Selected.SelectedModel)
}

// Load a state from a history item.
func loadState(item history.HistoryItem) {
	InputText.InputString = item.Query.InputString
	PromptText.PromptString = item.Query.PromptString
	Selected.SelectedModel = item.Query.SelectedModel
	OutputText.OutputString = item.Output
	updatePanes(item.Query.SelectedModel)
}

// Delete the selected history item from the history list
func deleteHistoryItem() {
	selectedIndex := HistoryPane.GetCurrentItem()
	History.HistoryList = append(
		History.HistoryList[:selectedIndex],
		History.HistoryList[selectedIndex+1:]...)
	err := history.SaveHistory(History)

	checkNilErr(err)
	updateHistoryPane()
}

// Update all panes with the current state.
func updatePanes(queryModel string) {
	UpdateInputPane()
	UpdatePromptPane()
	UpdateOutputPane()
	SelectModel(queryModel)
}
