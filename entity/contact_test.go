package entity

import (
	"testing"
	"goAmoCrm"
)

func TestContact_Save(t *testing.T) {
	var config = goAmoCrm.Config{
		"ebc.amocrm.ru",
		"dinukov@exbico.ru",
		"a89b3f1977893ae47556caf46acc800a",
	}

	var handler = goAmoCrm.Handler{
		config,
		goAmoCrm.Token{},
	}

	handler.Auth()

	var contact = Contact{}
	contact.Name = "Walter White"
	contact.CompanyName = "ЭБК"
	contact.Save(handler)
}