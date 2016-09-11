package main

import (
	"log"
	"github.com/Unknwon/goconfig"
	"github.com/figoxu/Figo"
	"github.com/figoxu/utee"
)

func main() {
	confFP := Figo.NewFilePath("./conf.txt")
	log.Println(confFP.FullPath())
	fp,err:=confFP.FullPath()
	cf,err:=goconfig.LoadConfigFile(fp)
	utee.Chk(err)
	cf.SetValue("Figo","Hello","World")
	 confFP.Open()
	w,err:=confFP.Writer()
	utee.Chk(err)
	log.Println(cf.GetSectionList())
	err=goconfig.SaveConfigData(cf,w)
	w.Flush()
	utee.Chk(err)
}
