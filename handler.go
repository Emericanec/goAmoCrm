package goAmoCrm

import (
	"net/http"
	"errors"
	"strings"
	"net/url"
	"io"
	"io/ioutil"
	"encoding/json"
)

const OK = 200;

type Token struct {
	Token string
}

type Config struct {
	Domain    string
	UserLogin string
	UserHash  string
}

type Handler struct {
	Config Config
}

func (this Handler) Auth() (Token, error) {
	var token = Token{}
	var request = Request{}
	request.CreateAuthRequest()
	values := url.Values{}
	values.Set("USER_LOGIN", this.Config.UserLogin)
	values.Set("USER_HASH", this.Config.UserHash)
	request.Body = strings.NewReader(values.Encode())
	response := this.Request(request)

	body, err := ioutil.ReadAll(response.Body)
	if(err != nil){
		panic(err.Error())
	}
	var s interface{}
	err = json.Unmarshal(body, &s)
	if(err != nil){
		panic(err.Error())
	}
	if(response.StatusCode == OK){
		cookieString := response.Header["Set-Cookie"][0]
		if(cookieString != "") {
			stringArray := strings.Split(cookieString, ";")
			sessionIdArray := strings.Split(stringArray[0], "=")
			token.Token = sessionIdArray[1]
		}
	} else {
		err = errors.New("Wrong http status: " + string(response.StatusCode))
	}

	return token, err
}

func (this Handler) Request(request Request) *http.Response {
	var urlString string

	urlString = "https://" + this.Config.Domain + "/private/api/" + request.Url


	var response *http.Response

	if (request.Post) {
		response = this.doPost(urlString, request.Body)
	} else {
		response = this.doGet(urlString)
	}

	return response
}

func (this Handler) doPost(url string, body io.Reader) *http.Response {
	response, err := http.Post(url, "application/x-www-form-urlencoded", body)
	if(err != nil){
		panic(err.Error())
	}
	return response
}

func (this Handler) doGet(url string) *http.Response {
	response, err := http.Get(url)
	if(err != nil){
		panic(err.Error())
	}
	return response
}
