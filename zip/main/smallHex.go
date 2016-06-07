package main

import (
	"log"
	"regexp"
	"strconv"
	"unsafe"
)

type SH struct {
	v string
}

func main() {
	v := "12139a532f3c29d678d879135532843838"
	log.Println(v)
	log.Println([]byte(v))
	v2 := hex2bytes(v)
	log.Println(v2)
	log.Println([]byte(v2))

	log.Println(revert([]byte(v2)))


	log.Println("@sizeOf v:",unsafe.Sizeof(v))
	log.Println("@sizeOf v2:",unsafe.Sizeof(&v2))
	v3:=`// 定義一個 Student 的結構
type Student struct {
    name    string
    id      int
    grade   float32
    friends []string
}
// 宣告一個內容為空的 student 變數代表這個結構
student := Student{}`
	log.Println("@sizeOf v3:",unsafe.Sizeof(&v3))
	log.Println("@cap v:",cap([]byte(v)))
	log.Println("@cap v2:",cap([]byte(v2)))
	log.Println("@cap v3:",cap([]byte(v3)))
	s1:= &SH{v:v}
	s2:= &SH{v:v2}
	s3:= &SH{v:v3}
	log.Println("szo  s1:",unsafe.Sizeof(s1.v))
	log.Println("szo  s2:",unsafe.Sizeof(s2.v))
	log.Println("szo  s3:",unsafe.Sizeof(s3.v))
	log.Println("szo  s1:",unsafe.Offsetof(s1.v))
	log.Println("szo  s2:",unsafe.Offsetof(s2.v))
	log.Println("szo  s3:",unsafe.Offsetof(s3.v))
}

func hex2bytes(hex string) string {

	s := []byte{}
	if m, _ := regexp.MatchString("^[0-9a-f]+$", hex); !m {
		log.Println("not hex")
		return hex
	}

	for i := 0; i < len(hex); i = i + 2 {
		var a int64
		if i+1 >= len(hex) {
			a, _ = strconv.ParseInt(string(hex[i]), 16, 10)
		} else {
			a, _ = strconv.ParseInt(string(hex[i:i+2]), 16, 10)
		}
		v := byte(a)
		s = append(s, v)
	}
	return string(s)

}

func revert(arr []byte) string{
	var str string
	for i := 0; i < len(arr); i++ {
		str += strconv.FormatInt(int64(arr[i]), 16)
	}
	return str
}
