package api

// analytics helps us to keep a track of different users and their usage of the application

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
	osType      string
	hostname    string
	username    string
	jsonData    []byte
	apiEndpoint string
)

type AnalyticReport struct {
	OS       string `json:"operatingsystem"`
	Hostname string `json:"hostname"`
	Username string `json:"username"`
}

func init() {
	osType = os.Getenv("OS")
	hostname, _ = os.Hostname()
	currentUser, _ := user.Current()
	username = currentUser.Username

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error sending analytic report")
		return
	}

	apiEndpoint = os.Getenv("ANALYTICS_API_ENDPOINT")

	report := AnalyticReport{
		OS:       osType,
		Hostname: hostname,
		Username: username,
	}

	jsonData, err = json.Marshal(report)
	if err != nil {
		log.Fatal("Error sending analytic report")
		return
	}
}

func SendAnalyticReport() {
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatal("Error sending analytic report")
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending analytic report")
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error sending analytic report")
	}

	log.Println("Analytic report sent successfully")
}
