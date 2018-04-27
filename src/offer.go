package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const offerEndpoint = "offers"

type Offer struct {
	ID                     string
	CompanyID              string
	Company                Company
	ProjectID              string
	Project                Project
	ListingID              string
	Listing                Listing
	AgentID                string
	Agent                  Agent
	ProspectID             string
	Prospect               Prospect
	Attachments            []OfferAttachment
	TimeOfOffer            string
	Price                  float64
	FundsReceived          float64
	Status                 string
	OfferStatusChangedDate string
}

func execRequestReturnAllOffers(s *Session, req *http.Request) ([]*Offer, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Offer

	list := make([]*Offer, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	for _, offer := range temp {
		list = append(list, offer)
	}

	return list, nil
}

func execRequestReturnSingleOffer(s *Session, req *http.Request) (*Offer, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	offer := &Offer{}
	err = json.Unmarshal(responseBytes, offer)

	return offer, err
}

//GetAllOffers returns a list of Offers -- this won't work yet
func (s *Session) GetAllOffers() ([]*Offer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			offerEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllOffers(s, req)
}

//GetPendingOffers returns a list of offers where status is pending
func (s *Session) GetPendingOffers() ([]*Offer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/pending",
			s.Auth.PortalEndpoint,
			offerEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllOffers(s, req)
}

//GetCompletedOffers returns a list of offers where status is complete
func (s *Session) GetCompletedOffers() ([]*Offer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/completed",
			s.Auth.PortalEndpoint,
			offerEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllOffers(s, req)
}

//GetOfferByID returns a single offer given the offer ID
func (s *Session) GetOfferByID(id string) (*Offer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			offerEndpoint,
			id),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}

//CreateOffer POSTs a new offer and reserves the relevant listing
func (o *Offer) CreateOffer(s *Session) (*Offer, error) {
	o.ID = ""
	body, err := json.Marshal(*o)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			offerEndpoint,
		),
		bytes.NewReader(body),
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}

//UpdateOffer PUTs new data to an existing offer, using an offer ID
func (o *Offer) UpdateOffer(s *Session) (*Offer, error) {
	body, err := json.Marshal(*o)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			offerEndpoint,
			o.ID),
		bytes.NewReader(body),
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}

//CompleteOffer PUTs new data to an existing offer, changing the offer status to complete and the relevant listing to sold.
func (o *Offer) CompleteOffer(s *Session) (*Offer, error) {
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v/compelete",
			s.Auth.PortalEndpoint,
			offerEndpoint,
			o.ID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}

// //TODO: Unsure how this is meant to work
// //AttachDoc POSTs an attachment to the offer
// func (o *Offer) AttachDoc(s *Session) (*Offer, error) {
// 	req, err := http.NewRequest(
// 		"POST",
// 		fmt.Sprintf(
// 			"%v/%v/%v/attachments",
// 			s.Auth.PortalEndpoint,
// 			offerEndpoint,
// 			o.ID),
// 		nil,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return execRequestReturnSingleAttachment(s, req)
// }

// //TODO: Unsure how this is meant to work
// //DownloadOfferAttachment downloads a specific attachment on an offer give the attachment & offer id
// func DownloadOfferAttachment() {
// }

//CancelOffer sets an offer's status to cancelled and reverts a listing to available
func (o *Offer) CancelOffer(s *Session) (*Offer, error) {
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v/cancel",
			s.Auth.PortalEndpoint,
			offerEndpoint,
			o.ID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}
