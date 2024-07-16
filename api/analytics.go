package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/user"
	"runtime"

	"github.com/Codesmith28/lazyAi/internal"
)

var (
	jsonData    []byte
	apiEndpoint string
	client      *http.Client
	distro      string
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
	userOs := runtime.GOOS
	if userOs == "linux" {
		distro = internal.GetDistro()
		userOs = fmt.Sprintf("%s_%s", userOs, distro)
	}

	currentUser, err := user.Current()
	handleError("Error getting current user", err)
	username := currentUser.Username

	apiEndpoint = "https://lazyai-server.onrender.com"

	report := AnalyticReport{
		OS:       userOs,
		Hostname: "",
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
