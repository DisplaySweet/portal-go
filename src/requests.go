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
	req.Header["X-DS-AUTH-APITOKEN"] = []string{s.Auth.APIKey}
	req.Header["X-DS-AUTH-USERAPITOKEN"] = []string{s.Auth.UserAPIKey}
	//req.Header["X-DS-AUTH-COMPANY"] = []string{s.Auth.Company}
	req.Header["X-DS-DATA-APITOKEN"] = []string{s.Auth.APIKey}
	req.Header["X-DS-DATA-COMPANY"] = []string{s.CompanyID}
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
		return errors.New("ERR: Request returned bad status code")
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
