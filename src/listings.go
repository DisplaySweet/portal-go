package portal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const listingEndpoint = "listings"

//Listing holds information of a listing
type Listing struct {
	ID            string
	Name          string  `json:"listing_name"`
	Availability  string  `json:"availability"`
	Floor         string  `json:"floor"`
	Building      string  `json:"building"`
	Price         float32 `json:"live_price"`
	OriginalPrice float32 `json:"price"`
	Bedrooms      string  `json:"bedrooms"`
	Bathrooms     string  `json:"bathrooms"`
	Study         string  `json:"study"`
	Carspaces     string  `json:"carspaces"`
	Aspect        string  `json:"aspect_orientation"`
	MarketingPlan string  `json:"marketing_plan"`
	InternalArea  float32 `json:"internal_area"`
	ExternalArea  float32 `json:"external_area"`
	TotalArea     float32 `json:"total_area"`
	s             *Session
}

// execute the HTTP requests and get the single Listing that should come out
func execRequestReturnSingleListing(s *Session, req *http.Request) (*Listing, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Listing

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	var v *Listing
	for k, v := range temp {
		v.ID = k
		v.s = s
	}

	return v, nil
}

//WIP
func execRequestReturnListings(s *Session, req *http.Request) ([]*Listing, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*Listing

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	list := make([]*Listing, 0, 0)

	for id, listing := range temp {
		listing.ID = id
		listing.s = s
		list = append(list, listing)
	}

	return list, nil
}

func execRequestReturnActivity(s *Session, req *http.Request) (*ListingActivity, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	activity := &ListingActivity{}

	err = json.Unmarshal(responseBytes, activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func execRequestReturnAllActivity(s *Session, req *http.Request) ([]*ListingActivity, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*ListingActivity

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	list := make([]*ListingActivity, 0, 0)

	for _, action := range temp {
		list = append(list, action)
	}

	return list, nil
}

func execRequestReturnAllStatusActivity(s *Session, req *http.Request) ([]*ListingStatusActivity, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	var temp map[string]*ListingStatusActivity

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	list := make([]*ListingStatusActivity, 0, 0)

	for _, action := range temp {
		list = append(list, action)
	}

	return list, nil
}

// // GetListingByID gets a single listing by its ID
// func (s *Session) GetListingByID(id string) (*Listing, error) {
// 	req, err := http.NewRequest(
// 		"GET",
// 		fmt.Sprintf(
// 			"%v/%v/%v",
// 			s.Auth.PortalEndpoint,
// 			listingEndpoint,
// 			id,
// 		),
// 		nil)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return execRequestReturnSingleListing(s, req)
// }

// // GetListingByName gets a single listing by its name
// func (s *Session) GetListingByName(name string) (*Listing, error) {
// 	req, err := http.NewRequest(
// 		"GET",
// 		fmt.Sprintf(
// 			"%v/%v?name=%v",
// 			s.Auth.PortalEndpoint,
// 			listingEndpoint,
// 			name,
// 		),
// 		nil,
// 	)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return execRequestReturnSingleListing(s, req)
// }

//GetListings returns a slice of all Listing
func (s *Session) GetListings() ([]*Listing, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			listingEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnListings(s, req)
}

//GetActivityByID GETs all the activity data for a particular listing (using ID)
func (s *Session) GetActivityByID(l *Listing) (*ListingActivity, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v/activity",
			s.Auth.PortalEndpoint,
			listingEndpoint,
			l.ID),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnActivity(s, req)
}

//GetAllActivity GETs all the activity for all portal listings
func (s *Session) GetAllActivity() ([]*ListingActivity, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/activity",
			s.Auth.PortalEndpoint,
			listingEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllActivity(s, req)
}

//GetAllStatusActivity GETs all listing status activity for every portal listing
func (s *Session) GetAllStatusActivity() ([]*ListingStatusActivity, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/activity/statuses",
			s.Auth.PortalEndpoint,
			listingEndpoint),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnAllStatusActivity(s, req)
}

//CreateListing POSTs a new Listing to the portal
func (s *Session) CreateListing(l *Listing) error {
	l.ID = "" // Make sure to blank out the ID
	body, err := json.Marshal(*l)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v",
			s.Auth.PortalEndpoint,
			listingEndpoint),
		bytes.NewReader(body),
	)

	if err != nil {
		return err
	}

	response, err := executeRequest(s, req)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 200:
	case 204:
		break
	default:
		return errors.New("Did not get a success code from the portal")
	}

	return nil
}

//DeleteListings removes the referenced listing from Listings
func (l *Listing) Delete() (int, error) {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			l.s.Auth.PortalEndpoint,
			listingEndpoint,
			l.ID,
		),
		nil,
	)
	if err != nil {
		return 0, err
	}

	return executeRequestAndGetStatusCode(l.s, req)
}

//I guess because []*Listing isnt a defined struct itself, we cant do this?

//WIP//

// func (l []*Listing) BulkUpdate(s *Session) err {
// 	body, err := json.Marshal(*l)
// 	if err != nil {
// 		return err
// 	}

// 	req, err := http.NewRequest(
// 		"PUT",
// 		fmt.Sprintf(
// 			"%v/%v",
// 			s.Auth.PortalEndpoint,
// 			listingEndpoint,
// 		),
// 		bytes.NewReader(body),
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	response, err := executeRequest(s, req)
// 	if err != nil {
// 		return err
// 	}

// 	switch response.StatusCode {
// 	case 200:
// 	case 204:
// 		break
// 	default:
// 		return errors.New("Did not get a success code from the portal")
// 	}

// 	return nil
// }

//I guess because []*Listing isnt a defined struct itself, we cant do this?

//WIP

// func (l []*Listing) BulkUpdateByRole(s *Session) err {
// 	body, err := json.Marshal(*l)
// 	if err != nil {
// 		return err
// 	}

// 	req, err := http.NewRequest(
// 		"PUT",
// 		fmt.Sprintf(
// 			"%v/%v/bulkbyrole",
// 			s.Auth.PortalEndpoint,
// 			listingEndpoint,
// 		),
// 		bytes.NewReader(body),
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	response, err := executeRequest(s, req)
// 	if err != nil {
// 		return err
// 	}

// 	switch response.StatusCode {
// 	case 200:
// 	case 204:
// 		break
// 	default:
// 		return errors.New("Did not get a success code from the portal")
// 	}

// 	return nil
// }

// SendUpdate saves changes to listing
func (l *Listing) Update() error {
	body, err := json.Marshal(*l)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v/byrole",
			l.s.Auth.PortalEndpoint,
			listingEndpoint,
			l.ID,
		),
		bytes.NewReader(body),
	)

	if err != nil {
		return err
	}

	response, err := executeRequest(l.s, req)
	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 200:
	case 204:
		break
	default:
		return errors.New("Did not get a success code from the portal")
	}

	return nil
}
