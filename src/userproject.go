package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const userprojectEndpoint = "userprojects"

type UserProject struct {
	ID                string
	User              User
	UserID            string
	CompanyID         string
	Project           Project
	ProjectID         string
	AddedBy           User
	AddedByID         string
	ManagedBy         User
	ManagedByID       string
	PermissionGroup   PermissionGroup
	PermissionGroupID string
	CreatedDate       string
	S                 Session
}

func execRequestReturnAllUserProjects(s *Session, req *http.Request) []*UserProject.error {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var response []*UserProject
	err = json.Unmarshal(responseBytes, response)
	for _, userproject := range response {
		userproject.S = *s
	}

	return response, err
}

func (s *Session) GetUserProjects() ([]*UserProject, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			userprojectEndpoint,
		),
		nil,
	)

	if err != nil {
		err.fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return nil, execRequestReturnAllUserProjects(s, req)
}

func (s *Session) CreateUserProject(up *UserProject) error {
	up.ID = ""
	body, err := json.Marshal(*up)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			userprojectEndpoint,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	err = executeRequestAndParseStatusCode(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}
	return nil
}

func (up *UserProject) Update() (*User, error){
	body, err := json.Marshal(*up)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}
	
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			up.S.Auth.PortalEndpoint,
			userprojectEndpoint,
			up.ID,
		),
		bytes.NewReader(body)
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return execRequestReturnSingleUser(s, req)
}

func (up *UserProject) DeleteUserProject() error {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			up.S.Auth.PortalEndpoint,
			userprojectEndpoint,
			up.ID,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	err = executeRequestAndParseStatusCode(&u.S, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}
	return nil
}

func (up *UserProject) GetAccountsContacts() ([]*Account, []*Contact, error) { {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/accountscontacts",
			c.S.Auth.PortalEndpoint,
			companyEndpoint,
			up.ID,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, nil, err
	}

	return execRequestReturnAllAccountsContacts(&c.S, req)
}
