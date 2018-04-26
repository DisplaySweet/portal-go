package portal

//ListingActivity handles the returned JSON listing activity
type ListingActivity struct {
	ID           string `json:"id"`
	Timestamp    string `json:"timestamp"`
	User         string `json:"user"`
	ChangedField string `json:"changedfield"`
	// ChangedValue object? `json:"changedvalue"`
}

//ListingStatusActivity handles the specifically returned JSON listing status activity
type ListingStatusActivity struct {
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	User      string `json:"user"`
	// Status object? `json:"status"`
}
