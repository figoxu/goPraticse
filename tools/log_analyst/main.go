package main

import (
	"log"
	"github.com/figoxu/Figo"
	"io/ioutil"
	"github.com/quexer/utee"
)

func main(){

	logPath := "/home/figo/delete_it/log.20170330.err"
	b,err:=ioutil.ReadFile(logPath)
	utee.Chk(err)
	parser := Figo.Parser{
		PrepareReg:[]string{"talkingdata.+","\\(.+java.+[0-9]+\\)"},
		ProcessReg:[]string{},
	}
	log.Println("hello")
	contents := parser.Exe(string(b))

	for _,content:= range contents {
		inc(content)
	}

	items := []Item{}
	for k,v := range m {
		items = append(items,Item{
			key:k,
			val:v,
		})
	}
	sortPrint(items)
}

var (
	m = map[string]int32{}
)

func inc(key string){
	m[key]=m[key]+1
}

type Item struct {
	key string
	val int32
}

func sortPrint(items []Item){
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j].val < items[j+1].val {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
	for _,item := range items {
		log.Println("代码:",item.key,"  出错次数:",item.val)
	}
}