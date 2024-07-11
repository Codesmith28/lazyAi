package internal

import "github.com/gdamore/tcell/v2"

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
	HistoryMap  map[string]HistoryItem
	HistoryList []HistoryItem
}

type Prompt struct {
	PromptString string
	Model        string
}

type Input struct {
	InputString string
}

type Model struct {
	SelectedModel string
}

type Output struct {
	OutputString string
}

type Theme struct {
	ActiveBorderCol   tcell.Color
	InactiveBorderCol tcell.Color
	PrimaryTextCol    tcell.Color
	SecondaryTextCol  tcell.Color
	TertiaryTextCol   tcell.Color
	TitleCol          tcell.Color
}
