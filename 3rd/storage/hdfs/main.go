package main

import (
	"github.com/colinmarc/hdfs"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"log"
	"os"
)

func main() {
	defer Figo.Catch()
	os.Setenv("HADOOP_USER_NAME", "root")
	client, _ := hdfs.New("192.168.108.131:9000")
	fname := "/hello4.txt"
	if _, err := client.Stat(fname); err != nil {
		err := client.CreateEmptyFile(fname)
		utee.Chk(err)
	}
	w, err := client.Append(fname)
	utee.Chk(err)
	w.Write([]byte("Hello Figo Over There"))
	w.Close()
	b, err := client.ReadFile(fname)
	utee.Chk(err)
	log.Println("===[Log I Read Is]===")
	log.Println(string(b))
	log.Println("===[Over]===")
}
