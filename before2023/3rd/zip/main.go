package main

import (
	"archive/zip"
	"fmt"
	"github.com/figoxu/utee"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	CompressZip()   //压缩
	DeCompressZip() //解压缩
}

func CompressZip() {
	const dir = "/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/3rd/zip/"
	//获取源文件列表
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fzip, _ := os.Create("3rd.zip")
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		content, err := ioutil.ReadFile(dir + file.Name())
		utee.Chk(err)
		n, err := fw.Write(content)
		utee.Chk(err)
		log.Println(n, " bytes write")
	}
}

func DeCompressZip() {
	const File = "3rd.zip"
	const dir = "tmp/"
	os.Mkdir(dir, 0777) //创建一个目录

	cf, err := zip.OpenReader(File) //读取zip文件
	utee.Chk(err)
	defer cf.Close()
	for _, file := range cf.File {
		rc, err := file.Open()
		utee.Chk(err)
		f, err := os.Create(dir + file.Name)
		utee.Chk(err)
		defer f.Close()
		n, err := io.Copy(f, rc)
		utee.Chk(err)
		log.Println(n, " bytes read")
	}

}
