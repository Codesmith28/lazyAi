package history

import (
	"encoding/json"
	"os"
	"time"

	"github.com/Codesmith28/cheatScript/internal"
)

type (
	Query       = internal.Query
	HistoryItem = internal.HistoryItem
	History     = internal.History
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

// LoadHistory loads history from the specified JSON file.
func LoadHistory(historyFile string) (*History, error) {
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

// SaveHistory saves history to the specified JSON file.
func SaveHistory(history *History, historyFile string) error {
	data, err := json.Marshal(history)
	checkNilErr(err)

	err = os.WriteFile(historyFile, data, 0644)
	checkNilErr(err)

	return nil
}

// AddHistoryItem adds a new history item and saves it to the specified JSON file.
func AddHistoryItem(history *History, query Query, output, historyFile string) error {
	if query.InputString == "" {
		return nil
	}

	historyItem := HistoryItem{
		Query:  query,
		Output: output,
		Date:   time.Now().Format("Monday, January 2, 2006"),
	}

	history.HistoryList = append(history.HistoryList, historyItem)
	return SaveHistory(history, historyFile)
}
