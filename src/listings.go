package portal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const listingEndpoint = "listings"
const GetListingByExtIdEndpoint = listingEndpoint + "/getByExternalId"

type Listing struct {
	ID            string  `json:"id"`
	ExternalID    string  `json:"externalid"`
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
	LotNumber     int     `json:"Lot_Number__c"`
	S             Session `json:"S"`
}

type ExportListing struct {
	ID            ID            `json:"id"`
	ExternalID    ExternalID    `json:"externalid"`
	Name          Name          `json:"Name"`
	Availability  Availability  `json:"Sale_Stage__c"`
	Floor         Floor         `json:"ds_rep_level"`
	Building      Building      `json:"Building__r"`
	Price         Price         `json:"ds_live_price"`
	OriginalPrice OriginalPrice `json:"Sales_Price__c"`
	MarketingPlan MarketingPlan `json:"marketing_plan"`
	Bedrooms      Bedrooms      `json:"Bedrooms__c"`
	Bathrooms     Bathrooms     `json:"Bathrooms__c"`
	Carspaces     Carspaces     `json:"Parking_Allocations__c"`
	Aspect        Aspect        `json:"Aspect__c"`
	InternalArea  InternalArea  `json:"Interior_Space__c"`
	ExternalArea  ExternalArea  `json:"Exterior_Space__c"`
	TotalArea     TotalArea     `json:"Total_Space_M__c"`
	LotNumber     LotNumber     `json:"Lot_Number__c"`
	S             Session       `json:"S"`
	// Study         Study         `json:"Has_Study__c"`
}

type ID struct {
	String string `json:"string"`
}
type ExternalID struct {
	String string `json:"string"`
}
type Name struct {
	String string `json:"string"`
}
type Availability struct {
	String string `json:"string"`
}
type Floor struct {
	String string `json:"string"`
}
type Building struct {
	String string `json:"string"`
}
type Bedrooms struct {
	String string `json:"string"`
}
type Bathrooms struct {
	String string `json:"string"`
}

// type Study struct {
// 	String string `json:"string"`
// }
type Carspaces struct {
	String string `json:"string"`
}
type Aspect struct {
	String string `json:"string"`
}
type Price struct {
	Float float32 `json:"float"`
}
type OriginalPrice struct {
	Float float32 `json:"float"`
}
type MarketingPlan struct {
	String string `json:"string"`
}
type InternalArea struct {
	Float float32 `json:"float"`
}
type ExternalArea struct {
	Float float32 `json:"float"`
}
type TotalArea struct {
	Float float32 `json:"float"`
}
type LotNumber struct {
	Int int `json:"int"`
}
type OpportunityPayload struct {
	ID            string `json:"id"`
	StageName     string `json:"StageName"`
	ListingStatus string `json:"ListingStatus"`
}

// execute the HTTP requests and get the single Listing that should come out
func execRequestReturnSingleListing(s *Session, req *http.Request) (*Listing, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var temp []*Listing

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var v *Listing
	for _, v := range temp {
		v.S = *s
	}

	return v, nil
}

//WIP
func execRequestReturnListings(s *Session, req *http.Request) ([]*Listing, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var temp []*Listing

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	list := make([]*Listing, 0, 0)

	for _, listing := range temp {
		listing.S = *s
		list = append(list, listing)
	}

	return list, nil
}

// execute the HTTP requests and get the single Listing that should come out
func execRequestReturnSingleisting(s *Session, req *http.Request) (*Listing, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var listing *Listing
	err = json.Unmarshal(responseBytes, &listing)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return listing, err
}

func execRequestReturnActivity(s *Session, req *http.Request) (*ListingActivity, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	activity := &ListingActivity{}

	err = json.Unmarshal(responseBytes, activity)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return activity, nil
}

func execRequestReturnAllActivity(s *Session, req *http.Request) ([]*ListingActivity, error) {
	responseBytes, err := executeRequestAndGetBodyBytes(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var temp map[string]*ListingActivity

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	var temp map[string]*ListingStatusActivity

	err = json.Unmarshal(responseBytes, &temp)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	list := make([]*ListingStatusActivity, 0, 0)

	for _, action := range temp {
		list = append(list, action)
	}

	return list, nil
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnListings(s, req)
}

//GetListings returns a slice of all Listing
func (s *Session) GetListingByExtId(extId string) (*Listing, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%v/%v/%v",
			s.Auth.PortalEndpoint,
			GetListingByExtIdEndpoint,
			extId),
		nil,
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnSingleisting(s, req)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return nil, err
	}

	return execRequestReturnAllStatusActivity(s, req)
}

//CreateListing POSTs a new Listing to the portal
func (s *Session) CreateListing(l *ExportListing) (error, string) {
	l.ID.String = "" // Make sure to blank out the ID
	body, err := json.Marshal(*l)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err, ""
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
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err, ""
	}

	response, err := executeRequest(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err, ""
	}

	switch response.StatusCode {
	case 200:
	case 204:
		break
	default:
		return fmt.Errorf("Error in file: %v line %v. Original ERR: Did not get a success code from the portal: %v", ErrorFile(), ErrorLine(), response.StatusCode), ""
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	id := string(bodyBytes)
	id = strings.Replace(id, "\"", "", 2)

	return nil, id
}

//UpdateListingContractStatus POSTs a new opportunity status to the portal
func (s *Session) UpdateListingContractStatus(payload OpportunityPayload, listingId string, prospectId string) (error, string) {
	body, err := json.Marshal(payload)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err, ""
	}
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(
			"%v/%v/%v/updatefromwebhook/%v",
			s.Auth.PortalEndpoint,
			listingEndpoint,
			listingId,
			prospectId),
		bytes.NewReader(body),
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err, ""
	}

	response, err := executeRequest(s, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err, ""
	}

	switch response.StatusCode {
	case 200:
	case 204:
		break
	default:
		return fmt.Errorf("Error in file: %v line %v. Original ERR: Did not get a success code from the portal: %v", ErrorFile(), ErrorLine(), response.StatusCode), ""
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	id := string(bodyBytes)
	id = strings.Replace(id, "\"", "", 2)

	return nil, id
}

//DeleteListings removes the referenced listing from Listings
func (l *Listing) Delete() error {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf(
			"%v/%v/%v",
			l.S.Auth.PortalEndpoint,
			listingEndpoint,
			l.ID,
		),
		nil,
	)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	return executeRequestAndParseStatusCode(&l.S, req)
}

// SendUpdate saves changes to listing
func (l *Listing) Update() error {
	body, err := json.Marshal(*l)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"%v/%v/%v/byrole",
			l.S.Auth.PortalEndpoint,
			listingEndpoint,
			l.ID,
		),
		bytes.NewReader(body),
	)

	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}

	err = executeRequestAndParseStatusCode(&l.S, req)
	if err != nil {
		err = fmt.Errorf("Error in file: %v line %v. Original ERR: %v", ErrorFile(), ErrorLine(), err)
		return err
	}
	return nil
}
