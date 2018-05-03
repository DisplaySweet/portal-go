package portal

// UserCompany holds information about a UserCompany
type UserCompany struct {
	ID          string `json:"Id"`
	UserID      string `json:"user"`
	User        Agent
	CompanyID   string `json:"companyid"`
	Company     Company
	Level       uint64 `json:"level"`
	Role        string `json:"role"`
	CreatedDate string `json:"createddate"`
}
