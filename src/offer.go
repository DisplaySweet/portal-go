package portal

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

//GetAllOffers returns a list of Offers
func GetAllOffers() {

}

//GetPendingOffers returns a list of offers where status is pending
func GetPendingOffers() {

}

//GetCompletedOffers returns a list of offers where status is complete
func GetCompletedOffers() {

}

//GetOfferByID returns a single offer given the offer ID
func GetOfferByID() {

}

//CreateOffer POSTs a new offer and reserves the relevant listing
func CreateOffer() {

}

//UpdateOffer PUTs new data to an existing offer, using an offer ID
func UpdateOffer() {

}

//CompleteOffer PUTs new data to an existing offer, changing the offer status to complete and the relevant listing to sold.
func CompleteOffer() {

}

//AttachDoc POSTs an attachment to the offer
func AttachDoc() {

}

//DownloadOfferAttachment downloads a specific attachment on an offer give the attachment & offer id
func DownloadOfferAttachment() {
}

//CancelOffer sets an offer's status to cancelled and reverts a listing to available
func CancelOffer() {

}
