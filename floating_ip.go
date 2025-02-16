package idcloudhost

type FloatingIP struct {
	ID                     int64
	Address                string
	UserID                 int64
	BillingAccountID       int64
	Type                   string
	NetworkID              int64
	Name                   string
	Enabled                bool
	CreatedAt              string
	UpdatedAt              string
	IsDeleted              bool
	IsVirtual              bool
	AssignedTo             string
	AssignedToResourceType string
	AssignedToPrivateIP    string
}
