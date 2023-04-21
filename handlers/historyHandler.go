package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func HandleGetRequestForHistoricalPercentage(w http.ResponseWriter, r *http.Request) {

	// Assign successive lines of raw CSV data to fields of the created structs
	var RenewableEnergyDataset = utils.RenewableEnergyDataset

	// Convert an array of structs to JSON using marshaling functions from the encoding/json package
	jsonData, err := json.MarshalIndent(RenewableEnergyDataset, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(string(jsonData)) // Print the JSON data to the console
	fmt.Fprintf(w, "%v", string(jsonData)) // Print the JSON data to the browser
}
