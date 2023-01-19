package main

import (
	"log"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/quexer/utee"
)

const (
	LUA        as.Language = "LUA"
	UDF_FILTER             = ``
)

func initLua(ac *as.Client) {
	regTask, err := ac.RegisterUDF(nil, []byte(UDF_FILTER), "udfFigoMapReduce.lua", LUA)
	utee.Chk(err)

	err = <-regTask.OnComplete()
	utee.Chk(err)
}

func main() {
	log.Println("Hello")
}
