package main
import (
	"fmt"
	"gopkg.in/mgo.v2"
	"github.com/figoxu/utee"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	DB = "figoTest"
	C = "st"

)

func main(){
	fmt.Println("hello")
	mg := connect("192.168.56.101/figoTest")
	cp := mg.Copy()
	defer cp.Close()

	s := &Student{
		Id:bson.NewObjectId(),
		ds:cp,
		name:"figo",
		visitTimes:10,
	}
	err := cp.DB(DB).C(C).Insert(s)
	utee.Chk(err)
	A := func(){
		s.incTest()
	}
	B := func(){
		s.setTest()
	}
	log.Println("inc 10000 times cost:",execute(A,10000))
	log.Println("set 10000 times cost:",execute(B,10000))
}
func connect(db_connection string) *mgo.Session {
	session, err := mgo.Dial(db_connection)
	utee.Chk(err)
	return session
}


func execute(method func(),t int) int64{
	b := utee.Tick()
	for i:=0;i<t;i++{
		method()
	}
	b = utee.Tick()-b
	return b
}


type Student struct {
	ds *mgo.Session
	Id      bson.ObjectId `bson:"_id" `
	name string
	visitTimes int
}

func (p *Student) incTest(){
	if err := p.ds.DB(DB).C(C).UpdateId(p.Id, bson.M{"$inc":bson.M{"visitTimes":1}});err!=nil {
		log.Println("@err:",err)
	}
}

func (p *Student) setTest(){
	if err := p.ds.DB(DB).C(C).UpdateId(p.Id, utee.J{"$set": bson.M{"visitTimes":1}});err!=nil {
		log.Println("@err:",err)
	}
}










