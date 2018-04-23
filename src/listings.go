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
	Name          string   `json:"listing_name"`
	Availability  string   `json:"availability"`
	Floor         string   `json:"floor"`
	Building      string   `json:"building"`
	LotNumbers    []string `json:"lots"`
	Price         float32  `json:"live_price"`
	OriginalPrice float32  `json:"price"`
	Bedrooms      string   `json:"bedrooms"`
	Bathrooms     string   `json:"bathrooms"`
	Study         string   `json:"study"`
	Carspaces     string   `json:"carspaces"`
	Aspect        string   `json:"aspect_orientation"`
	MarketingPlan string   `json:"marketing_plan"`
	InternalArea  float32  `json:"internal_area"`
	ExternalArea  float32  `json:"external_area"`
	TotalArea     float32  `json:"total_area"`
}

// execute the HTTP requests and get the single Listing that should come out
func execRequestReturnSingleListing(s *Session, req *http.Request) (*Listing, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	listingID, _ := setListingID(responseBytes)

	listing := &Listing{}
	err = json.Unmarshal(responseBytes, listing)
	listing.ID = listingID

	return listing, err
}

//https://stackoverflow.com/questions/17452722/how-to-get-the-key-value-from-a-json-string-in-go
func setListingID(responseBytes []byte) (string, error) {
	// a map container to decode the JSON structure into
	c := make(map[string]interface{})

	// unmarshal JSON
	e := json.Unmarshal(responseBytes, &c)

	// panic on error
	if e != nil {
		panic(e)
	}

	// a string slice to hold the keys
	k := make([]string, len(c))

	// iteration counter
	i := 0

	// copy c's keys into k
	for s, _ := range c {
		k[i] = s
		i++
	}

	//just want the parent key
	return k[0], nil
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
