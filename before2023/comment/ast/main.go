package main

import (
	"fmt"
	"github.com/figoxu/utee"
	"go/parser"
	"go/token"
	"log"
	"path"
	"reflect"
	"strings"
)

func main() {

	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	log.Println("hello")

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, path.Join("/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/comment/ast", "commentDoc", "doc.go"), nil, parser.ParseComments)
	utee.Chk(err)
	if f.Comments != nil {
		for _, c := range f.Comments {
			for _, s := range strings.Split(c.Text(), "\n") {
				log.Println("@s:", s)
			}
		}
	}
}
