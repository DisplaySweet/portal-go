package portal

// Session holds information about the current session
type Session struct {
	Auth              Auth
	ProjectID         string
	CompanyID         string
	DumpErrorPayloads bool
}
