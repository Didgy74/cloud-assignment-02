package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// NotificationHandler
//
// Handles all endpoints related to the notifications and webhooks.
//
// Base URL is the sub URL of the endpoint.
func NotificationHandler(serverState *utils.ServerState, baseUrl string, w http.ResponseWriter, r *http.Request) {

	// We need to extract the remaining URL after the endpoint URL.
	// This checks that the first part matches.
	if baseUrl != r.URL.Path[0:len(baseUrl)] {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	leftoverUrl := r.URL.Path[len(baseUrl):]

	if r.Method == http.MethodPost {
		// Make sure the URL doesn't have anything extra
		webhookRegistrationHandler(serverState, w, r)
	} else if r.Method == http.MethodDelete {
		webhookDeletion(serverState, leftoverUrl, w, r)
	} else if r.Method == http.MethodGet {
		webhookView(serverState, leftoverUrl, w, r)
	}
}

// WebhookRegistrationRequestBody
//
// This struct matches the schema of the request for webhook registration
type WebhookRegistrationRequestBody struct {
	Url         string `json:"url"`
	CountryCode string `json:"country"`
	Calls       int    `json:"calls"`
}

// WebhookRegistrationOutput
//
// Struct that matches the schema for endpoint for registering new webhook
type WebhookRegistrationOutput struct {
	WebhookId int `json:"webhook_id"`
}

// Handles the endpoint that creates a webhook.
func webhookRegistrationHandler(serverState *utils.ServerState, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.Header.Get("content-type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Decode the body of the request
	var command WebhookRegistrationRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&command)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	registrationId := serverState.InsertWebhook(toWebhookRegistration(command))

	output := WebhookRegistrationOutput{WebhookId: registrationId}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Handles the endpoint that deletes a webhook
func webhookDeletion(serverState *utils.ServerState, leftoverUrl string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, httpCode, err := parseLeftoverUrlAsInt(leftoverUrl)
	if err != nil {
		w.WriteHeader(httpCode)
		_, _ = fmt.Fprint(w, err.Error())
		return
	}
	if httpCode != http.StatusOK {
		w.WriteHeader(httpCode)
		return
	}
	if id == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deletionSuccess := serverState.DeleteWebhook(*id)
	if !deletionSuccess {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// ViewWebhookRegistrationOutput
//
// Struct that matches the schema of endpoint for viewing webhooks.
type ViewWebhookRegistrationOutput struct {
	WebhookId   int    `json:"webhook_id"`
	Url         string `json:"url"`
	CountryCode string `json:"country"`
	Calls       int    `json:"calls"`
}

// Handles the endpoint for viewing webhooks
func webhookView(serverState *utils.ServerState, leftoverUrl string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	id, httpCode, err := parseLeftoverUrlAsInt(leftoverUrl)
	if err != nil {
		w.WriteHeader(httpCode)
		_, _ = fmt.Fprint(w, err.Error())
		return
	}
	if httpCode != http.StatusOK {
		w.WriteHeader(httpCode)
		return
	}

	allWebhooks := utils.GetAllWebhooks(*serverState)

	output := combineWebhooksToWebhookViewOutput(allWebhooks, id)

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Parses the end of the URL as an integer.
//
// First return-value returns the integer if there is one,
// nil if there is nothing at the end.
// Second return is the http-code
// Third return is the error.
func parseLeftoverUrlAsInt(leftoverUrl string) (*int, int, error) {
	if leftoverUrl == "" {
		return nil, http.StatusOK, nil
	}
	// Interpret it as an int
	integer, err := strconv.Atoi(leftoverUrl)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	return &integer, http.StatusOK, nil
}

// Combines all webhooks into a list of webhooks to view
func combineWebhooksToWebhookViewOutput(
	allWebhooks map[int]utils.WebhookRegistration,
	wantedId *int) []ViewWebhookRegistrationOutput {

	var output []ViewWebhookRegistrationOutput

	if wantedId != nil {
		// Find the one id and add it
		id := *wantedId
		value, exists := allWebhooks[*wantedId]
		if exists {
			output = append(output, toWebhookRegistrationOutput(id, value))
		}
	} else {
		// Add all to the output if no id present
		for key, value := range allWebhooks {
			output = append(output, toWebhookRegistrationOutput(key, value))
		}
	}

	return output
}

func toWebhookRegistrationOutput(id int, value utils.WebhookRegistration) ViewWebhookRegistrationOutput {
	return ViewWebhookRegistrationOutput{
		WebhookId:   id,
		Url:         value.Url,
		CountryCode: value.CountryCode,
		Calls:       value.Calls,
	}
}

func toWebhookRegistration(value WebhookRegistrationRequestBody) utils.WebhookRegistration {
	return utils.WebhookRegistration{
		Url:         value.Url,
		CountryCode: value.CountryCode,
		Calls:       value.Calls}
}
