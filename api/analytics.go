package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/user"

	"github.com/joho/godotenv"
)

var (
	jsonData    []byte
	apiEndpoint string
	client      *http.Client
)

type AnalyticReport struct {
	OS       string `json:"operatingsystem"`
	Hostname string `json:"hostname"`
	Username string `json:"username"`
}

func handleError(message string, err error) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
}

func init() {
	err := godotenv.Load()
	handleError("Error loading .env file", err)

	osType := os.Getenv("OS")
	hostname, err := os.Hostname()
	handleError("Error getting hostname", err)

	currentUser, err := user.Current()
	handleError("Error getting current user", err)
	username := currentUser.Username

	apiEndpoint = os.Getenv("ANALYTICS_API_ENDPOINT")
	if apiEndpoint == "" {
		handleError("ANALYTICS_API_ENDPOINT not set in .env file", fmt.Errorf("empty endpoint"))
	}

	report := AnalyticReport{
		OS:       osType,
		Hostname: hostname,
		Username: username,
	}

	jsonData, err = json.Marshal(report)
	handleError("Error marshalling analytic report", err)
}

func SendAnalyticReport() {
	request, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonData))
	handleError("Error creating request", err)

	request.Header.Set("Content-Type", "application/json")
	client = &http.Client{}

	resp, err := client.Do(request)
	handleError("Error sending analytic report", err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		handleError("Error sending analytic report", fmt.Errorf("status code %d", resp.StatusCode))
	}
	log.Println("Analytic report sent successfully")
}
