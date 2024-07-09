package internal

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
