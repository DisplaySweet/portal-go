package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const prospectEndpoint = "prospects"

//Prospect holds the data for a prospect
type Prospect struct {
	ID          string
	BuyerCode   string
	OwnerID     string
	Owner       Company
	CompanyID   string
	Company     Company
	ProjectID   string
	Project     Project
	AgentID     string
	Agent       Agent
	SolicitorID string
	Solicitor   ProspectSolicitor
	DateCreated string
	Buyers      []ProspectBuyer
	EventID     string
	SeatCount   int
	// ScheduleID   string
	// Schedule     Schedule
	WasScheduled bool
	HasPaid      bool
	HasArrived   string
	ArrivalSent  bool
	ArrivalTime  string
	Offers       []Offer
	Attachments  []ProspectAttachment
	Deposits     []Deposit
	s            *Session
}

func execRequestReturnAllEventProspects(s *Session, req *http.Request) ([]*Prospect, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Prospect

	list := make([]*Prospect, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, prospect := range temp {
		prospect.s = s
		list = append(list, prospect)
	}

	return list, nil
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
		return nil, err
	}

	return execRequestReturnAllEventProspects(s, req)
}

//The relevant method in the prospects controller shouldn't really belong to a Prospect,
//TODO: confirm coverage of prospects controller at a later date
//func (s *Session) GetAllEventsForDate() (e []*Event)

//UpdateProspect PUTs updates a prospect with new data, using the prospect object
func (p *Prospect) UpdateProspect() (int, error) {
	body, err := json.Marshal(*p)
	if err != nil {
		return 0, nil
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
	return executeRequestAndGetStatusCode(p.s, req)
}

//UpdateAgent POSTs a newly assigned Agent to the Prospect
func (p *Prospect) UpdateAgent(a *Agent) (int, error) {
	body, err := json.Marshal(*a)
	if err != nil {
		return 0, nil
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
	return executeRequestAndGetStatusCode(p.s, req)
}

//UpdateSchedule updates the schedules assigned to a prospect
func (p *Prospect) UpdateSchedule(s *Schedule) (int, error) {
	body, err := json.Marshal(*s)
	if err != nil {
		return 0, nil
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
	return executeRequestAndGetStatusCode(p.s, req)
}
