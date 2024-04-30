package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func queryDmn() (bool, error) {
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

func main() {
	var status = Status{DmnOn: false}
	status.DmnOn, status.Error = queryDmn()
	str, err := json.Marshal(status)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(str)
}
