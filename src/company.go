package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const companyEndpoint = "companies"

const ( // Agent types
	// MasterAgent is the admin of this agency
	MasterAgent = iota
	// RegularAgent is a regular member of an agency
	RegularAgent = iota
)

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
	ExternalID      string        `json:"externalid"`
	S               Session       `json:"S"`
	//AllocationGroupAgencies []AllocationGroupAgency `json`
}

type CompanyProject struct {
	OwnerID   string
	CompanyID string
	ProjectID string
}

type UserAdd struct {
	Level string `json:"Level"`
}

// execute the HTTP requests and get the list of companies that should come out
func execRequestReturnAllCompanies(s *Session, req *http.Request) ([]*Company, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var companies []*Company

	err = json.Unmarshal(responseBytes, &companies)
	if err != nil {
		if s.DumpErrorPayloads {
			fmt.Printf("Dumping Error Payload: %v\n", string(responseBytes))
		}
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	for _, company := range companies {
		company.S = *s
	}

	return companies, nil
}

// execute the HTTP requests and get the single company that should come out
func execRequestReturnSingleCompany(s *Session, req *http.Request) (*Company, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var company *Company

	err = json.Unmarshal(responseBytes, &company)
	if err != nil {
		if s.DumpErrorPayloads {
			fmt.Printf("Dumping Error Payload: %v\n", string(responseBytes))
		}
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}
	company.S = *s

	return company, err
}

func execRequestReturnAssignedCompanyProjects(s *Session, req *http.Request) ([]*CompanyProject, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var companyProjects []*CompanyProject

	err = json.Unmarshal(responseBytes, &companyProjects)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return companyProjects, nil

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

// GetCompany creates the appropriate get request and calls the service function to execute and handle the request
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

func (c *Company) GetAssignedProjects() ([]*CompanyProject, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/projects",
			c.S.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnAssignedCompanyProjects(&c.S, req)
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

// Update PUTs new company details to an existing company (using ID) in the portal
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
			companyEndpoint,
			c.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(&c.S, req)

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

	return executeRequestAndParseStatusCode(&c.S, req)

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

	return execRequestReturnAllAccountsContacts(&c.S, req)
}

//AddCompanyUser adds an Agent 'user' to the company
func (c *Company) AddUsers(u []*UserAdd) error {
	body, err := json.Marshal(u)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v/%v/addusers",
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

	return executeRequestAndParseStatusCode(&c.S, req)

}

// AddUser adds a single user to a company
func (c *Company) AddUser(u *User, permissionlevel int64) error {
	body, err := json.Marshal(u)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v/%v/adduser/%v/%v",
			c.S.Auth.PortalEndpoint,
			companyEndpoint,
			c.ID,
			u.ID,
			permissionlevel,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(&c.S, req)
}
