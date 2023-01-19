package main

import (
	"fmt"
	"github.com/figoxu/utee"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"github.com/pborman/uuid"
)

const (
	DB = "figoTest"
	C  = "st1"
	C2 = "st2"
)

var ds *mgo.Session
func main() {
	fmt.Println("hello")
	mg := connect("192.168.56.101/figoTest")
	cp := mg.Copy()
	defer cp.Close()
	ds=mg
	err := cp.DB(DB).C(C2).EnsureIndex(mgo.Index{Key: []string{"name"}, Unique: true})
	utee.Chk(err)
	A := func() {
		insertATest()
	}
	B := func() {
		insertBTest()
	}
	log.Println("insert with    index 10000 times cost:", execute(B, 10000))
	log.Println("insert without index 10000 times cost:", execute(A, 10000))
}
func connect(db_connection string) *mgo.Session {
	session, err := mgo.Dial(db_connection)
	utee.Chk(err)
	return session
}

func execute(method func(), t int) int64 {
	b := utee.Tick()
	for i := 0; i < t; i++ {
		method()
	}
	b = utee.Tick() - b
	return b
}

type Student struct {
	Id         bson.ObjectId `bson:"_id" `
	Name       string		`bson:"name"`
	VisitTimes int		`bson:"vt"`
}

func insertATest() {
	n := fmt.Sprint("figo", uuid.NewUUID().String())
	s := &Student{
		Id:         bson.NewObjectId(),
		Name:       n,
		VisitTimes: 10,
	}
	err := ds.DB(DB).C(C).Insert(s)
	utee.Chk(err)
}

func  insertBTest() {
	n := fmt.Sprint("figo", uuid.NewUUID().String())
	s := &Student{
		Id:         bson.NewObjectId(),
		Name:       n,
		VisitTimes: 10,
	}
	err := ds.DB(DB).C(C2).Insert(s)
	utee.Chk(err)
}
