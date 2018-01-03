package main


import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case A:
		fmt.Println("type struct A")
	case B:
		fmt.Println("type struct B")
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type A struct {
	content string
}

type B struct {
	value int
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(A{"hello"})
	do(B{100})
}