package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationResult struct {
	Ip      string `json:"ip"`
	Country string `json:"country"`
	Cc      string `json:"cc"`
}

func GetMyLocation() (LocationResult, error) {
	var locationResponse LocationResult

	resp, err := http.Get("https://api.myip.com/")
	if err != nil {
		return locationResponse, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationResponse, err
	}

	if err := json.Unmarshal(body, &locationResponse); err != nil {
		return locationResponse, err
	}
	return locationResponse, nil
}
