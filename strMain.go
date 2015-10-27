package main
import (
	"log"
	"bytes"
	"github.com/quexer/utee"
	"strconv"
)

func main(){
	t:=utee.Tick()
	log.Println("@t:",t)
	buf := bytes.NewBufferString("hello")
	buf.WriteString(" world @time:"+strconv.FormatInt(t,10))
	log.Println(buf.String())
}