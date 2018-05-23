package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
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
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	ArchivedOn   string  `json:"archivedon"`
	LastLoggedIn string  `json:"lastloggedin"`
	CreatedDate  string  `json:"createddate"`
	S            Session `json:"S"`
}

func (u *User) triggerPasswordResetEmail(email string) error {
	value := []byte(u.Email)
	b := bytes.NewReader(value)

	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/auth/forgotpassword",
			u.S.Auth.PortalEndpoint,
		),
		b,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(&u.S, req)
}

// execute the HTTP requests and get the single Contact that should come out
func execRequestReturnSingleUser(s *Session, req *http.Request) (*User, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var user *User
	err = json.Unmarshal(responseBytes, &user)

	return user, err
}

func execRequestReturnMultipleUsers(s *Session, req *http.Request) ([]*User, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}
	log.Println(string(responseBytes))

	var users []*User

	err = json.Unmarshal(responseBytes, &users)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	for _, user := range users {
		user.S = *s
	}

	return users, err
}

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

	return execRequestReturnMultipleUsers(s, req)
}

func (u *User) GetAccountsContacts() ([]*Account, []*Contact, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/accountscontacts",
			u.S.Auth.PortalEndpoint,
			userEndpoint,
			u.ID,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, nil, err
	}

	return execRequestReturnAllAccountsContacts(&u.S, req)
}

func (u *User) GetCompanyUsers() ([]*User, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/mycompany",
			u.S.Auth.PortalEndpoint,
			userEndpoint,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnMultipleUsers(&u.S, req)
}

//////EOIUSers
func (s *Session) GetEOIUsers() ([]*User, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/eoi-user",
			s.Auth.PortalEndpoint,
			userEndpoint,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnMultipleUsers(s, req)

}

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
func (s *Session) CreateUser(u *User) (*User, error) {
	u.ID = "" // Make sure to blank out the ID
	body, err := json.Marshal(*u)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
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
		return nil, err
	}

	result, err := execRequestReturnSingleUser(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}
	return result, nil
}

//Update saves changes made to contact
func (u *User) Update() error {
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

func (u *User) Delete() error {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			u.S.Auth.PortalEndpoint,
			userEndpoint,
			u.ID,
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
