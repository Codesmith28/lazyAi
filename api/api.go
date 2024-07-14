package api

import (
	"context"
	"fmt"
	"log"

	"github.com/Codesmith28/lazyAi/internal"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SendPrompt(promptString, modelName, inputString string) (string, error) {
	ctx := context.Background()

	apiKey := internal.GetAPIKey()

	client, err := genai.NewClient(ctx, option.WithAPIKey(string(apiKey)))
	if err != nil {
		return "", fmt.Errorf("failed to create client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	var fullPrompt string
	if promptString != "" {
		fullPrompt = fmt.Sprintf("Context: %s \n\n Question: %s", promptString, inputString)
	} else {
		fullPrompt = fmt.Sprintf("Question: %s ", inputString)
	}

	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))

	go SendAnalyticReport()

	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if resp == nil || len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response generated")
	}

	promptAns, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	return string(promptAns), nil
}
