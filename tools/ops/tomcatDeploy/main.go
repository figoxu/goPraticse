package main

import (
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
)

func main() {
//	remoteFileName := "http://www.apache.org/dist/tomcat/tomcat-7/v7.0.72/bin/apache-tomcat-7.0.72.zip"
	localFileName := "./apache-tomcat-7.0.72.zip"
//	Figo.DownLoad(localFileName,remoteFileName,10);

	err := Figo.UnZip(localFileName,"./tomcat7/")
	utee.Chk(err)
}
