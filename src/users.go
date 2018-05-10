package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const userEndpoint = "users"

// User holds information of a user
type User struct {
	ID           string  `json:"id"`
	Active       bool    `json:"active"`
	SuperUser    bool    `json:"superuser"`
	Firstname    string  `json:"firstname"`
	Lastname     string  `json:"lastname"`
	Email        string  `json:"email"`
	Username     string  `json:"-"`
	Password     string  `json:"-"`
	ArchivedOn   string  `json:"archivedon"`
	LastLoggedIn string  `json:"lasloggedin"`
	CreatedDate  string  `json:"createddate"`
	S            Session `json:"S"`
}

// execute the HTTP requests and get the single Contact that should come out
func execRequestReturnSingleUser(s *Session, req *http.Request) (*User, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	user := &User{}
	err = json.Unmarshal(responseBytes, user)

	return user, err
}

///////GetUserCompany

//GetUsers gets all users
func (s *Session) GetUsers() ([]*User, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			userEndpoint,
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

	var response []*User
	err = json.Unmarshal(responseBytes, response)
	for _, user := range response {
		user.S = *s
	}

	return response, err
}

///////MycompanyGetUsers

//////EOIUSers

// GetUserByID gets a single contact by its ID
func (s *Session) GetUserByID(id string) (*User, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			userEndpoint,
			id,
		),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleUser(s, req)
}

// Create generates a new contact from the supplied data
// Create should return the user that was just created.
func (s *Session) CreateUser(u *User) error {
	u.ID = "" // Make sure to blank out the ID
	body, err := json.Marshal(*u)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			userEndpoint,
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

// SendUpdate saves changes made to contact
func (u *User) SendUpdate() error {
	body, err := json.Marshal(*u)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			u.S.Auth.PortalEndpoint,
			userEndpoint,
			u.ID,
		),
		bytes.NewReader(body),
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

///////DELETE

//GET ACCOUNTSCONTACTS
