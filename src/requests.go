package portal

import (
	"io/ioutil"
	"net/http"
)

// Add required headers and execute the request
func executeRequest(s *Session, req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-DS-AUTH-APITOKEN", s.Auth.APIKey)
	req.Header.Add("X-DS-DATA-COMPANY", s.Company.ID)

	return http.DefaultClient.Do(req)
}

// Execute the request and return the bytes that come back
func executeRequestAndGetBodyBytes(s *Session, req *http.Request) ([]byte, error) {
	response, err := executeRequest(s, req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}
