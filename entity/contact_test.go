package entity

import (
	"testing"
	"goAmoCrm"
)

func TestContact_Save(t *testing.T) {
	var config = goAmoCrm.Config{
		"yourdomain.amocrm.ru",
		"yourLogin",
		"yourApiKey",
	}

	var handler = goAmoCrm.Handler{
		config,
	}

	token, _ := handler.Auth()

	var contact = Contact{}
	contact.Name = "Walter White"
	contact.CompanyName = "ЭБК"
	contact.Save(handler, token.Token)
}