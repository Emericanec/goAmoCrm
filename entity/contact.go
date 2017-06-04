package entity

type CustomFiledValue struct {
	value   string
	enum    string
	subtype string
}

type CustomFiled struct {
	id     int
	values []CustomFiledValue
}

type ContactListParam struct {
	LimitRows         int
	LimitOffset       int
	Id                int
	Query             string
	ResponsibleUserId string
	Type              string
}

type Contact struct {
	Id                int `json:"id"`
	Name              string `json:"name"`
	RequestId         string `json:"request_id"`
	LastModified      int `json:"last_modified"`
	AccountId         int `json:"account_id"`
	DateCreate        int `json:"date_create"`
	CreatedUserId     int `json:"created_user_id"`
	ModifiedUserId    int `json:"modified_user_id"`
	ResponsibleUserId int `json:"responsible_user_id"`
	GroupId           int `json:"group_id"`
	ClosestTask       int `json:"closest_task"`
	LinkedCompanyId   string `json:"linked_company_id"`
	CompanyName       string `json:"company_name"`
	Tags              []Tag `json:"tags"`
	LinkedLeadsId     interface{} `json:"linked_leads_id"`
	CustomFields      []CustomFiled `json:"custom_fields"`
}

type ContactListResponse struct {
	Contacts   []Contact `json:"contacts"`
	ServerTime int `json:"server_time"`
}

type ContactListResponseRoot struct {
	Response ContactListResponse `json:"response"`
}

type ContactsSetRequestRoot struct {
	Request ContactsSetRequestContacts `json:"request"`
}

type ContactsSetRequestContacts struct {
	Contacts ContactsSetRequest `json:"contacts"`
}

type ContactsSetRequest struct {
	Add    []Contact `json:"add"`
	Update []Contact `json:"update"`
}

type ContactsSetResponseRoot struct {
	Response ContactsSetResponseContacts `json:"response"`
}

type ContactsSetResponseContacts struct {
	Contacts   ContactsSetResponse `json:"contacts"`
	ServerTime int `json:"server_time"`
}

type ContactsSetResponse struct {
	Add    []ContactsSetResponseAdd `json:"add"`
	Update []ContactsSetResponseUpdate `json:"update"`
}

type ContactsSetResponseAdd struct {
	Id        int `json:"id"`
	RequestId int `json:"request_id"`
}

type ContactsSetResponseUpdate struct {
	Id        int `json:"id"`
	RequestId int `json:"request_id"`
}
