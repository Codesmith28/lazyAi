package internal

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	filelocation    string
	historyLocation string
	apiKey          []byte
)

func init() {
	ostype := runtime.GOOS

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
	}

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
			panic(err)
		}

		return string(apiKey)
	}
	return string(apiKey)
}
