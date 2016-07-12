package main
import (
	"log"
//	"time"
//	"encoding/base64"
	"strings"
)

func main(){
	url1 := "localhost:3000/api/q/a/739ec0805cebd9b7"
	v := url1[ strings.Index(url1,"/q/a/")+len("/q/a/"):len(url1)]
	log.Println("@v:",v)



}
