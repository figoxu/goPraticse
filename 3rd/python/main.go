package main

import (
	"github.com/sbinet/go-python"
	"fmt"
)

var PyStr = python.PyString_FromString
var GoStr = python.PyString_AS_STRING


func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	hello := ImportModule("/Users/xujianhui/develop/golang/gopath/src/github.com/figoxu/goPraticse/3rd/python", "main")
	fmt.Printf("[MODULE] repr(hello) = %s\n", GoStr(hello.Repr()))


	b := hello.GetAttrString("test")

	bArgs := python.PyTuple_New(1)
	python.PyTuple_SetItem(bArgs, 0, PyStr("xixi"))

	res := b.Call(bArgs, python.Py_None)
	fmt.Printf("[CALL] b('xixi') = %s\n", GoStr(res))

	a := hello.GetAttrString("l_val_t")
	fmt.Printf("[VARS] a = %#v\n", python.PyInt_AsLong(a))
}


// InsertBeforeSysPath will add given dir to python import path
//func InsertBeforeSysPath(p string) string {
//	sysModule := python.PyImport_ImportModule("sys")
//	path := sysModule.GetAttrString("path")
//	python.PyList_Insert(path, 0, PyStr(p))
//	return GoStr(path.Repr())
//}

// ImportModule will import python module from given directory
func ImportModule(dir, name string) *python.PyObject {
	sysModule := python.PyImport_ImportModule("sys") // import sys
	path := sysModule.GetAttrString("path")          // path = sys.path
	python.PyList_Insert(path, 0, PyStr(dir))        // path.insert(0, dir)
	return python.PyImport_ImportModule(name)        // return __import__(name)
}