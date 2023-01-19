package main
import (
	"fmt"
	"github.com/figoxu/utee"
	"strconv"
)


func main(){
	fmt.Println("This is a program which mock to gen ios token")
	genCount := 1000;
	seed:="seedStr_bad__"
	f := utee.Md5Str(seed)
	for i:=0;i<genCount;i++ {

		v := fmt.Sprint( f(strconv.Itoa(i+12345)),f(strconv.Itoa(i+5678)))
		fmt.Println(v)
	}
}
