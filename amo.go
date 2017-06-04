package goAmoCrm

import (
	"github.com/emericanec/goAmoCrm/structures"
	"net/url"
	"strings"
	"net/http"
	"errors"
)

const (
	HTTP_OK = 200
	CONTACTS_PATH = "/private/api/v2/json/contacts"
)


type Amo struct {
	Config   structures.Configure
	Contacts structures.Contacts
}

func auth(configure structures.Configure) ([]*http.Cookie, error) {
	var cookie []*http.Cookie
	values := url.Values{}
	values.Set("USER_LOGIN", configure.Login)
	values.Set("USER_HASH", configure.Hash)
	body := strings.NewReader(values.Encode())
	urlString := configure.Url + "/private/api/auth.php"
	response, err := http.Post(urlString, "application/x-www-form-urlencoded", body)
	if response.StatusCode == HTTP_OK {
		cookie = response.Cookies()
	} else {
		err = errors.New("Wrong http status: " + string(response.StatusCode))
	}
	return cookie, err
}

func New(url, login, hash string) Amo {
	var err error

	config := structures.Configure{
		Url:   url,
		Login: login,
		Hash:  hash,
	}

	structures.Cookie, err = auth(config)
	if err != nil {
		panic(err.Error())
	}

	contacts := structures.Contacts{
		Config: config,
		Path:   CONTACTS_PATH,
	}

	amo := Amo{
		config,
		contacts,
	}

	return amo
}
