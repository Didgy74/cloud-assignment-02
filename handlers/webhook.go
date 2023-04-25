package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

// NotificationHandler
// Base URL is the sub URL of the endpoint.
func NotificationHandler(serverState *utils.ServerState, baseUrl string, w http.ResponseWriter, r *http.Request) {

	if baseUrl != r.URL.Path[0:len(baseUrl)] {
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

type WebhookRegistrationOutput struct {
	WebhookId int `json:"webhook_id"`
}

func webhookRegistrationHandler(serverState *utils.ServerState, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	if r.Header.Get("content-type") != "application/json" {
		return
	}

	var command WebhookRegistrationRequestBody
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&command)
	if err != nil {
		return
	}

	registration := utils.WebhookRegistration{
		Url:   command.Url,
		Event: command.CountryCode,
	}
	registrationId := serverState.InsertWebhook(registration)

	output := WebhookRegistrationOutput{
		WebhookId: registrationId,
	}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(output)
	if err != nil {
		return
	}
}

func webhookDeletion(serverState *utils.ServerState, leftoverUrl string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		return
	}

	id, httpCode, err := parseLeftoverUrlAsInt(leftoverUrl)
	if err != nil {
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

type ViewWebhookRegistrationOutput struct {
	WebhookId   int    `json:"webhook_id"`
	Url         string `json:"url"`
	CountryCode string `json:"country"`
	Calls       int    `json:"calls"`
}

func webhookView(serverState *utils.ServerState, leftoverUrl string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}

	id, httpCode, err := parseLeftoverUrlAsInt(leftoverUrl)
	if err != nil {
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

func combineWebhooksToWebhookViewOutput(
	allWebhooks map[int]utils.WebhookRegistration,
	wantedId *int) []ViewWebhookRegistrationOutput {

	var output []ViewWebhookRegistrationOutput

	if len(allWebhooks) != 0 {
		// Add all to the output if no id present
		if wantedId == nil {
			for key, value := range allWebhooks {
				output = append(output, ViewWebhookRegistrationOutput{
					WebhookId:   key,
					Url:         value.Url,
					CountryCode: value.CountryCode,
					Calls:       0,
				})
			}
		} else {
			// Find the one id
			value, exists := allWebhooks[*wantedId]
			if exists {
				output = append(output, ViewWebhookRegistrationOutput{
					WebhookId:   *wantedId,
					Url:         value.Url,
					CountryCode: value.CountryCode,
					Calls:       0,
				})
			}
		}
	}

	return output
}
