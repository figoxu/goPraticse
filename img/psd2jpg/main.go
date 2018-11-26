package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"os/exec"
	"bytes"
	"path/filepath"
	"os"
	"fmt"
	"strings"
)

func main(){
	dir := "/Users/xujianhui/develop/golang/gopath/src/github.com/figoxu/goPraticse/img/psd2jpg/input"
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if strings.Index(name,".psd")!=-1 && strings.Index(name,".jpg")==-1{
			outfileName:=fmt.Sprint(path,".jpg")
			cmd := fmt.Sprint("convert -layers flatten ",path," ",outfileName)
			system(cmd)
		}
		return nil
	})
}

func system(s string) string {
	defer Figo.Catch()
	cmd := exec.Command("/bin/sh", "-c", s) //调用Command函数
	var out bytes.Buffer                    //缓冲字节
	cmd.Stdout = &out                       //标准输出
	err := cmd.Run()                        //运行指令 ，做判断
	utee.Chk(err)
	return out.String() //输出执行结果
}
