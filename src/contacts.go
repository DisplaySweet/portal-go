package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const contactEndpoint = "usercontacts"

// Contact holds information of a contact
type Contact struct {
	ID          string  `json:"id"`
	Firstname   string  `json:"firstname"`
	Lastname    string  `json:"lastname"`
	Mobile      string  `json:"mobile"`
	Phone       string  `json:"phone"`
	Email       string  `json:"email"`
	DOB         string  `json:"dob"`
	Nationality string  `json:"nationality"`
	HasDeposit  bool    `json:"hasdeposit"`
	AddressID   string  `json:"addressid"`
	Address     Address `json:"address"`
	AgentID     string  `json:"agentid"`
	ManageByID  string  `json:"managebyid"`
	s           *Session
}

// execute the HTTP requests and get the single Contact that should come out
func execRequestReturnSingleContact(s *Session, req *http.Request) (*Contact, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	contact := &Contact{}
	err = json.Unmarshal(responseBytes, contact)

	return contact, err
}

// GetContactByID gets a single contact by its ID
func (s *Session) GetContactByID(id string) (*Contact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			contactEndpoint,
			id,
		),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleContact(s, req)
}

// GetContactByName gets the first contact that matches the provided name
func (s *Session) GetContactByName(firstname string, lastname string) (*Contact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v?firstname=%v&lastname=%v",
			s.Auth.PortalEndpoint,
			contactEndpoint,
			firstname,
			lastname,
		),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleContact(s, req)
}

// GetContacts gets all contacts
func (s *Session) GetContacts() ([]*Contact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			contactEndpoint,
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

	var response []*Contact
	err = json.Unmarshal(responseBytes, response)
	for _, contact := range response {
		contact.s = s
	}

	return response, err
}

//Update saves changes made to contact
func (c *Contact) Update() error {
	body, err := json.Marshal(*c)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			c.s.Auth.PortalEndpoint,
			contactEndpoint,
			c.ID,
		),
		bytes.NewReader(body),
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	err = executeRequestAndParseStatusCode(c.s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}
	return nil
}

// Create generates a new contact from the supplied data
// Create should return the contact that was just created.
func (s *Session) CreateContact(c *Contact) error {
	c.ID = "" // Make sure to blank out the ID
	body, err := json.Marshal(*c)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			contactEndpoint,
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
