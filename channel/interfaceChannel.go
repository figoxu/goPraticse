package main
import (
	"log"
	"fmt"
	"time"
)


type OpAs interface {
	GetKey() string
	Op(ds string)
}
type Op struct {
	key string
	before   func(string, string)
	after   func(string) int
}

func (p *Op) GetKey() string {
	return p.key;
}

func (p *Op) Op(ds string)  {
	if(p.before!=nil){
		p.before(ds,"test")
	}
	fmt.Println("working @val:",ds,"   @key:",p.GetKey());
	if(p.after!=nil){
		p.after(ds)
	}
}


const (
	MAX_OPQ_AS_SIZE = 200000
)

var (
	g_as_op_q = make(chan OpAs, MAX_OPQ_AS_SIZE)
)

func asOpConsume(ds string) {
	for item := range g_as_op_q {
		item.Op(ds)
	}
}

func asOpEnque(op OpAs) {
	select {
	case g_as_op_q <- op:
	default:
		log.Println("[warn] op as basic overstock")
	}
}


func main(){
	go asOpConsume("SYS_TEST_MSG")
	before :=func(v1,v2 string){
		fmt.Println("====before===")
		fmt.Println(v1,v2);
	}
	after :=func(v1 string)int{
		fmt.Println("====after===")
		fmt.Println(v1);
		return 200
	}
	op :=&Op{
		key:"test_key",
		before:before,
		after:after,
	}
	asOpEnque(op)


	time.Sleep(time.Second*5)
}
