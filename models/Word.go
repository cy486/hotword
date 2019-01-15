package models

import (
	"fmt"
	orm2 "github.com/astaxie/beego/orm"
	_ "github.com/go-mysql/mysql"
)

type HotWord struct {
	Id      int `orm:"pk;auto;"`
	Word    string
	UrL     orm2.TextField
	Content orm2.TextField
	IsA     string
}
type WordContent struct {
	Id      int
	Word    int `orm:"pk"`
	Title   string
	Concent orm2.TextField
	Url     string
}

func init() {
	err := orm2.RegisterDataBase("default", "mysql", "root:123456@tcp(localhost:3306)/text?charset=UTF8")
	orm2.RegisterModel(new(HotWord), new(WordContent))
	err = orm2.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
}
