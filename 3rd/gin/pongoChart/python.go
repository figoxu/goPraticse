package main

import (
	"github.com/sbinet/go-python"
	"strings"
)

var (
	metaModule *python.PyObject
	PyStr      = python.PyString_FromString
	PyInt      = python.PyInt_FromLong
	GoStr      = python.PyString_AsString
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
	metaModule = ImportModule("./", "meta")
}

func ImportModule(dir, name string) *python.PyObject {
	sysModule := python.PyImport_ImportModule("sys") // import sys
	path := sysModule.GetAttrString("path")          // path = sys.path
	python.PyList_Insert(path, 0, PyStr(dir))        // path.insert(0, dir)
	return python.PyImport_ImportModule(name)        // return __import__(name)
}

func getTableNames(drivername, database, username, password, host, port string) []string {
	getTableNamesFunc := metaModule.GetAttrString("get_table_names")
	bArgs := python.PyTuple_New(6)
	python.PyTuple_SetItem(bArgs, 0, PyStr(drivername))
	python.PyTuple_SetItem(bArgs, 1, PyStr(database))
	python.PyTuple_SetItem(bArgs, 2, PyStr(username))
	python.PyTuple_SetItem(bArgs, 3, PyStr(password))
	python.PyTuple_SetItem(bArgs, 4, PyStr(host))
	python.PyTuple_SetItem(bArgs, 5, PyStr(port))
	rsp := getTableNamesFunc.Call(bArgs, python.Py_None)
	resp := GoStr(rsp)
	return strings.Split(resp, ",")
}
