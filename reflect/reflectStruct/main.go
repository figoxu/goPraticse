package main
import (
	"fmt"
	"reflect"
)
type User struct {
	Id int
	Name string
	Age int
	Title string
}
func(this User)Test()string{
	return this.Name
}
func main(){
	// interface可以接受任何类型的参数，如果不用接口，就只能用 := 让go自己去判断类型了
	// o := &User{1,"Tom",12,"nan"}
	var o interface {} = &User{1,"Tom",12,"nan"}
	v := reflect.ValueOf(o)
	fmt.Println(v)
	m := v.MethodByName("Test")
	rets := m.Call([]reflect.Value{})
	fmt.Println(rets)
}