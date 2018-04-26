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
func (s *Session) GetAccountContacts(id string) ([]*AccountContact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/contacts",
			s.Auth.PortalEndpoint,
			accountEndpoint,
			id),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllAccountContacts(s, req)
}

//GetAccountDeposits GETs a list of all DepositResponses belonging to the Account, using their ID
func (s *Session) GetAccountDeposits(id string) ([]*Deposit, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/deposits",
			s.Auth.PortalEndpoint,
			accountEndpoint,
			id),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllAccountDeposits(s, req)
}

//CreateAccount POSTs a new Account to the portal
func (a *Account) CreateAccount(s *Session) (int, error) {
	a.ID = "" // Make sure to blank out the ID
	body, err := json.Marshal(*a)
	if err != nil {
		return 0, err
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
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}

//UpdateAccount PUTs new Account data to an existing Account, using their ID
func (a *Account) UpdateAccount(s *Session) (int, error) {
	body, err := json.Marshal(*a)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			accountEndpoint,
			a.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}

//DeleteAccount DELETEs an Account, using their ID
func (a *Account) DeleteAccount(s *Session) (int, error) {
	body, err := json.Marshal(*a)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			accountEndpoint,
			a.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}

//BulkDeleteAccounts DELETEs Accounts using a list of Account IDs
func (s *Session) BulkDeleteAccounts(list []*Account) (int, error) {
	body, err := json.Marshal(list)
	if err != nil {
		return 0, err
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
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}
