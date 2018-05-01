package portal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const projectEndpont = "projects"

// Project holds information about a project
type Project struct {
	ID   string
	Name string
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
		return nil, err
	}

	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	project := &Project{}
	err = json.Unmarshal(responseBytes, project)

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
		return nil, err
	}

	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	project := &Project{}
	err = json.Unmarshal(responseBytes, project)

	return project, err
}
