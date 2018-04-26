package portal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const prospectbuyerEndpoint = "prospectbuyers"

type ProspectBuyer struct {
	ID                 string `json:"id"`
	Phone              string `json:"phone"`
	Mobile             string `json:"mobile"`
	Email              string `json:"email"`
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	DOB                string `json:"dob"`
	Nationality        string `json:"nationality"`
	AddressID          string `json:"addressid"`
	Address            Address
	ProspectID         string `json:"prospectid"`
	Prospect           Prospect
	CompanyID          string `json:"company"`
	ProjectID          string `json:"project"`
	FIRB               int    `json:"firb"`
	Primary            bool   `json:"primary"`
	OriginContactUser  string `json:"origincontactuser"`
	OriginatingContact string `json:"originatingcontact"`
	OriginAccountUser  string `json:"originaccountuser"`
	OriginatingAccount string `json:"originatingaccount"`
}

func execRequestReturnAllProspectsBuyers(s *Session, req *http.Request) ([]*ProspectBuyer, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*ProspectBuyer

	list := make([]*ProspectBuyer, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, prospectbuyer := range temp {
		list = append(list, prospectbuyer)
	}

	return list, nil
}

//GetProspectBuyers GETs a list of ProspectBuyers that belong to the session
func (s *Session) GetProspectBuyers() ([]*ProspectBuyer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			prospectbuyerEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllProspectsBuyers(s, req)
}

//GetAvailableProspectBuyers GETs a list of ProspectBuyers where none of their Offers have been cancelled that belong to the session
func (s *Session) GetAvailableProspectBuyers() ([]*ProspectBuyer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/available",
			s.Auth.PortalEndpoint,
			prospectbuyerEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllProspectsBuyers(s, req)
}
