package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func createRenewableEnergyDatasetTask2(data [][]string) []utils.RenewableEnergy {
	// convert csv lines to array of structs
	var RenewableEnergyDataset []utils.RenewableEnergy
	for i, line := range data {
		if i > 0 { // omit header line
			var rec utils.RenewableEnergy
			var err error
			for j, field := range line {
				if j == 0 {
					rec.Entity = field
				} else if j == 1 {
					rec.Code = field
				} else if j == 2 {
					rec.Year, err = strconv.Atoi(field)
				} else if j == 3 {
					rec.Renewables, err = strconv.ParseFloat(field, 64)
					if err != nil {
						continue
					}
				}
			}
			RenewableEnergyDataset = append(RenewableEnergyDataset, rec)
		}
	}
	return RenewableEnergyDataset
}

func HandleGetRequestForHistoricalPercentage(w http.ResponseWriter, r *http.Request) {

	// Assign successive lines of raw CSV data to fields of the created structs
	var RenewableEnergyDataset = createRenewableEnergyDatasetTask2(utils.RenewableEnergyDataset)

	// Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonData, err := json.MarshalIndent(RenewableEnergyDataset, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(jsonData)) // Print the JSON data to the console
	fmt.Fprintf(w, "%v", string(jsonData)) // Print the JSON data to the browser
}
