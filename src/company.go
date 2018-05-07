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
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Active          bool          `json:"active"`
	StripeAccountID string        `json:"stripeaccountid"`
	CreatedByID     string        `json:"createdbyid"`
	CreatedDate     string        `json:"createddate"`
	CreatedBy       Agent         `json:"createdby"`
	Projects        []Project     `json:"projects"`
	UserCompanies   []UserCompany `json:"usercompanies"`
	Events          []Event       `json:"events"`
	S               *Session      `json:"S"`
	//AllocationGroupAgencies []AllocationGroupAgency `json`

}

type accountsContactsResponse struct {
	Accounts []*Account
	Contacts []*Contact
}

// execute the HTTP requests and get the list of companies that should come out
func execRequestReturnAllCompanies(s *Session, req *http.Request) ([]*Company, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var temp []*Company

	list := make([]*Company, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	for _, company := range temp {
		company.S = s
		list = append(list, company)
	}

	return list, nil
}

// execute the HTTP requests and get the single company that should come out
func execRequestReturnSingleCompany(s *Session, req *http.Request) (*Company, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	company := &Company{}
	err = json.Unmarshal(responseBytes, company)
	company.S = s

	return company, err
}

func execRequestReturnAllAccountsContacts(s *Session, req *http.Request) ([]*Account, []*Contact, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, nil, err
	}

	temp := &accountsContactsResponse{}

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, nil, err
	}

	return temp.Accounts, temp.Contacts, nil
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnAllCompanies(s, req)
}

//GetCompanyByID creates the appropriate get request and calls the service function to execute and handle the request
func (s *Session) GetCompany(id string) (*Company, error) {
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleCompany(s, req)
}

//CreateCompany POSTs a new company to the portal
func (s *Session) CreateCompany(company *Company) (*Company, error) {
	company.ID = ""
	body, err := json.Marshal(*company)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
func (c *Company) Update() error {
	body, err := json.Marshal(*c)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			c.S.Auth.PortalEndpoint,
			accountEndpoint,
			c.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(c.S, req)

}

//DeleteCompany removes an existing company (using ID) from the portal
func (c *Company) Delete() error {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			c.S.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(c.S, req)

}

//GetCompanyAccountsContacts GETs all existing accounts and contacts for this company
func (c *Company) GetAccountsContacts() ([]*Account, []*Contact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/accountscontacts",
			c.S.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, nil, err
	}

	return execRequestReturnAllAccountsContacts(c.S, req)
}

//AddCompanyUser adds an Agent 'user' to the company
func (c *Company) AddUsers(a []*Agent) error {
	body, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v/%v",
			c.S.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(c.S, req)

}
