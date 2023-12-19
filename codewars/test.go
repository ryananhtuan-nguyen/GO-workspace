package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// AllFarm represents a list of farms
type AllFarm []struct {
	FarmID  string
	TowerID string
	RSSI    float64
}

// BestTower represents the best tower with its average RSSI
type BestTower struct {
	TowerID     string
	AverageRSSI float64
}

// convertCSVtoJSON converts CSV data to JSON
func convertCSVtoJSON(csvData string) (AllFarm, error) {
	reader := csv.NewReader(strings.NewReader(csvData))
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(csvData)
		return nil, err
	}

	var jsonData AllFarm
	for _, record := range records {
		rssi, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}
		jsonData = append(jsonData, struct {
			FarmID  string
			TowerID string
			RSSI    float64
		}{record[0], record[1], rssi})
	}

	return jsonData, nil
}

// fetchData fetches data from the given API URL
func fetchData(apiURL string) (AllFarm, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to fetch data from %s", apiURL)
	}

	csvData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
	}

	return convertCSVtoJSON(string(csvData))
}

// findBestTowerByFarmID finds the best tower for a given farm ID
func findBestTowerByFarmID(farmID string) (*BestTower, error) {
	apiURL := "https://api.onizmx.com/lambda/tower_stream"
	allData, err := fetchData(apiURL)
	if err != nil {
		return nil, err
	}

	var bestTower BestTower
	towers := make(map[string][]float64)

	// Filter data for the given farm ID
	for _, item := range allData {
		if item.FarmID == farmID {
			towers[item.TowerID] = append(towers[item.TowerID], item.RSSI)
		}
	}

	// Calculate average RSSI for each tower
	for towerID, rssis := range towers {
		var sum float64
		for _, rssi := range rssis {
			sum += rssi
		}
		averageRSSI := sum / float64(len(rssis))
		towers[towerID] = []float64{averageRSSI}
	}

	// Find the tower with the highest average RSSI
	var maxAverageRSSI float64
	for towerID, averageRSSI := range towers {
		if averageRSSI[0] > maxAverageRSSI {
			maxAverageRSSI = averageRSSI[0]
			bestTower.TowerID = towerID
			bestTower.AverageRSSI = maxAverageRSSI
		}
	}

	return &bestTower, nil
}

func main() {
	farmID := "yourFarmID" // Replace with the actual farm ID
	result, err := findBestTowerByFarmID(farmID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Best Tower: %+v\n", result)
}
