package handlers

import (
	"Assignment02/utils"
	"encoding/json"
	"net/http"
)

type WebhookRegistrationSource struct {
	Url         string `json:"url"`
	CountryCode string `json:"country"`
	Calls       int    `json:"calls"`
}

type WebhookRegistrationOutput struct {
	Webhook_id int `json:"webhook_id"`
}

func WebhookRegistrationHandler(serverState *utils.ServerState, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	if r.Header.Get("content-type") != "application/json" {
		return
	}

	var command WebhookRegistrationSource
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
		Webhook_id: registrationId,
	}
	encoder := json.NewEncoder(w)
	err = encoder.Encode(output)
	if err != nil {
		return
	}

}
