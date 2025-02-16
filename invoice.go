package idcloudhost

import "encoding/json"

type Invoice struct {
	AccountSnapshot    json.RawMessage
	BillingAccountID   int
	Created            int64
	DiscountPercentage int
	DueDate            int
	ID                 int
	PaddedID           string
	PeriodEnd          int
	PeriodStart        int
	RecordsList        []InvoiceRecord
}

type InvoiceRecord struct {
	Amount       float64
	Created      int
	Description  string
	ID           int
	InvoiceID    int
	ItemPrice    float64
	LocationSlug string
	Name         string
	Qty          int
	QtyUnit      string
}
