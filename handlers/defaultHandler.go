package handlers

import (
	"Assignment02/utils"
	"fmt"
	"net/http"
)

// Empty handler as default handler
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	// Ensure interpretation as HTML by client (browser)
	w.Header().Set("content-type", "text/html")

	// Offer information for redirection to paths
	output := "This service does not provide any functionality on root path level. Please use paths <a href=\"" +
		utils.CURRENT_PATH + "\">" + utils.CURRENT_PATH + "</a>, <a href=\"" + utils.HISTORY_PATH + "\">" + utils.HISTORY_PATH + "</a> or " +
		"<a href=\"" + utils.STATUS_PATH + "\">" + utils.STATUS_PATH + "</a>"

	// Write output to client
	_, err := fmt.Fprintf(w, "%v", output)

	// Deal with error if any
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}
}
