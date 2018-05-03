package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const prospectsolicitorEndpoint = "eoi/prospectsolicitors"

////////
//ProspectSolicitors is EOI-related
//Leave EOI-related resources alone
///////

type ProspectSolicitor struct {
	ID          string `json:"id"`
	Phone       string `json:"phone"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	CompanyName string `json:"company_name"`
	FullName    string `json:"full_name"`
	AddressID   string `json:"addressid"`
	Address     Address
	CompanyID   string `json:"companyid"`
	ProjectID   string `json:"projectid"`
	s           *Session
}

func execRequestReturnSingleProspectSolicitor(s *Session, req *http.Request) (*ProspectSolicitor, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	prospectSolicitor := &ProspectSolicitor{}
	err = json.Unmarshal(responseBytes, prospectSolicitor)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return prospectSolicitor, err
}

//GetSolicitor GETs a Solicitor
func (s *Session) GetSolicitor() (*ProspectSolicitor, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			prospectsolicitorEndpoint,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleProspectSolicitor(s, req)
}

//Create POSTs new solicitor information to a Prospect
func (ps *ProspectSolicitor) Create() error {
	ps.ID = ""
	body, err := json.Marshal(*ps)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			ps.s.Auth.PortalEndpoint,
			prospectsolicitorEndpoint,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(ps.s, req)
}

//Update PUTs new information to an existing solicitor
func (ps *ProspectSolicitor) Update() error {
	body, err := json.Marshal(*ps)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			ps.s.Auth.PortalEndpoint,
			prospectsolicitorEndpoint,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(ps.s, req)
}

// //Delete DELETEs a solicitor from a prospect EOI RELATED
// func (ps *ProspectSolicitor) Delete() (int, error) {
// 	req, err := http.NewRequest(
// 		"DELETE",
// 		fmt.Sprintf(
// 			"%v/%v",
// 			ps.s.Auth.PortalEndpoint,
// 			prospectsolicitorEndpoint,
// 		),
// 		bytes.NewReader(body),
// 	)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return executeRequestAndGetStatusCode(ps.s, req)
// }
