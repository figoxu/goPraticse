package main

import (
//	"github.com/jinzhu/copier"
//	"github.com/mitchellh/copystructure"

	"github.com/ulule/deepcopier"

//	"github.com/mohae/deepcopy"

//	deepcopy "github.com/ld9999999999/go-interfacetools"
//	"log"
	"math/rand"
	"time"
	"fmt"
//	"github.com/figoxu/utee"
	"github.com/figoxu/utee"
)

type Obj struct {
	Name string `json:"Name"`
	Score int `json:"score"`
//	childs []Obj
	Member *Obj `json:"member"`
//	member Obj `deepcopier:"field:member"`
}


func (p *Obj) display() {
	fmt.Println("@name:", p.Name, " @score:", p.Score, "@member:", p.Member)
//	fmt.Println("@name:", p.Name, " @score:", p.Score)
//	for _, obj := range p.childs {
//		log.Println(obj)
//	}
}

func randomeObj()Obj{

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		rv := func(v string) string {
			return fmt.Sprint(v, "-", r.Int())
		}
	return Obj{
		Name: rv("name"),
		Score:r.Int(),
		Member : &Obj{
			Name:rv("member"),
			Score:r.Int(),
		},
	}
//	return Obj{
//		name: rv("name"),
//		score:r.Int(),
//	}
}


// Prop

func main(){
	obj :=randomeObj()
	var obj2 Obj
	err := deepcopier.Copy(&obj).To(&obj2)
//	err := deepcopy.CopyOut(&obj,&obj2)
	utee.Chk(err)
//	obj2:=deepcopy.Copy(&obj)

//	log.Println(tmp)
//	copier.Copy(&obj2,&obj)
//	copier.Copy(&obj,&obj2)
//obj2 := tmp.(*Obj)
	obj2.Member.Score=100
	obj.display()
	obj2.display()

}