package entity

import (
	"goAmoCrm"
	"encoding/json"
	"fmt"
)

type CustomFiledValue struct {
	value string
	enum string
	subtype string
}

type CustomFiled struct {
	id int
	values []CustomFiledValue
}

type Contact struct {
	Id int
	LastModified string
	Name string
	CompanyName string
	ResponsibleUserId int
	Tags string
	LinkedLeadsId []int
	CustomFields []CustomFiled
}

func (this Contact) Save(handler goAmoCrm.Handler, token string)  {

	var request = goAmoCrm.Request{}
	request.Url = "v2/json/contacts/set"
	request.Post = true
	action := "update"
	if(this.Id == 0) {
		action = "add"
	}
	var params = map[string]map[string]map[string]map[string]string{
		"request": {
			"contacts": {
				action: {
					"id" : string(this.Id),
					"last_modified" : this.LastModified,
					"name" : this.Name,
					"company_name" : this.CompanyName,
					"responsible_user_id" : string(this.ResponsibleUserId),
					"tags" : this.Tags,
					"linked_leads_id" : "",
					"custom_fields" : "",
				},
			},
		},
	}
	str, _ :=json.Marshal(params)
	fmt.Println(string(str))
	resonse := handler.Request(request)
	fmt.Println(resonse)
	//request.Body =

}