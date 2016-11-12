package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := list.New()
	a.PushBack("大鸟")
	a.PushBack("小菜")
	a.PushBack("行李")
	a.PushBack("老外")
	a.PushBack("公交内部员工")
	a.PushBack("小偷")
	for i := a.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
