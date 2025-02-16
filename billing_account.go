package idcloudhost

type BillingAccount struct {
	AdditionalData            string        `json:"additional_data"`
	AddressLine1              string        `json:"address_line1"`
	AllowDebt                 bool          `json:"allow_debt"`
	CanPay                    bool          `json:"can_pay"`
	City                      string        `json:"city"`
	CompanyName               string        `json:"company_name"`
	CompanyRegCode            string        `json:"company_reg_code"`
	CompanyVatNumber          string        `json:"company_vat_number"`
	Country                   string        `json:"country"`
	County                    string        `json:"county"`
	Created                   int64         `json:"created"`
	CreditAmount              float64       `json:"credit_amount"`
	DiscountPercentage        float64       `json:"discount_percentage"`
	Email                     string        `json:"email"`
	ID                        int64         `json:"id"`
	IsActive                  bool          `json:"is_active"`
	IsDefault                 bool          `json:"is_default"`
	IsDeleted                 bool          `json:"is_deleted"`
	IsRecurringPaymentEnabled bool          `json:"is_recurring_payment_enabled"`
	PayingByInvoice           bool          `json:"paying_by_invoice"`
	PostIndex                 string        `json:"post_index"`
	PrimaryCard               CardData      `json:"primary_card"`
	RecurringPaymentAmount    int64         `json:"recurring_payment_amount"`
	RecurringPaymentThreshold int64         `json:"recurring_payment_threshold"`
	ReferralShareCode         string        `json:"referral_share_code"`
	Reseller                  string        `json:"reseller"`
	RestrictionLevel          string        `json:"restriction_level"`
	RunningTotals             RunningTotals `json:"running_totals"`
	SendInvoiceEmail          bool          `json:"send_invoice_email"`
	Site                      string        `json:"site"`
	SuspendReason             string        `json:"suspend_reason"`
	Title                     string        `json:"title"`
	UnpaidAmount              int64         `json:"unpaid_amount"`
	UserID                    int64         `json:"user_id"`
	VatPercentage             int64         `json:"vat_percentage"`
}

type CardData struct {
	ID            string            `json:"id"`
	ExpireMonth   int64             `json:"expire_month"`
	ExpireYear    int64             `json:"expire_year"`
	Last4         string            `json:"last4"`
	CardType      string            `json:"card_type"`
	CardHolder    string            `json:"card_holder"`
	Type          string            `json:"type"`
	ProcessorData CardProcessorData `json:"processor_data"`
	IsVerified    bool              `json:"is_verified"`
}

type CardProcessorData struct {
	ID            string                     `json:"id"`
	Object        string                     `json:"object"`
	BillingDetail CardProcessorBillingDetail `json:"billing_details"`
	Card          Card                       `json:"card"`
	Created       int64                      `json:"created"`
	Customer      string                     `json:"customer"`
	LiveMode      bool                       `json:"livemode"`
	MetaData      interface{}                `json:"metadata"`
	Type          string                     `json:"type"`
}

type CardProcessorBillingDetail struct {
	Address CardProcessorBillingAddress `json:"address"`
	Email   string                      `json:"email"`
	Name    string                      `json:"name"`
	Phone   string                      `json:"phone"`
}

type CardProcessorBillingAddress struct {
	City       string `json:"city"`
	Country    string `json:"contry"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}

type Card struct {
	Brand             string       `json:"brand"`
	Checks            interface{}  `json:"checks"`
	Country           string       `json:"country"`
	ExpMonth          int64        `json:"exp_month"`
	ExpYear           int64        `json:"exp_year"`
	Fingerprint       string       `json:"fingerprint"`
	Funding           string       `json:"funding"`
	GeneratedFrom     interface{}  `json:"generated_from"`
	Last4             string       `json:"last4"`
	ThreeDSecureUsage ThreeDSecure `json:"three_d_secure_usage"`
	Wallet            interface{}  `json:"wallet"`
}

type ThreeDSecure struct {
	Supported string `json:"supprted"`
}

type RunningTotals struct {
	CreditAmount    float64 `json:"credit_amount"`
	CreditAvailable float64 `json:"credit_available"`
	DiscountAmount  float64 `json:"discount_amount"`
	Ongoing         float64 `json:"ongoing"`
	Subtotal        float64 `json:"subtotal"`
	Total           float64 `json:"total"`
	VatTax          float64 `json:"vat_tax"`
}
