package portal

// Listing holds information of a listing
type Listing struct {
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

// GetListingByID gets a single listing by its ID
func (s *Session) GetListingByID(id string) (*Listing, error) {
	return nil, nil
}

// GetListingByName gets a single listing by its name
func (s *Session) GetListingByName(name string) (*Listing, error) {
	return nil, nil
}

// SendUpdate saves changes to listing
func (l *Listing) SendUpdate() error {
	return nil
}
