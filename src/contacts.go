package portal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const contactEndpoint = "usercontacts"

// Contact holds information of a contact
type Contact struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

// execute the HTTP requests and get the single Contact that should come out
func execRequestReturnSingleContact(s *Session, req *http.Request) (*Contact, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
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
		return nil, err
	}

	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var response []*Contact
	err = json.Unmarshal(responseBytes, response)

	return response, err
}

// SendUpdate saves changes made to contact
func (c *Contact) SendUpdate(s *Session) error {
	body, err := json.Marshal(*c)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			contactEndpoint,
			c.ID,
		),
		bytes.NewReader(body),
	)

	if err != nil {
		return err
	}

	response, err := executeRequest(s, req)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 200:
	case 204:
		break
	default:
		return errors.New("Did not get a success code from the portal")
	}

	return nil
}

// Create generates a new contact from the supplied data
func (c *Contact) Create(s *Session) error {
	c.ID = "" // Make sure to blank out the ID
	body, err := json.Marshal(*c)
	if err != nil {
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
		return err
	}

	response, err := executeRequest(s, req)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 204:
		break
	default:
		return errors.New("Not implemented")
	}

	return nil
}
