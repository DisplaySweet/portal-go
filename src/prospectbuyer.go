package portal

type ProspectBuyer struct {
	ID                 string `json:"id"`
	Phone              string `json:"phone"`
	Mobile             string `json:"mobile"`
	Email              string `json:"email"`
	Firstname          string `json:"firstname"`
	Lastname           string `json:"lastname"`
	DOB                string `json:"dob"`
	Nationality        string `json:"nationality"`
	AddressID          string `json:"addressid"`
	Address            Address
	ProspectID         string `json:"prospectid"`
	Prospect           Prospect
	CompanyID          string `json:"company"`
	ProjectID          string `json:"project"`
	FIRB               int    `json:"firb"`
	Primary            bool   `json:primary"`
	OriginContactUser  string `json:"origincontactuser"`
	OriginatingContact string `json:"originatingcontact"`
	OriginAccountUser  string `json:"originaccountuser"`
	OriginatingAccount string `json:"originatingaccount"`
}

//GetProspectBuyers GETs a list of ProspectBuyers that belong to the session
func GetProspectBuyers() {

}

//GetAvailableProspectBuyers GETs a list of ProspectBuyers where none of their Offers have been cancelled that belong to the session
func GetAvailableProspectBuyers() {

}
