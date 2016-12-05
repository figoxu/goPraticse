package main

import "fmt"
import "reflect"
import "encoding/xml"

type ReflectMethodSample struct{
}

func (this *ReflectMethodSample)Echo(){
	fmt.Println("echo()")
}

func (this *ReflectMethodSample)Echo2(){
	fmt.Println("echo--------------------()")
}

var xmlstr string=`<root>
    <func>Echo</func>  
    <func>Echo2</func>  
    </root>`

type FuncList struct{
	Names []string `xml:"func"`
}

func main() {
	funcList := FuncList{}
	xml.Unmarshal([]byte(xmlstr), &funcList)

	sample := &ReflectMethodSample{}
	reflectValue := reflect.ValueOf(sample)
	for _,name := range funcList.Names {
		reflectValue.MethodByName(name).Call(nil)
	}
}