package portal

//Address holds any address information
type Address struct {
	ID           string `json:"id"`
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	City         string `json:"city"`
	Country      string `json:"country"`
	Postcode     string `json:"postcode"`
	State        string `json:"state"`
}

type BillingAddress struct {
	ID       string `json:"id"`
	Street   string `json:"Street"`
	City     string `json:"city"`
	Country  string `json:"country"`
	Postcode string `json:"PostalCode"`
	State    string `json:"state"`
}
