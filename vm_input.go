package idcloudhost

type VMInput struct {
	Backup           bool   `form:"backup"`
	BillingAccountID int64  `form:"billing_account_id"`
	Description      string `form:"description"`
	Name             string `form:"name"`
	OSName           string `form:"os_name"`
	OSVersion        string `form:"os_version"`
	Password         string `form:"password"`
	PublicKey        string `form:"public_key"`
	RAM              int64  `form:"ram"`
	SourceReplica    string `form:"source_replica"`
	SourceUUID       string `form:"source_uuid"`
	Username         string `form:"username"`
	VCPU             int64  `form:"vcpu"`
	ReservePublicIP  bool   `form:"reserve_public_ip"`
	NetworkUUID      string `form:"network_uuid"`
	CloudInit        string `form:"cloud_init"`
	DiskUUID         string `form:"disk_uuid"`
	Disks            string `form:"disks"`
}
