package utils

import (
	"fmt"
	"path/filepath"
)

func getCSVFilepath() string {
	absPath, err := filepath.Abs(CSV_FILE_PATH)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
	}
	return absPath
}

// Absolute path to the CSV file
var CSVFilePath = getCSVFilepath()

// "Storage" for the renewable energy data. It is set in the main function
// var RenewableEnergyDataset []RenewableEnergy
var RenewableEnergyDataset [][]string

var RestCountriesDataset []RESTCountries
