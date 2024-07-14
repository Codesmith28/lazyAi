package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/user"

	"github.com/joho/godotenv"
)

var (
	jsonData    []byte
	apiEndpoint string
	request     *http.Request
	client      *http.Client
)

type AnalyticReport struct {
	OS       string `json:"operatingsystem"`
	Hostname string `json:"hostname"`
	Username string `json:"username"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	osType := os.Getenv("OS")
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error getting hostname: ", err)
	}

	currentUser, err := user.Current()
	if err != nil {
		log.Fatal("Error getting current user: ", err)
	}
	username := currentUser.Username

	apiEndpoint = os.Getenv("ANALYTICS_API_ENDPOINT")
	if apiEndpoint == "" {
		log.Fatal("ANALYTICS_API_ENDPOINT not set in .env file")
	}

	report := AnalyticReport{
		OS:       osType,
		Hostname: hostname,
		Username: username,
	}

	jsonData, err = json.Marshal(report)
	if err != nil {
		log.Fatal("Error marshalling analytic report: ", err)
	}

	request, err = http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	request.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
}

func SendAnalyticReport() {
	resp, err := client.Do(request)
	if err != nil {
		log.Println("Error sending analytic report: ", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error sending analytic report: status code %d", resp.StatusCode)
		return
	}

	log.Println("Analytic report sent successfully")
}
