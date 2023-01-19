package main

import (
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/mohae/deepcopy"
)

type Person struct {
	Id     int
	Name   string
	Friend *Person
}

type Student struct {
	Id      int
	Name    string
	Father  *Person
	Mother  *Person
	Teacher *Person
}

func main() {
	lily := &Person{
		Id:   14,
		Name: "Lily.Nichonas",
	}
	lucy := &Person{
		Id:     13,
		Name:   "Lucy.Nichonas",
		Friend: lily,
	}
	//lily.Friend = lucy   //todo 如果存在相互引用，则会是死循环，直到退出
	lilei := &Person{
		Id:     11,
		Name:   "李磊",
		Friend: lucy,
	}
	hanMM := &Person{
		Id:     12,
		Name:   "韩梅梅",
		Friend: lily,
	}

	student := Student{
		Id:      552223,
		Name:    "测试员",
		Father:  lilei,
		Mother:  hanMM,
		Teacher: lucy,
	}

	fmt.Println(Figo.JsonString(student))

	var s2 Student
	v := deepcopy.Copy(student)
	s2 = v.(Student)

	lilei.Name = "李四"
	lucy.Name = "Lucy.Sun"
	fmt.Println(Figo.JsonString(student))

	fmt.Println(Figo.JsonString(s2))

}
