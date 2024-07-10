package api

func CheckCredentials(apiInput string) bool {

	if apiInput == "" {
		return false
	}

	_, err := SendPrompt("", "gemini-1.5-flash", "Hello, World!", &apiInput)

	return err == nil
}
