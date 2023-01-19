package main
import (
	"github.com/quexer/utee"
	"strconv"
	"log"
	"encoding/json"
	"fmt"
	"unsafe"
)

type Result struct {
	inter	 *Result
	_type *Result
	data  unsafe.Pointer
	Item  string `json:"item"`
	Value int   `json:"value"`
}

func (p *Result) getSortVal() int{
	return utee.Fint(p.Item)
}

func main(){
	m := make(map[string]int)

	m["67"] = m["67"] + 10
	m["hello"] = 1024
	fmt.Println("-------------------")
	fmt.Println(m)
	fmt.Println("------delete key hello-------------")
	delete(m,"hello")
	fmt.Println(m)
	fmt.Println("-------------------")

	m1 := make(map[string]bool)
	log.Println("@m1['not exist key'] : ",m1["not exist key"])

	start := utee.Fint("60")
	end :=utee.Fint("70")
	for i:=start;i<=end;i++{
		if( m[strconv.Itoa(i)] <=0 ){
			m[strconv.Itoa(i)] = 0;
		}
	}

	a := make([]Result, 0, len(m))
	for k, v := range m {
		fmt.Println("@key:",k)
		a = append(a, Result{
			Item:  k,
			Value: v,
		})
	}


	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j].getSortVal() > a[j+1].getSortVal() {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	for index,v := range a{
		log.Println("@index:",index,"  @v:",v.Item)
	}

//	if tmp,ok := a.([]interface{});ok{
//
//		vaa :=userSort(tmp)
//		if valA ,ok := vaa.([]Result);ok{
//			a=valA
//		}
//	}

	val, _ := json.Marshal(a)
	log.Println("@val:",string(val))
}

type SortObj interface {
	getSortVal()int
}
//todo 只是想抽象一个排序而已
//func userSort(array []interface{}) []interface{} {
//	for i := 0; i < len(array); i++ {
//		for j := 0; j < len(array)-i-1; j++ {
//			if array[j].(SortObj).getSortVal() < array[j+1].(SortObj).getSortVal() {
//				array[j], array[j+1] = array[j+1], array[j]
//			}
//		}
//	}
//	return array
//}

