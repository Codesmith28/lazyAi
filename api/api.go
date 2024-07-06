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
		panic(err)
	}
}

func SendPrompt(promptString string, modelName string, inputString string) (string, error) {
	ctx := context.Background()
	err := godotenv.Load()
	checkNilErr(err)

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	checkNilErr(err)
	defer client.Close()

	model := client.GenerativeModel(modelName)

	fullPrompt := fmt.Sprintf("%s\n\nQuestion: %s", promptString, inputString)

	if promptString == "" {
		fullPrompt = fmt.Sprintf("Question: %s", inputString)
	}

	log.Printf("Full prompt: %s", fullPrompt)

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
