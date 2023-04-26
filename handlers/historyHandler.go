package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// This function is the entry point for historical renewable energy percentage statistics
func HandleGetRequestForHistoricalPercentage(w http.ResponseWriter, r *http.Request) {

	//only accept GET method requests
	switch r.Method {
	case http.MethodGet:
		getHandler(w, r)
	default:
		http.Error(w, "REST Method "+r.Method+" not supported. Currently only "+http.MethodGet+
			" is supported.", http.StatusNotImplemented)
		return
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	//Setting the header
	w.Header().Set("content-type", "application/json")

	// Retrieve the dataset variable
	data := parse(utils.RenewableEnergyDataset)

	//Retrieve the URL path
	parts := strings.Split(r.URL.Path, "/")

	//Check the path length and return an error in case there are too many parts
	if len(parts) != 6 {
		http.Error(w, "Malformed URL. Expected format: ... "+utils.HISTORY_PATH+"<country code>?begin=<year>&end=<year>", http.StatusBadRequest)
		return
	}

	// In case of no country code was provided
	if parts[5] == "" {

		//check if any time boundaries were specified
		if r.URL.RawQuery == "" { //In case no time boundaries were not specified
			//Return the information for all years
			meanHistoricalPercentage(w, data, 1965, 2021)
		}
	} else { //in case a country code and time boundaries are a part of the URL
		if len(parts[5]) != 3 { // Check for the length of the country code and return an error if the code is not valid
			http.Error(w, "Please enter a valid 3-letter country code.", http.StatusBadRequest)
			return
		}
		//Check if any time parameters were specified
		if r.URL.RawQuery == "" { //In case no time boundaries were not specified
			historyPercentageByCountry(w, data, strings.ToUpper(parts[5]), 1965, 2021)
		}
	}
	//Return OK status code when finished
	http.Error(w, "OK", http.StatusOK)

}

// This function calculates the mean renewable energy percentage in a specific time frame and writes it to the response writer
func meanHistoricalPercentage(w http.ResponseWriter, data []utils.RenewableEnergy, start int, finish int) {

	//Instantiate the response variable to be returned to the client
	var response []utils.MeanRenewableEnergy

	//A variable to keep track of the sum of percentages for each country
	var sum float64 = 0

	//A variable to keep track of the number of percentages in the sum
	num := 0

	//Loop over the dataset calculating the mean percentage within the time boundaries for each country
	name := data[0].Entity
	for _, element := range data {
		// Reset the variables and register the results whenever the loop encounters a new country
		if element.Entity != name {

			if num == 0 {
				log.Println(element)
				num += 1
			}
			temp := utils.MeanRenewableEnergy{Entity: name, Code: element.Code, Renewables: sum / float64(num)}

			response = append(response, temp)
			name = element.Entity
			sum = 0
			num = 0
		}

		//Check that the year is within the time limits before updating the variables
		if element.Year >= start && element.Year <= finish {
			sum += element.Renewables
			num += 1
		}
	}
	//Encode and send the response to the client or return an error if the encoding fails
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error occurred while encoding coding the response "+err.Error(), http.StatusInternalServerError)
		return
	}

}

// This function retrieves the renewable energy percentages for the specifies country and time limits and writes it to the response writer
func historyPercentageByCountry(w http.ResponseWriter, data []utils.RenewableEnergy, countryCode string, start int, finish int) {

	//Instantiate the response variable to be returned to the client
	var response []utils.RenewableEnergy

	//Loop over the data set adding all the object with a matching the 3-letter code and falling withing the time limits
	for index, element := range data {
		country := element.Entity
		if element.Code == countryCode { // Check if the current element matches the 3-letter code
			i := index
			for data[i].Entity == country { //Check that loop hasn't reached a new country
				if data[i].Year >= start && data[i].Year <= finish { //Check that the element falls withing the time limits

					//Add the element to the response slice
					response = append(response, data[i])
				}
				i++
			}
			break

		}
	}
	//Handling the case where no country with the specified code was found
	if response == nil || len(response) == 0 {
		http.Error(w, "No country matching the code was found, please make sure the code you entered is valid.", http.StatusNotFound)
		return
	}
	//Encode and send the response to the client or return an error if the encoding fails
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error occurred while decoding the response", http.StatusInternalServerError)
		return
	}

}

func parse(dataset [][]string) []utils.RenewableEnergy {
	// convert csv lines to array of structs
	var RenewableEnergyDataset []utils.RenewableEnergy
	for i, line := range dataset {
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
