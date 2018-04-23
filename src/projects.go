package portal

// Project holds information about a project
type Project struct {
	ID   string
	Name string
}

// GetProjectByName returns a project queried by name
func (s *Session) GetProjectByName(name string) (*Project, error) {
	return nil, nil
}

// GetProjectByID returns a project queried by ID
func (s *Session) GetProjectByID(id string) (*Project, error) {
	return nil, nil
}
