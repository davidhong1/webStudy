package models

type GromPerson struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

//设置表名
func (GromPerson) TableName() string {
	return "grom_person"
}
