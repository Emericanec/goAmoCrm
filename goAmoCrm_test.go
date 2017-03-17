package goAmoCrm

import (
	"testing"
	"fmt"
)

func TestRequest_CreateGetRequest(t *testing.T) {
	var config = Config{
		"yourdomain.amocrm.ru",
		"yourLogin",
		"yourApiKey",
	}

	var handler = Handler{
		config,
	}

	token, _ := handler.Auth()
	fmt.Println(token.Token)
}