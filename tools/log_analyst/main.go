package main

import (
	"log"
	"github.com/figoxu/Figo"
	"io/ioutil"
	"github.com/quexer/utee"
	"fmt"
)

func main(){
	logPath := "/home/figo/delete_it/log.20170330.err"
	packageStr := "talkingdata"
	logAnalysist(logPath,packageStr)
}


func logAnalysist(logPath,packageStr string){

	type Item struct {
		key string
		val int32
	}

	sortPrint:=func (items []Item){
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

	b,err:=ioutil.ReadFile(logPath)
	utee.Chk(err)
	parser := Figo.Parser{
		PrepareReg:[]string{fmt.Sprint(packageStr,".+"),"\\(.+java.+[0-9]+\\)"},
		ProcessReg:[]string{},
	}
	contents := parser.Exe(string(b))
	m := map[string]int32{}
	for _,content:= range contents {
		m[content]=m[content]+1
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