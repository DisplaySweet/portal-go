package portal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const userCompanyEndpoint = "usercompanies"

//UserCompany holds information about a UserCompany
type UserCompany struct {
	User        User    `json:"user, omitempty"`
	Company     Company `json:"company, omitempty"`
	ID          string  `json:"id"`
	UserID      string  `json:"userid"`
	CompanyID   string  `json:"companyid"`
	Role        int     `json:"role"`
	CreatedDate string  `json:"createddate"`
}

// type UserCompany struct {
// 	User    User
// 	Company Company
// }

func execRequestReturnAllUserCompanies(s *Session, req *http.Request) ([]*UserCompany, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var UCs []*UserCompany

	err = json.Unmarshal(responseBytes, &UCs)
	if err != nil {
		if s.DumpErrorPayloads {
			fmt.Printf("Dumping Error Payload: %v\n", string(responseBytes))
		}
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return UCs, nil
}

//GetAllCompanies creates the appropriate get request and calls the service function to execute and handle the request
func (s *Session) GetAllUserCompanies() ([]*UserCompany, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			userCompanyEndpoint),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnAllUserCompanies(s, req)
}
