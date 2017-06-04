package goAmoCrm

import (
	"testing"
	"fmt"
	"github.com/emericanec/goAmoCrm/entity"
)

func TestRequest_CreateGetRequest(t *testing.T) {
	amo := New("https://new5933c838c4aa8.amocrm.ru", "sale@roque.ru", "f2e1e217146b1d59e4d26c5c55c8aed5")

	contact := entity.Contact{}
	contact.Name = "James Bond"
	result := amo.Contacts.Save(contact)
	fmt.Println(result)

	params := entity.ContactListParam{}
	contacts := amo.Contacts.List(params)
	fmt.Println(contacts)


}