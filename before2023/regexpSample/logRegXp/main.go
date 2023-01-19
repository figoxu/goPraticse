package main

import (
	"io/ioutil"
	"github.com/quexer/utee"
	"strings"
	"time"
	"github.com/figoxu/Figo"
	"fmt"
	"log"
)

var local *time.Location

func init(){
	local, _ = time.LoadLocation("Asia/Shanghai")
}

func main(){
	b,err:=ioutil.ReadFile("/Users/xujianhui/Documents/err.log")
	utee.Chk(err)
	contents:=strings.Split(string(b),"--")
	for _,content:=range contents {
		pingTime:=minTime("PING1",content)
		onlineTime:=minTime("ONLINE",content)
		sec:=onlineTime.Unix()-pingTime.Unix()
		fmt.Println("掉线时间：",pingTime,"\t上线时间:",onlineTime,"\t故障总时间:",sec,"秒")
	}
}

func minTime(feature,content string) time.Time{
	defer Figo.Catch()
	parser:=Figo.Parser{
		PrepareReg:[]string{fmt.Sprint(".+",feature),"\\d+/\\d+/\\d+ \\d+:\\d+:\\d+"},
		ProcessReg: []string{},
	}
	contents := parser.Exe(content)
	parseT := func (str string) (time.Time, error){
		defer Figo.Catch()
		return time.ParseInLocation("2006/01/02 15:04:05", str,local)
	}
	t,err:= parseT(contents[0])
	if err!=nil {
		defer Figo.Catch()
		log.Println("@content[0]",contents[0]," @content[1]:",contents[1])
		t,err= parseT(contents[1])
	}
	return t
}
