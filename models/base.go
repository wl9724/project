package models

import "github.com/beego/beego/v2/client/orm"

type User struct {
	Id   int
	Name string
}

func init() {
	orm.RegisterModel(new(User))
}
