package main
import (
	"log"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"errors"
)

func main(){
	defer log.Println("defer 001")
	defer log.Println("defer 002")
	defer log.Println("defer 003")
	defer Figo.Catch()
	log.Println("execute 001")
	utee.Chk(errors.New("Error"))
	log.Println("execute 002")
}
//2016/11/09 10:28:58 execute 001
//2016/11/09 10:28:58 Error  (recover)
//2016/11/09 10:28:58 defer 003
//2016/11/09 10:28:58 defer 002
//2016/11/09 10:28:58 defer 001