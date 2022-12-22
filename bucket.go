package idcloudhost

type Bucket struct {
	BillingAccountID int64  `json:"billing_account_id"`
	Name             string `json:"name"`
	NumObjects       int64  `json:"num_objects"`
	SizeBytes        int64  `json:"size_bytes"`
	CreatedAt        string `json:"created_at"`
	ModifiedAt       string `json:"modified_at"`
	IsSuspended      bool   `json:"is_suspended"`
}
