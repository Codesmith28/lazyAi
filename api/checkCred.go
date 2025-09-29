package api

import (
	"net/http"
	"os"
	"strings"
	"time"
)

func CheckCredentials(FileLocation string, inputApiKey *string) bool {
	var apiKey []byte

	if inputApiKey != nil {
		key := *inputApiKey
		apiKey = []byte(key)
	} else {
		var err error

		apiKey, err = os.ReadFile(FileLocation)
		if err != nil {
			return false
		}
	}

	// Validate the API key by listing models
	apiKeyStr := strings.TrimSpace(string(apiKey))
	client := &http.Client{Timeout: 10 * time.Second}
	url := "https://generativelanguage.googleapis.com/v1beta/models"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false
	}

	req.Header.Set("x-goog-api-key", apiKeyStr)

	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
