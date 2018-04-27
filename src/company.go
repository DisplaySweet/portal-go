package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const companyEndpoint = "companies"

//Company holds information about companies
type Company struct {
	ID                      string `json:"id"`
	Name                    string `json:"name"`
	Active                  bool   `json:"active"`
	StripeAccountID         string `json:"stripeaccountid"`
	CreatedByID             string `json:"createdbyid"`
	CreatedDate             string `json:"createddate"`
	CreatedBy               Agent
	Projects                []Project
	UserCompanies           []UserCompany
	Events                  []Event
	AllocationGroupAgencies []AllocationGroupAgency
}

// execute the HTTP requests and get the list of companies that should come out
func execRequestReturnAllCompanies(s *Session, req *http.Request) ([]*Company, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Company

	list := make([]*Company, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, company := range temp {
		list = append(list, company)
	}

	return list, nil
}

// execute the HTTP requests and get the single company that should come out
func execRequestReturnSingleCompany(s *Session, req *http.Request) (*Company, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	company := &Company{}
	err = json.Unmarshal(responseBytes, company)

	return company, err
}

func execRequestReturnAllAccountsContacts(s *Session, req *http.Request) ([]*Account, []*Contact, error) {

	return nil, nil, nil
}

//GetAllCompanies creates the appropriate get request and calls the service function to execute and handle the request
func (s *Session) GetAllCompanies() ([]*Company, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			companyEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllCompanies(s, req)
}

//GetCompanyByID creates the appropriate get request and calls the service function to execute and handle the request
func (s *Session) GetCompanyByID(id string) (*Company, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			companyEndpoint,
			id),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleCompany(s, req)
}

//CreateCompany POSTs a new company to the portal
func (c *Company) CreateCompany(s *Session) (*Company, error) {
	c.ID = ""
	body, err := json.Marshal(*c)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			companyEndpoint),
		bytes.NewReader(body),
	)

	return execRequestReturnSingleCompany(s, req)
}

//UpdateCompany PUTs new company details to an existing company (using ID) in the portal
func (c *Company) UpdateCompany(s *Session) (int, error) {
	body, err := json.Marshal(*c)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			accountEndpoint,
			c.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}

//DeleteCompany removes an existing company (using ID) from the portal
func (c *Company) DeleteCompany(s *Session) (int, error) {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID,
		),
		nil,
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}

//GetCompanyAccountsAndContacts GETs all existing accounts and contacts for this company
func (c *Company) GetCompanyAccountsAndContacts(s *Session) ([]*Account, []*Contact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/accountscontacts",
			s.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID,
		),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	return execRequestReturnAllAccountsContacts(s, req)
}

//AddCompanyUser adds an Agent 'user' to the company
func (c *Company) AddCompanyUser(s *Session, a []*Agent) (int, error) {
	body, err := json.Marshal(a)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}
