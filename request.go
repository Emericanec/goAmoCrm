package goAmoCrm

import (
	"net/url"
	"fmt"
	"io"
)

type Object struct {
	objectName string
	methodName string
}

type Request struct {
	Object Object
	Body io.Reader
	Post bool
	Url string
}

func (this *Request) CreateAuthRequest()  {
	this.Post = true
	this.Url = "auth.php?type=json"
}

func (this *Request) CreateInfoRequest() {
	this.Post = false
	this.Url = "v2/json/accounts/current"
}

func (this *Request) CreateGetRequest(params map[string]string) {
	this.Post = false
	this.Url = "v2/json/"+this.Object.objectName + "/" + this.Object.methodName
	if(len(params) > 0) {
		query := url.Values{}
		for k, v := range params{
			query.Set(k, v)
		}
		this.Url += "?" + query.Encode()
	}
	fmt.Println(this.Url)
}