package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"fmt"
	"github.com/astaxie/beego"
)

type Student struct {
	Id   int64
	Name string
	Age  int
}

//Jsonb string `orm:"type(jsonb);null"`
func init() {
	orm.RegisterModel(new(Student))
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user=xxx password=xxx dbname=sdz_dev host=47.93.118.90 port=3432 sslmode=disable")
	orm.RunSyncdb("default", false, true)
}

func main(){
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default")
	stu := new(Student)
	stu.Name = "tom"
	stu.Age = 25
	fmt.Println(o.Insert(stu))
	beego.Run()
}
