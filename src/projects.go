package portal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const projectEndpoint = "projects"

//TODO: Not all modelled fields have yet been included
// Project holds information about a project
type Project struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	CompanyID  string    `json:"companyid"`
	Company    Company   `json:"omitempty"`
	Active     bool      `json:"active"`
	Listings   []Listing `json:"omitempty"`
	Events     []Event   `json:"omitempty"`
	Offers     []Offer   `json:"omitempty"`
	ExternalID string    `json:"externalid"`
	S          Session   `json:"S"`
}

func (s *Session) GetAllProjects() ([]*Project, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			projectEndpoint,
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

	var projects []*Project

	err = json.Unmarshal(responseBytes, &projects)
	if err != nil {
		if s.DumpErrorPayloads {
			fmt.Printf("Dumping Error Payload: %v\n", string(responseBytes))
		}
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	for _, project := range projects {
		project.S = *s
	}

	return projects, nil

}

// GetProjectByName returns a project queried by name
func (s *Session) GetProjectByName(name string) (*Project, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v?name=%v",
			s.Auth.PortalEndpoint,
			projectEndpoint,
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

// GetProjectByID returns a project qußeried by ID
func (s *Session) GetProjectByID(id string) (*Project, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			projectEndpoint,
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
		return nil, errB
	}

	var project *Project

	err = json.Unmarshal(responseBytes, &project)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}
	project.S = *s

	return project, err
}

//AddCompany includes a company in a project
func (p *Project) AddCompany(id string) error {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v/%v/addcompany/%v",
			p.S.Auth.PortalEndpoint,
			projectEndpoint,
			p.ID,
			id,
		),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(&p.S, req)
}
