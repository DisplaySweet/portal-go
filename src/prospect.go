package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const prospectEndpoint = "prospects"
const CreateProspectFromContact = prospectEndpoint + "/createFromContact"

//Prospect holds the data for a prospect
type Prospect struct {
	ID string `json:"id"`
	//BuyerCode   string ``
	OwnerID     string `json:"ownerid"`
	Owner       Company
	CompanyID   string `json:"companyid"`
	Company     Company
	ProjectID   string `json:"projectid"`
	Project     Project
	AgentID     string `json:"agentid"`
	Agent       Agent
	SolicitorID string `json:"solicitorid"`
	Solicitor   ProspectSolicitor
	DateCreated string `json:"DateCreated"`
	Buyers      []ProspectBuyer
	EventID     string `json:"eventid"`
	SeatCount   int    `json:"seatcount"`
	// ScheduleID   string
	// Schedule     Schedule
	WasScheduled bool `json:"wasscheduled"`
	HasPaid      bool `json:"haspaid"`
	//HasArrived   string `json:"hasarrived"`
	//ArrivalSent bool
	//ArrivalTime string
	Offers      []Offer
	Attachments []ProspectAttachment
	Deposits    []Deposit
	s           *Session
}

func execRequestReturnAllEventProspects(s *Session, req *http.Request) ([]*Prospect, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var temp map[string]*Prospect

	list := make([]*Prospect, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	for _, prospect := range temp {
		prospect.s = s
		list = append(list, prospect)
	}

	return list, nil
}

// execute the HTTP requests and get the single Prospect that should come out
func execRequestReturnSingleProspect(s *Session, req *http.Request) (*Prospect, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	prospect := &Prospect{}
	err = json.Unmarshal(responseBytes, prospect)

	return prospect, err
}

//GetAllEventProspects GETs a list of all prospects for an event
func (s *Session) GetAllEventProspects(eventID string) ([]*Prospect, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"/%v/%v/%v",
			s.Auth.PortalEndpoint,
			prospectEndpoint,
			eventID,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnAllEventProspects(s, req)
}

//The relevant method in the prospects controller shouldn't really belong to a Prospect,
//TODO: confirm coverage of prospects controller at a later date
//func (s *Session) GetAllEventsForDate() (e []*Event)

//UpdateProspect POSTs updates a prospect with new data, using the prospect object
func (s *Session) CreateProspect(contactId string) (*Prospect, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			CreateProspectFromContact,
			contactId,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleProspect(s, req)
}

//UpdateProspect PUTs updates a prospect with new data, using the prospect object
func (p *Prospect) Update() error {
	body, err := json.Marshal(*p)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"/%v/%v",
			p.s.Auth.PortalEndpoint,
			prospectEndpoint,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(p.s, req)
}

//UpdateAgent POSTs a newly assigned Agent to the Prospect
func (p *Prospect) UpdateAgent(a *Agent) error {
	body, err := json.Marshal(*a)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"/%v/%v/%v/agent",
			p.s.Auth.PortalEndpoint,
			prospectEndpoint,
			p.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(p.s, req)
}

//UpdateSchedule updates the schedules assigned to a prospect
func (p *Prospect) UpdateSchedule(s *Schedule) error {
	body, err := json.Marshal(*s)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"/%v/%v/%v/schedules",
			p.s.Auth.PortalEndpoint,
			prospectEndpoint,
			p.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(p.s, req)
}
