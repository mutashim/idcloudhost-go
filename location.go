package idcloudhost

type Location struct {
	DisplayName string `json:"display_name"`
	IsDefault   bool   `json:"is_default"`
	IsPreferred bool   `json:"is_preferred"`
	Description string `json:"description"`
	OrderNr     int64  `json:"order_nr"`
	Slug        string `json:"slug"`
	CountryCode string `json:"country_code"`
}
