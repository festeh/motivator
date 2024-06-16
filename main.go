package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Status struct {
	// Is Distact Me not browser extension enabled
	DmnOn bool  `json:"DmnOn"`
	Error error `json:"error"`
}

// Custom marshalling for the Status struct
func (s Status) MarshalJSON() ([]byte, error) {
	if s.Error != nil {
		return []byte(`{"DmnOn":false,"error":"` + s.Error.Error() + `"}`), nil
	}
	return []byte(`{"DmnOn":` + fmt.Sprintf("%t", s.DmnOn) + `}`), nil
}

type Extension struct {
	Manifest struct {
		Name string `json:"name"`
	} `json:"manifest"`

	State int `json:"state"`
}

type ChromePref struct {
	Extensions struct {
		Settings map[string]Extension `json:"settings"`
	} `json:"extensions"`
}

func queryExtensionEnabled() (bool, error) {
	// Get Google chrome preferences
	// Check if Distact Me is enabled
	// Set status.DmnOn to true or false
	// get username on a Linux machine
	userName := os.Getenv("USER")
	// get path to the chrome preferences file
	chromePrefPath := "/home/" + userName + "/.config/google-chrome/Default/Preferences"
	// read the chrome preferences file
	data, err := os.ReadFile(chromePrefPath)
	if err != nil {
		return false, err
	}
	var chromePref ChromePref
	// unmarshal the data into the chromePref struct
	err = json.Unmarshal(data, &chromePref)
	if err != nil {
		return false, err
	}
	for _, v := range chromePref.Extensions.Settings {
		if v.Manifest.Name == "Distract Me Not" {
			if v.State == 1 {
				return true, nil
			}
			return false, nil
		}
	}
	return false, fmt.Errorf("Distact Me Not extension not found")
}

func queryStatus() (bool, error) {
	status := true
	chromeExtPath := filepath.Join("/home", os.Getenv("USER"), ".config", "google-chrome", "Default", "Local Extension Settings", "lkmfokajfoplgdkdifijpffkjeejainc", "000003.log")
	data, err := os.ReadFile(chromeExtPath)
	if err != nil {
		return false, err
	}
	lines := strings.Split(string(data), "\n")
	lastLine := lines[len(lines)-1]
	enabledIndex := strings.LastIndex(lastLine, "Enabled")
	if enabledIndex == -1 {
		return status, nil
	}
	enabledValue := strings.TrimSpace(lastLine[enabledIndex+8:])
	if strings.ToLower(enabledValue) == "false" {
		status = false
	}
	return status, nil
}

func writeStatus(status Status) {
	str, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(str)
}

func main() {
	var status = Status{DmnOn: false}
	status.DmnOn, status.Error = queryExtensionEnabled()
	if status.DmnOn == false {
		writeStatus(status)
		os.Exit(0)
	}
	status.DmnOn, status.Error = queryStatus()
	writeStatus(status)
}
