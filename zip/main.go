package main
import (
	"log"
	"bytes"
	"compress/gzip"
	"io"
)
func main() {
	s := zipByGzip("5660ad8f4012a5f0867622024550165")
	log.Println(s)
	revert(s)
	s = zipByGzip("5555550000001111112222")
	log.Println(s)
	revert(s)
}

func zipByGzip(str string) string {
	var in bytes.Buffer
	w := gzip.NewWriter(&in)
	w.Write([]byte(str))
	w.Close()
	log.Println("@before:",len([]byte(str))," @after:",len(in.Bytes()))
	return string(in.Bytes())
}

func revert(str string) {
	var out bytes.Buffer
	r, _ := gzip.NewReader(bytes.NewBufferString(str))
	io.Copy(&out, r)
	log.Println(out.String())
}
