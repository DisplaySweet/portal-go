package portal

const eventScheduleEndpoint = "eventschedules"

//TODO: This is EOI related, return to complete at a later date

type Schedule struct {
	ID           string `json:"id"`
	EventID      string `json:"eventId"`
	EventDateID  string `json:"eventDateId"`
	EventDate    EventDate
	ScheduleTime string `json:"scheduleTime"`
}

//func (s *Session) GetEventSchedules()
