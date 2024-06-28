package internal

type HistoryItem struct {
	Query string
	Date  string
	Model string
	SrNo  int
}

type History struct {
	HistoryList []HistoryItem
}

type Prompt struct {
	PromptString string
}

type Input struct {
	InputString string
}

type Model struct {
	SelectedModel string
}

type FormattedInput struct {
	InputString  string
	ModelName    string
	PromptString string
}

type Output struct {
	OutputString string
}
