package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func NotificationHandler(serverState *utils.ServerState, w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Make sure the URL doesn't have anything extra

		webhookRegistrationHandler(serverState, w, r)
	} else if r.Method == http.MethodDelete {
		webhookDeletion(serverState, w, r)
	}
}

type WebhookRegistrationBody struct {
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

	var command WebhookRegistrationBody
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
	encoder := json.NewEncoder(w)
	err = encoder.Encode(output)
	if err != nil {
		return
	}

}

func webhookDeletion(serverState *utils.ServerState, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		return
	}

	// Extract the ID from the path
	remainingPath := r.URL.Path[len(utils.NOTIFICATIONS_PATH):]
	if remainingPath == "" {
		return
	}

	// Interpret it as an int
	id, err := strconv.Atoi(remainingPath)
	if err != nil {
		return
	}

	deletionSuccess := serverState.DeleteWebhook(id)
	if !deletionSuccess {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
