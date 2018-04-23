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
	Name          string
	Availability  string
	LotNumbers    string
	Price         float32
	OriginalPrice float32
	Bedrooms      string
	Bathrooms     string
	Carspaces     string
	InternalArea  float32
	ExternalArea  float32
	TotalArea     float32
}

// execute the HTTP requests and get the single Listing that should come out
func execRequestReturnSingleListing(s *Session, req *http.Request) (*Listing, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		return nil, err
	}

	listing := &Listing{}
	err = json.Unmarshal(responseBytes, listing)

	return listing, err
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
