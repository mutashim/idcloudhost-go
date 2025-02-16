package idcloudhost

type User struct {
	CookieID     string      `json:"cookie_id"`
	ID           int         `json:"id"`
	LastActivity string      `json:"last_activity"`
	Name         string      `json:"name"`
	Profile      interface{} `json:"profile"`
	ProfileData  Profile     `json:"profile_data"`
	State        interface{} `json:"state"`
}
