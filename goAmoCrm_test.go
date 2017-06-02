package goAmoCrm

import (
	"testing"
	"fmt"
	"github.com/emericanec/goAmoCrm/entity"
)

func TestRequest_CreateGetRequest(t *testing.T) {
	amo := New("domain.amocrm.ru", "email@gmail.com", "YOUR TOKEN")
	params := entity.ContactListParam{}
	contacts := amo.Contacts.List(params)
	fmt.Println(contacts)
}