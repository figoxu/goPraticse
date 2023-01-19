package main
import (
	"github.com/alicebob/qr"
	"time"
)


func main(){
	q, err := qr.New(
	"/tmp/",
	"example",
	qr.OptionBuffer(100),
	qr.OptionTest("your datatype"),
)
	if err != nil {
		panic(err)
	}
	defer q.Close()
//	go func() {
//		for e := range q.Dequeue() {
//			fmt.Printf("We got: %v\n", e)
//		}
//	}()

	time.Sleep(time.Duration(time.Minute*2))
}