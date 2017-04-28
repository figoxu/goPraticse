package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"fmt"
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/figoxu/utee"
	"log"
)

type Student struct {
	Id   int64
	Name string
	Age  int
	Hobby string `orm:"type(jsonb);null"`

}

type Hobby struct {
	Sport string
	Music string
	Game string
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

	hb := Hobby{
		Sport:"FootBall",
		Music:"We will rock you",
		Game:"LOL",
	}
	b,err:=	json.Marshal(hb)
	utee.Chk(err)
	log.Println(string(b))
	stu.Hobby=string(b)
	fmt.Println(o.Insert(stu))
	beego.Run()
}
