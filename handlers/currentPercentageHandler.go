package handlers

import (
	"Assignment02/utils"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	//"strings"
)

func createRenewableEnergyDataset(data [][]string) []utils.StructTest {
	// convert csv lines to array of structs
	var RenewableEnergyDataset []utils.StructTest
	for i, line := range data {
		if i > 0 { // omit header line
			var rec utils.StructTest
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

// Dedicated handler for GET requests
func HandleGetRequestForCurrentPercentage(w http.ResponseWriter, r *http.Request) {

	// Get the country name from the request URL
	//URLParts := strings.Split(r.URL.String(), "/")
	//countryName := URLParts[??]

	// Open the CSV file
	fd, err := os.Open(utils.CsvFilePath)
	if err != nil {
		fmt.Println("Error opening CSV file.")
		fmt.Println(err)
	}
	fmt.Println("Successfully opened the CSV file.")

	// Remember to close the file at the end of the program
	defer fd.Close()

	// Read the CSV file
	fileReader := csv.NewReader(fd)
	data, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file.")
		fmt.Println(err)
	}
	//fmt.Println(data) // Print the CSV data to the console
	//fmt.Fprintf(w, "%v", data) // Print the CSV data to the browser

	// Assign successive lines of raw CSV data to fields of the created structs
	RenewableEnergyDataset := createRenewableEnergyDataset(data)

	// Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonData, err := json.MarshalIndent(RenewableEnergyDataset, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(jsonData)) // Print the JSON data to the console
	fmt.Fprintf(w, "%v", string(jsonData)) // Print the JSON data to the browser

}
