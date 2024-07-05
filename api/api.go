package api

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func checkNilErr(err error) (string, error) {
	if err != nil {
		return "", err
	}
	return "", nil
}

func SendPrompt(prompt string, modelName string) (string, error) {
	ctx := context.Background()
	err := godotenv.Load()
	checkNilErr(err)

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	checkNilErr(err)
	defer client.Close()

	model := client.GenerativeModel(modelName)
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	checkNilErr(err)

	if resp != nil && len(resp.Candidates) > 0 {
		promptAns, _ := resp.Candidates[0].Content.Parts[0].(genai.Text)
		return string(promptAns), nil
	}

	return "", nil
}
