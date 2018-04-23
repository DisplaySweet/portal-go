package portal

// Contact holds information of a contact
type Contact struct {
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

// GetContactByID gets a single contact by its ID
func (s *Session) GetContactByID(id string) (*Contact, error) {
	return nil, nil
}

// GetContactByName gets the first contact that matches the provided name
func (s *Session) GetContactByName(firstname string, lastname string) (*Contact, error) {
	return nil, nil
}

// SendUpdate saves changes made to contact
func (c *Contact) SendUpdate() error {
	return nil
}
