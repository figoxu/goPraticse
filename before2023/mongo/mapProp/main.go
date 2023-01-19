package main

import (
	"github.com/figoxu/utee"
	"github.com/pborman/uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	DB = "figoTest"
	C  = "st"
)

func main() {
	log.Println("hello")
	mg := connect("192.168.56.101/figoTest")
	cp := mg.Copy()
	defer cp.Close()

	apps := map[string]string{}
	apps["com.test.com"] = "特斯特"
	apps["com.cool.me"] = "酷儿"
	apps["com.mi.me"] = "米密"

	m := &MapTest{
		Id:   bson.NewObjectId(),
		Dvid: uuid.NewUUID().String(),
		Ct:   time.Now(),
		Mt:   utee.Tick(),
		Apps: apps,
	}

	cp.DB(DB).C(C).Insert(m)

	//	func (p *Student) incTest(){
	//		if err := p.ds.DB(DB).C(C).UpdateId(p.Id, bson.M{"$inc":bson.M{"visitTimes":1}});err!=nil {
	//			log.Println("@err:",err)
	//		}
	//	}

}

func connect(db_connection string) *mgo.Session {
	session, err := mgo.Dial(db_connection)
	utee.Chk(err)
	return session
}

type MapTest struct {
	Id   bson.ObjectId     `bson:"_id"`
	Dvid string            `bson:"dvid"`
	Ct   time.Time         `bson:"ct"`
	Mt   int64             `bson:"mt"`
	Apps map[string]string `bson:"apps"`
}
