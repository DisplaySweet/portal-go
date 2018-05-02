package portal

const eventScheduleEndpoint = "eventschedules"

//TODO: This is EOI related, return to complete at a later date

type Schedule struct {
	ID           string
	EventID      string
	EventDateID  string
	EventDate    EventDate
	ScheduleTime string
}

//func (s *Session) GetEventSchedules()
