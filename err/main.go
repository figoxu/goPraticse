package main

import (
	"errors"
	"github.com/figoxu/goPraticse/err/a"
	"github.com/jinzhu/gorm"
	pkgerror "github.com/pkg/errors"
	"log"
)

func main() {
	log.Println("hello")
	e := errors.New("test")
	log.Println(e.Error())
	log.Println(e.Error() == "test")
	errNotFound := pkgerror.WithStack(gorm.ErrRecordNotFound)
	if pkgerror.Cause(errNotFound) == gorm.ErrRecordNotFound {
		log.Println("SAME")
	} else {
		log.Panic("DIFFERENT")
	}
	a.TMethod()
}
