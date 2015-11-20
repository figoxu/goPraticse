package main
import (
	"fmt"
	glru "github.com/hashicorp/golang-lru"
	"log"
)



type Test struct {
	Name  string
	Tp    string
	Count int
}

func main(){
	fmt.Println("hello")
	m,_ := glru.New(1000)
	m.Add("test",Test{Name:"figo",Tp:"android",Count:1024})
	for k:= range m.Keys(){
		log.Println("@k:",k)
		d,_:=m.Get("test")
		v:=d.(Test)
		log.Println("@name:",v.Name,"@count:",v.Count,"@tp:",v.Tp)
	}
	d,_:=m.Get("test")
	v:=d.(Test)
	log.Println("@name:",v.Name,"@count:",v.Count,"@tp:",v.Tp)
}
