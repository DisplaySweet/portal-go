package portal

type Agent struct {
	ID           string `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Active       bool   `json:"active"`
	Superuser    bool   `json:"superuser"`
	LastLoggedIn string `json:"lastloggedin"`
	CreatedDate  string `json:"createddate"`
	ArchivedOn   string `json:"archivedon"`
}
