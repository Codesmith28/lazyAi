package api

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func SendPrompt(prompt string) (string, error) {
	ctx := context.Background()
	err := godotenv.Load()

	if err != nil {
		return "", err
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if resp != nil && len(resp.Candidates) > 0 {
		promptAns, _ := resp.Candidates[0].Content.Parts[0].(genai.Text)
		return string(promptAns), nil
	}

	return "", nil
}
