package portal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Add required headers and execute the request
func executeRequest(s *Session, req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header["X-DS-AUTH-APITOKEN"] = []string{s.Auth.APIKey}
	req.Header["X-DS-AUTH-COMPANY"] = []string{s.Auth.Company}
	req.Header["X-DS-DATA-APITOKEN"] = []string{s.Auth.APIKey}
	req.Header["X-DS-DATA-COMPANY"] = []string{s.Company.ID}
	req.Header["X-DS-DATA-PROJECT"] = []string{s.Project.ID}

	return http.DefaultClient.Do(req)
}

// Execute the request and return the bytes that come back
func executeRequestAndGetBodyBytes(s *Session, req *http.Request) ([]byte, error) {
	response, err := executeRequest(s, req)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("Request Body Bytes: failed to execute http request: ERR: %v", err))
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func executeRequestAndParseStatusCode(s *Session, req *http.Request) error {
	response, err := executeRequest(s, req)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("Request Status Code: execute request and retrieve status code: ERR: %v", err))
		return err
	}

	switch response.StatusCode {
	case 200:
	case 204:
		break
	default:
		return errors.New("Not implemented")
	}

	return nil
}
