package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
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

func fetchData() ([]MyData, []string, error) {
	linksApiEndpoint := "https://api.onizmx.com/lambda/tower_stream"

	links, err := getApiLinks(linksApiEndpoint)
	if err != nil {
		fmt.Println("Error fetching links:", err)
		return nil, nil, err
	}
	var failedLinks []string

	var allData []MyData

	// Make a GET request to each API link
	for _, link := range links {
		response, err := http.Get(link)
		if err != nil {
			fmt.Println("Error making the request:", err)
			failedLinks = append(failedLinks, link)
			continue
		}
		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			fmt.Printf("Error fetching CSV from %s: Status code %d\n", link, response.StatusCode)
			failedLinks = append(failedLinks, link)
			response.Body.Close()
			continue
		}

		// Create a CSV reader
		csvReader := csv.NewReader(response.Body)
		csvReader.LazyQuotes = true

		// Skip the header row
		_, err = csvReader.Read()
		if err != nil {
			fmt.Println("Error reading CSV header:", err)
			continue
		}

		// Initialize a slice to store parsed data
		var data []MyData

		// Read the rest of the CSV records
		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error reading CSV:", err)
				continue
			}

			// Ensure the record has at least 3 columns
			if len(record) < 3 {
				fmt.Println("Invalid CSV record:", record)
				continue
			}

			// Parse CSV record and create MyData object
			myData := MyData{
				FarmID:  record[0],
				TowerID: record[1],
			}

			// Parse RSSI field as float64
			rssi, err := strconv.ParseFloat(strings.TrimSpace(record[2]), 64)
			if err != nil {
				fmt.Printf("Error parsing RSSI for record %v: %v\n", record, err)
				continue // Skip this record if RSSI parsing fails
			}
			myData.RSSI = rssi

			// Append MyData object to the slice
			data = append(data, myData)
		}

		// Access data as needed, for example:
		allData = append(allData, data...)
	}
	return allData, failedLinks, nil
}
