package utils

type ServerState struct {
	startTime   time.Time
	_useMocking bool
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