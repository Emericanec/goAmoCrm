package goAmoCrm

import (
	"github.com/emericanec/goAmoCrm/structures"
	"net/url"
	"strings"
	"net/http"
	"errors"
)

const OK = 200

type Amo struct {
	Config   structures.Configure
	Contacts structures.Contacts
}

func auth(configure structures.Configure) (string, error) {
	var token string
	var err error
	values := url.Values{}
	values.Set("USER_LOGIN", configure.Login)
	values.Set("USER_HASH", configure.Hash)
	body := strings.NewReader(values.Encode())
	urlString := "https://" + configure.Url + "/private/api/auth.php"
	var response *http.Response
	response, err = http.Post(urlString, "application/x-www-form-urlencoded", body)
	if response.StatusCode == OK {
		cookieString := response.Header["Set-Cookie"][0]
		if cookieString != "" {
			stringArray := strings.Split(cookieString, ";")
			sessionIdArray := strings.Split(stringArray[0], "=")
			token = sessionIdArray[1]
		}
	} else {
		err = errors.New("Wrong http status: " + string(response.StatusCode))
	}
	return token, err
}

func New(url, login, hash string) Amo {

	config := structures.Configure{
		Url:   url,
		Login: login,
		Hash:  hash,
	}

	token, err := auth(config)
	if err != nil {
		panic(err.Error())
	}

	config.Token = token

	contacts := structures.Contacts{
		Config: config,
		Path:   "/private/api/v2/json/contacts",
	}

	amo := Amo{
		config,
		contacts,
	}
	return amo
}
