package utils

import "time"

type ServerState struct {
	startTime   time.Time
	_useMocking bool

	webhook_id_tracker int
	webHooks           []WebhookRegistration
}

func (state ServerState) UseMocking() bool {
	return state._useMocking
}

func (state ServerState) UptimeInSeconds() float64 {
	return time.Since(state.startTime).Seconds()
}

func (state *ServerState) InsertWebhook(registration WebhookRegistration) int {
	state.webHooks = append(state.webHooks, registration)
	output := state.webhook_id_tracker
	state.webhook_id_tracker++
	return output
}

type CountryRenewableOutput struct {
	CountryName         string  `json:"name"`
	IsoCode             string  `json:"isoCode"`
	Year                int     `json:"year"`
	RenewablePercentage float32 `json:"percentage"`
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
	Url   string `json:"url"`
	Event string `json:"event"`
}

type StructTest struct {
	Entity     string  `json:"Entity"`
	Code       string  `json:"Code"`
	Year       int     `json:"Year"`
	Renewables float64 `json:"Renewables (% equivalent primary energy)"`
}
