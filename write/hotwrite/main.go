package main

import (
	"github.com/astaxie/beego/orm"
	"github.com/figoxu/Figo"
	"log"
)

type HotColdInfo struct {
	Id          int
	IdentityNo  string
	Name        string
	Age         int
	Salary      int
	Description string
}

var rsWriter RedisqlWriter

func main() {
	log.Println("main")
	rp := Figo.RedisPool("127.0.0.1", "")
	orm.Debug = true
	conf := Figo.MysqlConf{
		Host:       "127.0.0.1",
		Pwd:        "xujianhui0915",
		User:       "root",
		Port:       3306,
		Name:       "figo_research",
		ConnIdle:   2,
		ConnActive: 2,
	}
	conf.Conf(new(HotColdInfo))
	orm.RunSyncdb("default", false, true)

	rsWriter = RedisqlWriter{
		hotWriter:  rp,
		coldWriter: orm.NewOrm(),
	}

}
