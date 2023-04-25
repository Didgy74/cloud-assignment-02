package utils

import (
	"time"
)

type ServerState struct {
	startTime   time.Time
	_useMocking bool

	webhookIdTracker int
	webHooks         map[int]WebhookRegistration
}

func MakeServerState() ServerState {
	out := ServerState{}
	out.webHooks = make(map[int]WebhookRegistration)
	return out
}

func (state ServerState) UseMocking() bool {
	return state._useMocking
}

func (state ServerState) UptimeInSeconds() float64 {
	return time.Since(state.startTime).Seconds()
}

func (state *ServerState) InsertWebhook(registration WebhookRegistration) int {
	output := state.webhookIdTracker
	state.webHooks[output] = registration
	state.webhookIdTracker++
	return output
}

func (state *ServerState) DeleteWebhook(id int) bool {
	_, exists := state.webHooks[id]
	if exists {
		delete(state.webHooks, id)
	}
	return exists
}

// GetAllWebhooks
//
// Makes a copy and returns it
func GetAllWebhooks(state ServerState) map[int]WebhookRegistration {
	out := make(map[int]WebhookRegistration)
	for key, value := range state.webHooks {
		out[key] = value
	}
	return out
}

type CountryItemName struct {
	Common string `json:"common"`
}

type CountryItem struct {
	Name         CountryItemName   `json:"name"`
	Languages    map[string]string `json:"languages"`
	Borders      []string          `json:"borders"`
	Cca2         string            `json:"cca2"`
	MapsInternal map[string]string `json:"maps"`
}

type WebhookRegistration struct {
	Url         string
	Event       string
	CountryCode string
	Calls       int
}

type RenewableEnergy struct {
	Entity     string  `json:"Entity"`
	Code       string  `json:"Code"`
	Year       int     `json:"Year"`
	Renewables float64 `json:"Renewables (% equivalent primary energy)"`
}

type RESTCountries struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Borders []string `json:"borders"`
	CCA3    string   `json:"cca3"`
}
