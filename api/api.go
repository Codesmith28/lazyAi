package api

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/Codesmith28/lazyAi/internal"
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
		return fmt.Sprintf("## Failed to generate because:\n %s \n \n for more info, **please visit**: https://cloud.google.com/vertex-ai/generative-ai/docs/multimodal/configure-safety-attributes", err), nil
	}

	if resp == nil || len(resp.Candidates) == 0 {
		return "Failed to generate content", nil
	}

	promptAns, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "Failed to generate content because of unexpected response format from API.", nil
	}

	return string(promptAns), nil
}
