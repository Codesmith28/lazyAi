package history

import (
	"encoding/json"
	"os"
	"time"
)

const historyFile = "history.json"

type Query struct {
	InputString   string
	PromptString  string
	SelectedModel string
}

type HistoryItem struct {
	Query  Query
	Output string
	Date   string
}

type History struct {
	HistoryList []HistoryItem
}

// LoadHistory loads history from the JSON file.
func LoadHistory() (*History, error) {
	history := &History{}
	if _, err := os.Stat(historyFile); os.IsNotExist(err) {
		return history, nil
	}

	data, err := os.ReadFile(historyFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, history)
	if err != nil {
		return nil, err
	}

	return history, nil
}

// SaveHistory saves history to the JSON file.
func SaveHistory(history *History) error {
	data, err := json.Marshal(history)
	if err != nil {
		return err
	}

	err = os.WriteFile(historyFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// AddHistoryItem adds a new history item and saves it to the JSON file.
func AddHistoryItem(history *History, query Query, output string) error {

	if query.InputString == "" {
		return nil
	}

	historyItem := HistoryItem{
		Query:  query,
		Output: output,
		Date:   time.Now().Format("Monday, January 2, 2006"),
	}

	history.HistoryList = append(history.HistoryList, historyItem)
	return SaveHistory(history)
}
