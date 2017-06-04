package structures

import (
	"github.com/emericanec/goAmoCrm/entity"
	"encoding/json"
)

type Contacts struct {
	Config Configure
	Path string
}

func (c *Contacts) getUrl() string {
	return c.Config.Url + c.Path
}

func (c *Contacts) List(params entity.ContactListParam) []entity.Contact {
	res := DoGet(c.getUrl() + "/list", map[string]string{})
	data := &entity.ContactListResponseRoot{}
	err := json.Unmarshal(res, data)
	if err != nil {
		panic(err.Error())
	}
	return data.Response.Contacts
}

func (c *Contacts) Save(contact entity.Contact) entity.Contact {
	var requestRoot entity.ContactsSetRequestRoot
	if contact.Id == 0 {
		requestRoot = c.add(contact)
	} else {
		requestRoot = c.update(contact)
	}

	res := DoPost(c.getUrl() + "/set", requestRoot)
	response := &entity.ContactsSetResponseRoot{}
	err := json.Unmarshal(res, response)
	if err != nil {
		panic(err.Error())
	}
	if contact.Id == 0 {
		for _, add := range response.Response.Contacts.Add {
			contact.Id = add.Id
		}
	}

	return contact
}

func (c *Contacts) add(contact entity.Contact) entity.ContactsSetRequestRoot {
	requestRoot := entity.ContactsSetRequestRoot{}
	requestRoot.Request.Contacts.Add = append(requestRoot.Request.Contacts.Add, contact)
	return requestRoot
}

func (c *Contacts) update(contact entity.Contact) entity.ContactsSetRequestRoot {
	requestRoot := entity.ContactsSetRequestRoot{}
	requestRoot.Request.Contacts.Update = append(requestRoot.Request.Contacts.Update, contact)
	return requestRoot
}
