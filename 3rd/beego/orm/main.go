package main

import (
	"log"
	"github.com/astaxie/beego/orm"

	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"github.com/astaxie/beego"
)

type User struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}


type Student struct {
	Id    int // 主键
	Name  string
	Age   int
	Sex   string
	Score float32
	Addr  string
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Student))
	orm.RegisterModel(new(User), new(Profile),new(Tag),new(Post))

	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")
	orm.RunSyncdb("default", false, true)
}

func main() {
	log.Println("hello")
	o := orm.NewOrm()
	o.Using("default")
	stu := new(Student)

	stu.Name = "bei"
	stu.Age = 25
	stu.Sex = "m"
	stu.Score = 88
	stu.Addr = "hunan.leiyang"

	fmt.Println(o.Insert(stu))
	beego.Run()
}
