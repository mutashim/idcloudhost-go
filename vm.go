package idcloudhost

type VM struct {
	Backup         bool
	BillingAccount int64
	CreatedAt      string
	Description    string
	Hostname       string
	ID             int64
	MAC            string
	Memory         int64
	Name           string
	OSName         string
	OSVersion      string
	PrivateIPv4    string
	PublicIPv6     string
	Status         string
	Storage        []Storage
	Tags           interface{}
	UpdatedAt      string
	UserID         int64
	Username       string
	UUID           string
	VCPU           int64
}
