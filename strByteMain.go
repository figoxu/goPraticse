package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	st := "hello" + "world"
	log.Println(st)
	v := strconv.Itoa(1024768)
	fmt.Println(v)
	v = fmt.Sprint("hellohowareyou", ",1")
	d := []byte(v)

	v2 := string(d)

	fmt.Println("==================[1]===============")
	fmt.Println(v)
	fmt.Println("==================[2]===============")
	fmt.Println(d)
	fmt.Println("==================[3]===============")
	fmt.Println(v2)

	s := strings.Split("abc,abc", ",")
	fmt.Println(s, len(s))

	v_int, _ := strconv.Atoi("1024")
	fmt.Println(v_int)

}
