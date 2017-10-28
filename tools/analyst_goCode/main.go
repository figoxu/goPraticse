package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
)

func main() {
	goFileName:="/Users/xujianhui/develop/golang/gopath/src/github.com/figoxu/goPraticse/tdTalker/lecture01/channel/pAndC/main.go"

	fu:=Figo.FileUtee{}
	fs,err:=fu.ReadLinesSlice(goFileName)
	utee.Chk(err)
	var codeStack CodeStack = make([]Code,0)
	for _,f:=range fs {
		codeStack = AppendCode(codeStack,f)
	}
	rootNode:=ParseStackToNode(codeStack,"{","}")
	Figo.PrintJson("===>",rootNode)
}
