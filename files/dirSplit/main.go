package main
import (
	"fmt"
	"encoding/hex"
	"github.com/figoxu/utee"
)

func main(){
	s := "test9324u5uhowareyouwahtever"
	fmt.Println(getOkayDir(s))
}

func getOkayDir(s string) string{
	v := hex.EncodeToString(utee.Md5([]byte(s)))
	v = v[8:24]  //16位md5
	d := fmt.Sprint(v[0:2],"/",v[2:4],"/",v[4:6],"/",v[6:8],"/",v[8:10],"/",v[10:12],"/",v[12:14])
	return d
}
