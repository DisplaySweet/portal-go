package portal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const projectEndpont = "projects"

//TODO: Not all modelled fields have yet been included
// Project holds information about a project
type Project struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyID string `json:"companyid"`
	Company   Company
	Active    bool `json:"active"`
	Listings  []Listing
	Events    []Event
	Offers    []Offer
}

// GetProjectByName returns a project queried by name
func (s *Session) GetProjectByName(name string) (*Project, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v?name=%v",
			s.Auth.PortalEndpoint,
			projectEndpont,
			name,
		),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	project := &Project{}
	err = json.Unmarshal(responseBytes, project)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return project, err
}

// GetProjectByID returns a project queried by ID
func (s *Session) GetProjectByID(id string) (*Project, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v?name=%v",
			s.Auth.PortalEndpoint,
			projectEndpont,
			id,
		),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	project := &Project{}
	err = json.Unmarshal(responseBytes, project)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return project, err
}
