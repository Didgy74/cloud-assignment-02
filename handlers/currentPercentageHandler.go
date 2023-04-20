package handlers

import (
  "encoding/csv"
  "fmt"
  "os"
  "net/http"
  "strings"
)


// Dedicated handler for GET requests
func HandleGetRequestForCurrentPercentage(w http.ResponseWriter, r *http.Request) {

	// Get the country name from the request URL
	//URLParts := strings.Split(r.URL.String(), "/")
	//countryName := URLParts[??]


	// Open the CSV file
	fd, err := os.Open("renewable-share-energy.csv")
	if err != nil {
		fmt.Println("Error opening CSV file.")
		fmt.Println(err)
	}
	fmt.Println("Successfully opened the CSV file.")

	// Remember to close the file at the end of the program
	defer fd.Close()

	// Read the CSV file
	fileReader := csv.NewReader(fd)
	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file.")
		fmt.Println(err)
	}
	fmt.Println(records)
}