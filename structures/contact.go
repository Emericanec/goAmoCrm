package structures

import "github.com/emericanec/goAmoCrm/entity"

type Contacts struct {
	Config Configure
	Path string
}

func (c *Contacts) getUrl() string {
	return c.Config.Url + c.Path
}

func (c *Contacts) List(params entity.ContactListParam) []entity.Contact {
	url := c.getUrl() + "/list"
	tmp := []entity.Contact{}
	return tmp
}

func (c *Contacts) Save(contact entity.Contact) int {
	return 1
}