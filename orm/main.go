package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/test?charset=utf8&loc=Asia%2FShanghai")

	orm.RegisterModel(new(User))
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := User{}
	o.Raw("select id, name from user where "+
		"name = ?", "kate").QueryRow(&user)

	fmt.Println(user)

	// user := new(User)
	// user.Id = 1
	// user.Age = 20
	// user.Name = "slene"

	// user2 := new(User)
	// user2.Id = 2
	// user2.Age = 21
	// user2.Name = "kate"

	// fmt.Println(o.Insert(user))
	// fmt.Println(o.Insert(user2))

}
