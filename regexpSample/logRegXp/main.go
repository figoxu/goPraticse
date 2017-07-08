package main

import (
	"io/ioutil"
	"github.com/quexer/utee"
	"strings"
	"time"
	"github.com/figoxu/Figo"
	"fmt"
)

var local *time.Location

func init(){
	local, _ = time.LoadLocation("Asia/Shanghai")
}

func main(){
	b,err:=ioutil.ReadFile("/Users/xujianhui/Documents/err02.log")
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
	parser:=Figo.Parser{
		PrepareReg:[]string{fmt.Sprint(".+",feature),"\\d+/\\d+/\\d+ \\d+:\\d+:\\d+"},
		ProcessReg: []string{},
	}
	t,err:= time.ParseInLocation("2006/01/02 15:04:05",parser.Exe(content)[0],local)
	utee.Chk(err)
	return t
}
