package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const companyEndpoint = "companies"

// Company holds information about companies
type Company struct {
	ID                      string `json:"id"`
	Name                    string `json:"name"`
	Active                  bool   `json:"active"`
	StripeAccountID         string `json:"stripeaccountid`
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

func CreateCompany() {

}

func UpdateCompany() {

}

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

func GetCompanyAccountsAndContacts() {

}

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
