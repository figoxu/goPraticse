package main

import (
	"bytes"
	"github.com/figoxu/utee"
	"log"
	"os/exec" //这个包是主要用来调用cmd命令
	"github.com/codeskyblue/go-sh"
)

func main() {
	log.Println(system("who "))
	log.Println(system("ps -ef "))
	err:=sh.Command("ping","127.0.0.1").Run()
	utee.Chk(err)
}

//调用系统指令的方法，参数s 就是调用的shell命令
func system(s string) string {
	cmd := exec.Command("/bin/sh", "-c", s) //调用Command函数
	var out bytes.Buffer                    //缓冲字节
	cmd.Stdout = &out                       //标准输出
	err := cmd.Run()                        //运行指令 ，做判断
	utee.Chk(err)
	return out.String() //输出执行结果
}

