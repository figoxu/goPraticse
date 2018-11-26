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
	"time"
	"log"
)

func main(){
	log.Println("本程序由Figo开发")
	log.Println("依赖外部库：ftp://ftp.imagemagick.org/pub/ImageMagick/binaries ")
	log.Println("请执行确定下载")
	for {
		log.Println("Ctrl+C 或 关闭命令窗口 即可")
		log.Println("请输入需要处理的目录:")
		var dir string
		fmt.Scanln(&dir)
		log.Println("准备处理目录:",dir)
		time.Sleep(time.Second*time.Duration(2))
		processImg(dir)
		log.Println("本批次处理结束")
	}
	fmt.Println("欢迎下次使用")
}

func processImg(dir string){
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if strings.Index(name,".psd")!=-1 && strings.Index(name,".jpg")==-1{
			outfileName:=fmt.Sprint(path,".jpg")
			cmd := fmt.Sprint("convert -layers flatten ",path," ",outfileName)
			log.Println(cmd)
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
