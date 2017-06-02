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
	LimitRows int
	LimitOffset int
	Id int
	Query string
	ResponsibleUserId string
	Type string
}

type Contact struct {
	Id                int
	LastModified      string
	Name              string
	CompanyName       string
	ResponsibleUserId int
	Tags              string
	LinkedLeadsId     []int
	CustomFields      []CustomFiled
}
