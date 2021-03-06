package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const offerEndpoint = "offers"
const GetOfferByExtIdEndpoint = offerEndpoint + "/getByExternalId"

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
	Price                  string
	FundsReceived          float64
	Status                 string
	OfferStatusChangedDate string
	s                      *Session
}

func execRequestReturnAllOffers(s *Session, req *http.Request) ([]*Offer, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var temp []*Offer

	list := make([]*Offer, 0, 0)

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		log.Println(string(responseBytes))
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	for _, offer := range temp {
		offer.s = s
		list = append(list, offer)
	}

	return list, nil
}

func execRequestReturnSingleOffer(s *Session, req *http.Request) (*Offer, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	offer := &Offer{}
	err = json.Unmarshal(responseBytes, offer)
	offer.s = s

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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}

//GetOfferByID returns a single offer given the offer ID
func (s *Session) GetOfferByExtID(extId string) (*Offer, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			GetOfferByExtIdEndpoint,
			extId),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}

//CreateOffer POSTs a new offer and reserves the relevant listing
func (s *Session) CreateOffer(o *Offer) (*Offer, error) {
	o.ID = ""
	body, err := json.Marshal(*o)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleOffer(s, req)
}

//UpdateOffer PUTs new data to an existing offer, using an offer ID
func (o *Offer) Update() (*Offer, error) {
	body, err := json.Marshal(*o)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			o.s.Auth.PortalEndpoint,
			offerEndpoint,
			o.ID),
		bytes.NewReader(body),
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleOffer(o.s, req)
}

//CompleteOffer PUTs new data to an existing offer, changing the offer status to complete and the relevant listing to sold.
func (o *Offer) Complete() (*Offer, error) {
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v/compelete",
			o.s.Auth.PortalEndpoint,
			offerEndpoint,
			o.ID),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleOffer(o.s, req)
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
func (o *Offer) Cancel() (*Offer, error) {
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v/cancel",
			o.s.Auth.PortalEndpoint,
			offerEndpoint,
			o.ID),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleOffer(o.s, req)
}
