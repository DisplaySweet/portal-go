package portal

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
}

//GetAllEventProspects GETs a list of all prospects for an event
func GetAllEventProspects() {

}

//UpdateProspect PUTs updates a prospect with new data, using the prospect object
func UpdateProspect() {

}

//UpdateProspectAgent POSTs a newly assigned Agent to the Prospect
func UpdateProspectAgent() {

}
