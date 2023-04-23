package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func createRenewableEnergyDatasetTask1(data [][]string, country string) []utils.RenewableEnergy {
	optionalParameter := strings.Split(country, "?")

	var neighbours bool = false

	if len(optionalParameter) > 1 {
		country = optionalParameter[0]
		if optionalParameter[1] == "neighbours=true" {
			neighbours = true
		}
	}

	// Prepare empty array to populate
	var borders []string

	// Populate the array
	for _, line := range utils.RestCountriesDataset {
		if strings.EqualFold(line.Name.Common, country) {
			borders = line.Borders
		}
	}

	// convert csv lines to array of structs
	var RenewableEnergyDataset []utils.RenewableEnergy
	for i, line := range data {
		if i > 0 { // omit header line
			var rec utils.RenewableEnergy
			var err error
			var flag bool = true
			for j, field := range line {
				if j == 0 {
					if !(country == "" || strings.EqualFold(field, country)) {
						flag = false
					}
					rec.Entity = field
				} else if j == 1 {
					if neighbours == true {
						for _, line := range borders {
							if country == "" || strings.EqualFold(field, line) || flag == true {
								flag = true
								break
							}
						}
					}

					if country == "" || strings.EqualFold(field, country) || flag == true {
						flag = true
					}
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
			if flag {
				RenewableEnergyDataset = append(RenewableEnergyDataset, rec)
			}
		}
	}
	return RenewableEnergyDataset
}

// Dedicated handler for GET requests
func HandleGetRequestForCurrentPercentage(w http.ResponseWriter, r *http.Request) {

	// Get the country name from the request URL
	countryName := strings.Split(r.URL.String(), "/")[5]

	fmt.Print("Displaying every item ")
	if countryName != "" {
		fmt.Print("where Entity/Code = ", countryName)
	}
	fmt.Println()

	// Assign successive lines of raw CSV data to fields of the created structs
	var RenewableEnergyDataset = createRenewableEnergyDatasetTask1(utils.RenewableEnergyDataset, countryName)

	// Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonData, err := json.MarshalIndent(RenewableEnergyDataset, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(jsonData)) // Print the JSON data to the console
	fmt.Fprintf(w, "%v", string(jsonData)) // Print the JSON data to the browser

}
