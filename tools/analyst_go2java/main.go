package main

import (
	"os"
	"io"
	"github.com/figoxu/utee"
	"log"
	"github.com/figoxu/Figo"
)

func main(){
	fpath:="D:/figo/workspace/workspace_praticse/workspace_go/sdz-mobile-app/handler-CampeSite.go"
	content:=read(fpath)
	parser:=Figo.Parser{
		PrepareReg :[]string{"type.+?struct.+?\\{[\\s\\S]+?\\}"},
		ProcessReg :[]string{},
	}
	tpStruts:=parser.Exe(content)
	log.Println(tpStruts)
	log.Println(len(tpStruts))
}

func read(path string)string{
	fi,err := os.Open(path)
	utee.Chk(err)
	defer fi.Close()
	chunks := make([]byte,1024,1024)
	buf := make([]byte,1024)
	for{
		n,err := fi.Read(buf)
		if err != nil && err != io.EOF{panic(err)}
		if 0 ==n {break}
		chunks=append(chunks,buf[:n]...)
	}
	return string(chunks)
}
