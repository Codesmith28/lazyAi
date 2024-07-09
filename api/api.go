package api

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SendPrompt(promptString, modelName, inputString, FileLocation string) (string, error) {
	ctx := context.Background()
	err := godotenv.Load()
	checkNilErr(err)

	apiKeyFile := FileLocation
	apiKey, err := os.ReadFile(apiKeyFile)
	checkNilErr(err)

	client, err := genai.NewClient(ctx, option.WithAPIKey(string(apiKey)))
	checkNilErr(err)
	defer client.Close()

	model := client.GenerativeModel(modelName)

	var fullPrompt string

	if promptString != "" {
		fullPrompt = fmt.Sprintf("Context: %s \n\n Question: %s", promptString, inputString)
	} else {
		fullPrompt = fmt.Sprintf("Question: %s ", inputString)
	}

	resp, err := model.GenerateContent(ctx, genai.Text(fullPrompt))
	checkNilErr(err)

	if resp != nil && len(resp.Candidates) > 0 {
		promptAns, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
		if !ok {
			return "", fmt.Errorf("unexpected response format")
		}
		return string(promptAns), nil
	}

	return "", fmt.Errorf("no response generated")
}
