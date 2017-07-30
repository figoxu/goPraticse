package main

import (
	"time"
	"io/ioutil"
	"strings"
	"github.com/quexer/utee"
	"github.com/figoxu/Figo"
	"fmt"
	"log"
)

var local *time.Location

func init(){
	local, _ = time.LoadLocation("Asia/Shanghai")
}

//cat nohup.out|grep OfflineTracker|grep -B1 -A1 By > err.log
func main(){

	b,err:=ioutil.ReadFile("/Users/xujianhui/Documents/err.log")
	utee.Chk(err)
	contents:=strings.Split(string(b),"--")
	for _,content:=range contents {
		pingTimeB4Failure:=minTime("HeartBeat",content)
		pingTimeAfterFailure:=maxTime("HeartBeat",content)
		sec:=pingTimeAfterFailure.Unix()-pingTimeB4Failure.Unix()
		fmt.Println("超时前Ping时间：",pingTimeB4Failure,"\t超时后Ping时间:",pingTimeAfterFailure,"\t间隔总时间:",sec,"秒")

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


func maxTime(feature,content string) time.Time{
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
	t,err:= parseT(contents[len(contents)-1])
	if err!=nil {
		defer Figo.Catch()
		log.Println("@content[0]",contents[0]," @content[1]:",contents[1])
		t,err= parseT(contents[1])
	}
	return t
}