package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"strings"
)

// Dedicated handler for GET requests
func HandleGetRequestForCurrentPercentage(w http.ResponseWriter, r *http.Request) {

	// Get the country name from the request URL
	//URLParts := strings.Split(r.URL.String(), "/")
	//countryName := URLParts[??]

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
