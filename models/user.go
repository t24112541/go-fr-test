package models

type CustomerEntity struct {
	Id   int    `json:"id" form:"id" sql:"auto_increment"`
	Name string `json:"name" form:"name" sql:"not_null"`
}

func (u *CustomerEntity) RestPath() string {
	return "customer"
}

func (u *CustomerEntity) TableName() string {
	return "customers"
}
