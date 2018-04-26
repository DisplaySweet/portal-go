package portal

type Account struct {
	ID              string `json:""`
	Phone           string `json:""`
	Mobile          string `json:""`
	Email           string `json:""`
	AccountName     string `json:""`
	AccountType     string `json:""`
	Industry        string `json:""`
	Website         string `json:""`
	AddressID       string `json:""`
	Address         Address
	ReferralCode    string `json:""`
	OwnerID         string `json:""`
	Owner           Company
	AgentID         string `json:""`
	Agent           Agent
	Notes           string `json:""`
	AccountContacts []AccountContact
}

//GetAllAccounts GETs a list of all accounts
func GetAllAccounts() {

}

//GetAccount GETs one account, using their ID
func GetAccount() {

}

//GetAccountContacts GETs a list of AccountContacts that belong to the Account, using their ID
func GetAccountContacts() {

}

//GetAccountDeposits GETs a list of all DepositResponses belonging to the Account, using their ID
func GetAccountDeposits() {

}

//CreateAccount POSTs a new Account to the portal
func CreateAccount() {

}

//UpdateAccount PUTs new Account data to an existing Account, using their ID
func UpdateAccount() {

}

//DeleteAccount DELETEs an Account, using their ID
func DeleteAccount() {

}

//BulkDeleteAccounts DELETEs Accounts using a list of Account IDs
func BulkDeleteAccounts() {

}
