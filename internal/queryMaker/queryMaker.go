package querymaker

import "github.com/Codesmith28/cheatScript/internal"

type Query = internal.Query

func MakeQuery(inputString string, promptString string, selectedModel string) Query {
	return Query{
		InputString:   inputString,
		PromptString:  promptString,
		SelectedModel: selectedModel,
	}
}
