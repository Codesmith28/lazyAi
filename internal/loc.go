package internal

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	filelocation    string
	historyLocation string
	apiKey          []byte
)

func init() {
	ostype := runtime.GOOS
	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()
	user := currentUser.Username

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if ostype == "windows" {
		filelocation = filepath.Join(homeDir, "AppData\\Local\\lazyAi\\lazy_ai_api")
		historyLocation = filepath.Join(homeDir, "AppData\\Local\\lazyAi\\history.json")
	} else if ostype == "darwin" {
		filelocation = filepath.Join(homeDir, "Library/Application Support/lazyAi/lazy_ai_api")
		historyLocation = filepath.Join(homeDir, "Library/Application Support/lazyAi/history.json")
	} else {
		filelocation = filepath.Join(homeDir, ".local/share/lazyAi/lazy_ai_api")
		historyLocation = filepath.Join(homeDir, ".local/share/lazyAi/history.json")
		distro := getDistro()
		fmt.Println(distro)
	}

	fmt.Println(hostname, user, ostype)

	apiKey, _ = os.ReadFile(filelocation)
}

func GetFileLocation() string {
	return filelocation
}

func GetHistoryLocation() string {
	return historyLocation
}

func GetAPIKey() string {
	if apiKey == nil {
		// try to read the file again
		apiKey, err := os.ReadFile(filelocation)
		if err != nil {
			return "" // this on read by api.go will be throwing out an error, moreover this will only be happening if the file is not found
		}

		return string(apiKey)
	}
	return string(apiKey)
}

func getDistro() (distro string) {
	filename := "/etc/os-release"
	file, err := os.Open(filename)
	if err != nil {
		return "unknown"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), "\"")

		switch key {
		case "NAME":
			distro = value
		}
	}

	if err := scanner.Err(); err != nil {
		return "unknown"
	}

	return distro
}
