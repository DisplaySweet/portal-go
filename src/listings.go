package portal

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const listingEndpoint = "listings"

// Listing holds information of a listing
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
}

type Listings struct {
	Listings []*Listing
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
	}

	return v, nil
}

//WIP
func execRequestReturnListings(s *Session, req *http.Request) (*Listings, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	//What the response should look like.
	// {[
	// 		ID [fields:{data:val, data:val}]
	// 		ID [fields:{data:val, data:val}]
	//	 	ID [fields:{data:val, data:val}]
	// ]}

	//	first unmarshal the raw json into a valid interface
	//	then iterate through one result at a time
	//	the structure should be the same as a listing, just an array
	var temp []map[string]*Listing

	//	unmarshal each listing into a listing struct,
	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		return nil, err
	}

	var list *Listings
	var listing *Listing
	//	squeeze the leading ID into the listing struct
	// 	then append them to a list and push them into the listings struct
	for index, entry := range temp {
		for id, listing := range entry {
			listing.ID = id
			list.Listings = append(list.Listings, listing)
		}
	}

	//need to unpack the leading string into the listing ID

	return list, nil

}

// GetListingByID gets a single listing by its ID
func (s *Session) GetListingByID(id string) (*Listing, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			listingEndpoint,
			id,
		),
		nil)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleListing(s, req)
}

// GetListingByName gets a single listing by its name
func (s *Session) GetListingByName(name string) (*Listing, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v?name=%v",
			s.Auth.PortalEndpoint,
			listingEndpoint,
			name,
		),
		nil,
	)

	if err != nil {
		return nil, err
	}

	return execRequestReturnSingleListing(s, req)
}

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

// SendUpdate saves changes to listing
func (l *Listing) SendUpdate(s *Session) error {
	body, err := json.Marshal(*l)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			listingEndpoint,
			l.ID,
		),
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
