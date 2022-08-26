package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fatih/color"
)

var spinnerdexAPI = "http://localhost:5173/api"

type Spinner struct {
	Name    string `json:"name"`
	Twitter string `json:"twitter,omitempty"`
	Youtube string `json:"youtube,omitempty"`
	Board   string `json:"board,omitempty"`
}

type Error struct {
	Code    int
	Message string
}

func AddSpinner(spinner Spinner, apiKey string) {
	spinnerKey := strings.ToLower(spinner.Name)
	url := spinnerdexAPI + "/spinners/" + spinnerKey
	body, err := json.Marshal(spinner)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 409 {
		color.Red("A spinner with this name already exists. Try updating it instead.")
	} else if resp.StatusCode == 401 {
		color.Red("You must provide a valid API key to add spinners to the database")
	} else if resp.StatusCode == 403 {
		color.Red("You do not have permission to add spinners to the database")
	} else if resp.StatusCode == 500 {
		color.Red("Something went wrong. Please try again later")
	} else {
		color.Green("Spinner added successfully")
	}
	defer resp.Body.Close()
}
