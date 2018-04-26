package portal

//Deposit holds the deposit data
type Deposit struct {
	ID              string
	CompanyID       string
	ProjectID       string
	AccountID       string
	ProspectID      string
	ProspectBuyerID string
	Amount          int
	DepositType     string
	DepositStatus   string
	StripeChargeID  string
	CreatedDate     string
}
