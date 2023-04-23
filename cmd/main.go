package main

import (
	"Assignment02/handlers"
	"Assignment02/utils"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	serverState := utils.ServerState{}

	// Open the CSV file
	fd, err := os.Open(utils.CSVFilePath)
	if err != nil {
		fmt.Println("Error opening CSV file.")
		fmt.Println(err)
	}
	fmt.Println("Successfully opened the CSV file.")
	defer fd.Close() // Remember to close the file at the end of the program

	// Read the CSV file
	fileReader := csv.NewReader(fd)
	data, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file.")
		fmt.Println(err)
	}

	//fmt.Println(data) // Print the CSV data to the console
	//fmt.Fprintf(w, "%v", data) // Print the CSV data to the browser

	// Store the CSV data in a variable
	utils.RenewableEnergyDataset = data

	// Extract PORT variable from the environment variables0
	port := os.Getenv("PORT")

	// Override port with default port if not provided (e.g. local deployment)
	if port == "" {
		log.Println("$PORT has not been set. Default: 8080")
		port = "8080"
	}

	http.HandleFunc(utils.DEFAULT_PATH, handlers.DefaultHandler)
	http.HandleFunc(utils.CURRENT_PATH, handlers.HandleGetRequestForCurrentPercentage)
	http.HandleFunc(utils.HISTORY_PATH, handlers.HandleGetRequestForHistoricalPercentage)

	// Add the
	http.HandleFunc(
		utils.NOTIFICATIONS_PATH,
		func(w http.ResponseWriter, r *http.Request) {
			handlers.WebhookRegistrationHandler(&serverState, w, r)
		})

	log.Println("Starting server on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

/*
func FindCountrySearchFromURL(r url.URL) (string, error) {
	split := strings.Split(r.Path, "/")
	if len(split) > 2 {
		return "", errors.New("URL is not long enough.")
	} else if len(split) == 2 {
		return split[1], nil
	}

	return "", nil
}

func SearchCountryName(country string) (*utils.CountryItem, int, error) {
	country = strings.ReplaceAll(country, " ", "%20")

	countryRequest, err := http.NewRequest(
		http.MethodGet,
		"https://restcountries.com/v3.1/name/"+country,
		nil)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	response, err := http.DefaultClient.Do(countryRequest)
	if err != nil {
		return nil, http.StatusFailedDependency, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, response.StatusCode, errors.New("Response is not okay.")
	}
	if response.Header.Get("content-type") != "application/json" {
		return nil, http.StatusFailedDependency, err
	}
	var countryItems []utils.CountryItem
	err = json.NewDecoder(response.Body).Decode(&countryItems)
	if err != nil {
		return nil, http.StatusFailedDependency, err
	}
	if len(countryItems) == 0 {
		return nil, http.StatusOK, nil
	} else {
		return &countryItems[0], http.StatusOK, nil
	}
}

func yo(serverState *ServerState, w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Grab country search code
	countrySearch, err := FindCountrySearchFromURL(*r.URL)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if countrySearch != "" {
		country, httpCode, err := SearchCountryName(countrySearch)
		if err != nil {

		}
	}

	returnVal := utils.CountryRenewableOutput{}

	w.Header().Set("content-type", "application/json")
	err = json.NewEncoder(w).Encode(returnVal)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
*/
