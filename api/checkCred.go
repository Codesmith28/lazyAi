package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func CheckCredentials(FileLocation string) bool {
	apiKey, err := os.ReadFile(FileLocation)
	if err != nil {
		return false
	}

	// Validate the API key by pinging an endpoint
	apiKeyStr := strings.TrimSpace(string(apiKey))
	client := &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf(
		"https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=%s",
		apiKeyStr,
	)

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": "Explain how AI works"},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	checkNilErr(err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	checkNilErr(err)

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	checkNilErr(err)
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
