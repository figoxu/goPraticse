package main


import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	CompressZip()   //压缩
	DeCompressZip() //解压缩
}

func CompressZip() {
	const dir = "/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/files/"
	//获取源文件列表
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	fzip, _ := os.Create("fileCode.zip")
	w := zip.NewWriter(fzip)
	defer w.Close()
	for _, file := range f {
		fw, _ := w.Create(file.Name())
		filecontent, err := ioutil.ReadFile(dir + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		n, err := fw.Write(filecontent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}

func DeCompressZip() {
	const File = "/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/fileCode.zip"
	const dir = "/home/figo/delete/tmp/"
	os.Mkdir(dir, 0777) //创建一个目录

	cf, err := zip.OpenReader(File) //读取zip文件
	if err != nil {
		fmt.Println(err)
	}
	defer cf.Close()
	for _, file := range cf.File {
		rc, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}

		f, err := os.Create(dir + file.Name)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		n, err := io.Copy(f, rc)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}

}