package api

import (
	"context"
	"fmt"

	"github.com/Codesmith28/cheatScript/internal"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func SendPrompt(promptString, modelName, inputString string, apiKeyValidate *string) (string, error) {
	ctx := context.Background()

	var apiKey string
	if apiKeyValidate != nil {
		apiKey = *apiKeyValidate
	} else {
		apiKey = internal.GetAPIKey()
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))

	if err != nil {
		return "", err
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
	if err != nil {
		return "", err
	} else if apiKeyValidate != nil {
		return "", nil
	}

	if resp != nil && len(resp.Candidates) > 0 {
		promptAns, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
		if !ok {
			return "", fmt.Errorf("unexpected response format")
		}
		return string(promptAns), nil
	}

	return "", fmt.Errorf("no response generated")
}
