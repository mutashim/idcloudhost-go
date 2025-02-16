package idcloudhost

type Token struct {
	BillingAccountID int64  `json:"billing_account_id"`
	ConsumerID       string `json:"consumer_id"`
	CreatedAt        string `json:"created_at"`
	Description      string `json:"description"`
	ID               int64  `json:"id"`
	KongID           string `json:"kong_id"`
	Restricted       bool   `json:"restricted"`
	Token            string `json:"token"`
	UpdatedAt        string `json:"updated_at"`
	UserID           int64  `json:"user_id"`
}
