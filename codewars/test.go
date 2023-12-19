package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Custom struct for CSV record
type MyData struct {
	FarmID  string
	TowerID string
	RSSI    float64
}

func getApiLinks(s string) ([]string, error) {
	// GET request to fetch api links
	response, err := http.Get(s)
	if err != nil {
		// Return an error if failed to fetch
		return nil, fmt.Errorf("failed to fetch API link: %v", err)
	}
	defer response.Body.Close()

	// Check the HTTP status code
	if response.StatusCode != http.StatusOK {
		// Return an error if the status code indicates a problem
		return nil, fmt.Errorf("failed to fetch API link, status code: %d", response.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// Return an error if the body is invalid
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse JSON array of URLs
	var links []string
	err = json.Unmarshal(body, &links)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON array: %v", err)
	}

	return links, nil
}
