package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const eventEndpoint = "events"

// Event holds information about an event
type Event struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CompanyID   string `json:"companyid"`
	Company     Company
	ProjectID   string `json:"projectid"`
	Project     Project
	Terminals   int32  `json:"terminals"`
	CreatedDate string `json:"createddate"`
	TZInfo      string `json:"timezoneinfoid"`
	Dates       []EventDate
	ArchivedOn  string `json:"archivedon"`
	s           *Session
}

func execRequestReturnAllEvents(s *Session, req *http.Request) ([]*Event, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Event

	list := make([]*Event, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, event := range temp {
		event.s = s
		list = append(list, event)
	}

	return list, nil
}

func execRequestReturnSingleEvent(s *Session, req *http.Request) (*Event, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	event := &Event{}
	err = json.Unmarshal(responseBytes, event)
	event.s = s

	return event, err
}

//GetEvents GETs a list of all Events
func (s *Session) GetEvents() ([]*Event, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			eventEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllEvents(s, req)
}

//GetEventByID GETs an individual event using the event ID
func (s *Session) GetEventByID(id string) (*Event, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%v/%v",
			s.Auth.PortalEndpoint,
			eventEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleEvent(s, req)
}

//CreateEvent POSTs new event data to the endpoint
func (s *Session) CreateEvent(e *Event) (int, error) {
	e.ID = ""
	body, err := json.Marshal(*e)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			eventEndpoint,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(s, req)
}

//UpdateEventByID PUTs new data to an existing event using it's ID
func (e *Event) Update() (int, error) {
	body, err := json.Marshal(*e)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			e.s.Auth.PortalEndpoint,
			eventEndpoint,
			e.ID,
		),
		bytes.NewReader(body),
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(e.s, req)
}

//DeleteEventByID DELETEs an existing event using it's ID
func (e *Event) Delete() (int, error) {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			e.s.Auth.PortalEndpoint,
			eventEndpoint,
			e.ID,
		),
		nil,
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(e.s, req)
}
