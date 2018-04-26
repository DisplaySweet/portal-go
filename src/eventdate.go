package portal

type EventDate struct {
	ID          string
	StartDate   string
	EndDate     string
	TimeWindow  string
	BreakLength string
	Schedules   []Schedule
	CompanyID   string
	ProjectID   string
	EventID     string
}
