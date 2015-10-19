package main
import (
	"github.com/alicebob/qr"
	"fmt"
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
	go func() {
		for e := range q.Dequeue() {
			fmt.Printf("We got: %v\n", e)
		}
	}()

	// elsewhere:
	q.Enqueue("aap")
	q.Enqueue("noot")
}