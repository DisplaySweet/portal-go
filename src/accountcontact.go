package portal

type AccountContact struct {
	ID        string
	AgentID   string
	AccountID string
	Account   Account
	ContactID string
	Contact   Contact
}
