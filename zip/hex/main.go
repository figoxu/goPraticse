package main
import (
//	"regexp"
	"log"
//	"strconv"
	_ "net/http/pprof"
	_ "expvar"
	"net/http"
	"os"
	"bufio"
	"regexp"
	"strconv"
)
var data []string

func main(){
	data= getDvIds("/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/data/appleToken/nohup.out")

	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(":6666", nil))
}


func handler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
//	hex1 := hex2bytes(id)
	log.Println(id, "exists", data)
}


func getDvIds(idFile string) []string {
	ids := []string{}
	f, err := os.Open(idFile)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
//		ids = append(ids, scanner.Text())

		id := scanner.Text();
		log.Println(".............")
		log.Println("--[01]--- @id:",id)
		log.Println("--[02]--- @id:",[]byte(id) )
		log.Println("--[03]--- @id:",hex2bytes(id) )
		log.Println("--[04]--- @id:",string(hex2bytes(id)) )
		log.Println(".............")
	}
	if err != nil {
		log.Println("err: ", err)
	}
	return ids
}

func hex2bytes(hex string) []byte {

	var s []byte
	if m, _ := regexp.MatchString("^[0-9a-f]+$", hex); !m {
		log.Println("not hex")
		return []byte(hex)
	}

	for i := 0; i < len(hex); i = i + 2 {
		var a int64
		if i+1 >= len(hex) {
			a, _ = strconv.ParseInt(string(hex[i]), 16, 10)
		} else {
			a, _ = strconv.ParseInt(string(hex[i:i+2]), 16, 10)
		}
		v := byte(a-128)
		//		log.Println("@v:",v+128 ," @a:",a)
		s = append(s, v)
	}
	//	log.Println("befor convert", hex)
	//	log.Println("convert array", s)
	return s

}
