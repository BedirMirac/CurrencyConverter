package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Currency struct { // key names for exchangerate-api // If you use something else like fixer.io it could be different
	Bases string             `json:"base_code"`
	Date  string             `json:"time_last_update_utc"`
	Rates map[string]float64 `json:"conversion_rates"`
}

var Data Currency

// Fetching rates
func SaveData() string {
	err := godotenv.Load()
	if err != nil {
		return fmt.Sprintln("Error loading .env file")
	}

	apiURL := os.Getenv("API_URL")
	fileName := os.Getenv("OUTPUT_FILE")

	resp, err := http.Get(apiURL)

	if err != nil {
		return fmt.Sprintln("Failed to fetch data")
	}
	defer resp.Body.Close()

	rawData, _ := io.ReadAll(resp.Body)

	json.Unmarshal(rawData, &Data)

	fileData, _ := json.MarshalIndent(Data, "", "    ")

	err = os.WriteFile(fileName, fileData, 0644)
	if err != nil {
		return fmt.Sprintln("Failed to save file")
	}

	return " "
}
