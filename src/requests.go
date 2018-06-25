package portal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type accountsContactsResponse struct {
	Accounts []*Account
	Contacts []*Contact
}

// Add required headers and execute the request
func executeRequest(s *Session, req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header["X-DS-AUTH-USERAPITOKEN"] = []string{s.Auth.UserAPIKey}
	req.Header["X-DS-AUTH-COMPANY"] = []string{s.Auth.Company}
	req.Header["X-DS-DATA-COMPANY"] = []string{s.Auth.Company}
	req.Header["X-DS-DATA-PROJECT"] = []string{s.ProjectID}

	return http.DefaultClient.Do(req)
}

// Execute the request and return the bytes that come back
func executeRequestAndGetBodyBytes(s *Session, req *http.Request) ([]byte, error) {
	//fmt.Fprintln(os.Stderr, req.URL)
	response, err := executeRequest(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		err = fmt.Errorf("Error in file: %v line %v. Expected content in response but got status code %v", ErrorFile(), ErrorLine(), response.StatusCode)
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}

// Execute the request and return the bytes that come back
func executeRequestAndGetBodyBytesModifiedStatus(s *Session, req *http.Request) ([]byte, int, error) {
	//fmt.Fprintln(os.Stderr, req.URL)
	response, err := executeRequest(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, 500, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 && response.StatusCode != 201 {
		err = fmt.Errorf("Error in file: %v line %v. Expected content in response but got status code %v", ErrorFile(), ErrorLine(), response.StatusCode)
		return nil, response.StatusCode, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, response.StatusCode, err
	}

	return body, response.StatusCode, err
}

func executeRequestAndParseStatusCode(s *Session, req *http.Request) error {
	response, err := executeRequest(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	switch response.StatusCode {
	case 200:
	case 204:
		break
	default:
		err := fmt.Sprintf("ERR: Request returned bad status code: %v", response.StatusCode)
		return errors.New(err)
	}

	return nil
}

func execRequestReturnAllAccountsContacts(s *Session, req *http.Request) ([]*Account, []*Contact, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, nil, err
	}

	temp := &accountsContactsResponse{}

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, nil, err
	}

	return temp.Accounts, temp.Contacts, nil
}
