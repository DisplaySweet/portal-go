package portal

// Session holds information about the current session
type Session struct {
	Auth    Auth
	Project Project
	Company Company
}
