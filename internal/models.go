package internal

type HistoryItem struct {
	Query string
	Date  string
	Model string
	SrNo  int
}

type History struct {
	History []HistoryItem
}

type Prompt struct {
	Prompt string
}

type Input struct {
	Input string
}

type FormattedInput struct {
	Input  string
	Model  string
	Prompt string
}

type Output struct {
	Output string
}
