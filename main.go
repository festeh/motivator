package main

import (
	"fmt"
	"os"
)

type Status struct {
	// Is Distact Me not browser extension enabled
	DmnOn bool `json:"DmnOn"`
}

type ChromePref struct {
	Extensions struct {
		Settings map[string]interface{} `json:"settings"`
	} `json:"extensions"`
}

func queryDmn() (bool, error) {
	// Get Google chrome preferences
	// Check if Distact Me is enabled
	// Set status.DmnOn to true or false
	// get username on a Linux machine
	userName := os.Getenv("USER")
	// get path to the chrome preferences file
	chromePrefPath := "/home/" + userName + "/.config/google-chrome/Default/Preferences"
	// read the chrome preferences file
	chromePref, err := os.ReadFile(chromePrefPath)
	if err != nil {
		return false, err
	}
	fmt.Println(string(chromePref))
	// unmarshal the chrome preferences file

	return true, nil
}

func main() {
	// var status = Status{IsDmnOn: false}
	queryDmn()
}
