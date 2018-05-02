package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const accountEndpoint = "accounts"

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
	s               *Session
}

func execRequestReturnAllAccounts(s *Session, req *http.Request) ([]*Account, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Account

	list := make([]*Account, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, account := range temp {
		account.s = s
		list = append(list, account)
	}

	return list, nil
}

func execRequestReturnSingleAccount(s *Session, req *http.Request) (*Account, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	account := &Account{}
	err = json.Unmarshal(responseBytes, account)
	account.s = s

	return account, err
}

func execRequestReturnAllAccountContacts(s *Session, req *http.Request) ([]*AccountContact, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*AccountContact

	list := make([]*AccountContact, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, ac := range temp {
		list = append(list, ac)
	}

	return list, nil
}

func execRequestReturnAllAccountDeposits(s *Session, req *http.Request) ([]*Deposit, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Deposit

	list := make([]*Deposit, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, deposit := range temp {
		list = append(list, deposit)
	}

	return list, nil
}

//GetAllAccounts GETs a list of all accounts
func (s *Session) GetAllAccounts() ([]*Account, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			accountEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllAccounts(s, req)
}

//GetAccount GETs one account, using their ID
func (s *Session) GetAccount(id string) (*Account, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			accountEndpoint,
			id),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleAccount(s, req)
}

//GetAccountContacts GETs a list of AccountContacts that belong to the Account, using their ID
func (a *Account) GetOwnedContacts() ([]*AccountContact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/contacts",
			a.s.Auth.PortalEndpoint,
			accountEndpoint,
			a.ID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllAccountContacts(a.s, req)
}

//GetAccountDeposits GETs a list of all DepositResponses belonging to the Account, using their ID
func (a *Account) GetOwnedDeposits() ([]*Deposit, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/deposits",
			a.s.Auth.PortalEndpoint,
			accountEndpoint,
			a.ID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllAccountDeposits(a.s, req)
}

//CreateAccount POSTs a new Account to the portal
func (s *Session) CreateAccount(account *Account) error {
	account.ID = "" // Make sure to blank out the ID
	body, err := json.Marshal(*account)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			accountEndpoint,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	return executeRequestAndParseStatusCode(s, req)
}

//UpdateAccount PUTs new Account data to an existing Account, using their ID
func (a *Account) Update() error {
	body, err := json.Marshal(*a)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			a.s.Auth.PortalEndpoint,
			accountEndpoint,
			a.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	return executeRequestAndParseStatusCode(a.s, req)

}

//Delete DELETEs an Account, using their ID
func (a *Account) Delete() error {
	body, err := json.Marshal(*a)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			a.s.Auth.PortalEndpoint,
			accountEndpoint,
			a.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	return executeRequestAndParseStatusCode(a.s, req)

}

//Delete DELETEs Accounts using a list of Account IDs
func (s *Session) Delete(list []*Account) error {
	body, err := json.Marshal(list)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/bulk",
			s.Auth.PortalEndpoint,
			accountEndpoint,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return err
	}

	return executeRequestAndParseStatusCode(s, req)
}
