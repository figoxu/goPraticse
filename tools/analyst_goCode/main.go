package main

import (
	"github.com/figoxu/Figo"
	//"log"
	"github.com/quexer/utee"
)

func main() {
	//goFileName := Figo.ReadInput("请输入需要解析的Golang文件名:",Figo.THEME_Red,Figo.THEME_Green)
	//log.Println(goFileName)

	goFileName:="/Users/xujianhui/develop/golang/gopath/src/github.com/figoxu/goPraticse/tdTalker/lecture01/channel/pAndC/main.go"

	fu:=Figo.FileUtee{}
	fs,err:=fu.ReadLinesSlice(goFileName)
	utee.Chk(err)
	var codeStack CodeStack = make([]Code,0)
	for _,f:=range fs {
		codeStack = AppendCode(codeStack,f)
	}
	Figo.PrintJson("===>",codeStack)
}
