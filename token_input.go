package idcloudhost

type TokenInput struct {
	BillingAccountID int64  `form:"billing_account_id"`
	Description      string `form:"description"`
	Restricted       bool   `form:"restricted"`
}
