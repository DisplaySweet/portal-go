package portal

// Event holds information about an event
type Event struct {
	ID          string
	Name        string
	CompanyID   string
	Company     Company
	ProjectID   string
	Project     Project
	Terminals   int32
	CreatedDate string
	TZInfo      string
	Dates       []EventDate
	ArchivedOn  string
}

//GetEvents GETs a list of all Events
func GetEvents() {

}

//GetEventByID GETs an individual event using the event ID
func GetEventByID() {

}

//CreateEvent POSTs new event data to the endpoint
func CreateEvent() {

}

//UpdateEventByID PUTs new data to an existing event using it's ID
func UpdateEventByID() {

}

//DeleteEventByID DELETEs an existing event using it's ID
func DeleteEventByID() {

}
