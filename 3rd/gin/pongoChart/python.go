package main

import (
	"github.com/sbinet/go-python"
	"strings"
	"regexp"
	"github.com/segmentio/objconv/json"
	"github.com/quexer/utee"
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
	packParam := func(drivername string, database string, username string, password string, host string, port string) *python.PyObject {
		bArgs := python.PyTuple_New(6)
		python.PyTuple_SetItem(bArgs, 0, PyStr(drivername))
		python.PyTuple_SetItem(bArgs, 1, PyStr(database))
		python.PyTuple_SetItem(bArgs, 2, PyStr(username))
		python.PyTuple_SetItem(bArgs, 3, PyStr(password))
		python.PyTuple_SetItem(bArgs, 4, PyStr(host))
		python.PyTuple_SetItem(bArgs, 5, PyStr(port))
		return bArgs
	}
	getTableNamesFunc := metaModule.GetAttrString("get_table_names")
	bArgs := packParam(drivername, database, username, password, host, port)
	rsp := getTableNamesFunc.Call(bArgs, python.Py_None)
	resp := GoStr(rsp)
	return strings.Split(resp, ",")
}

func getColumn(tablename, drivername, database, username, password, host, port string) []TableInfo {

	packParam := func(tablename string, drivername string, database string, username string, password string, host string, port string) *python.PyObject {
		bArgs := python.PyTuple_New(7)
		python.PyTuple_SetItem(bArgs, 0, PyStr(tablename))
		python.PyTuple_SetItem(bArgs, 1, PyStr(drivername))
		python.PyTuple_SetItem(bArgs, 2, PyStr(database))
		python.PyTuple_SetItem(bArgs, 3, PyStr(username))
		python.PyTuple_SetItem(bArgs, 4, PyStr(password))
		python.PyTuple_SetItem(bArgs, 5, PyStr(host))
		python.PyTuple_SetItem(bArgs, 6, PyStr(port))
		return bArgs
	}
	formatJson := func(rsp *python.PyObject) string {
		resp := GoStr(rsp)
		resp = strings.Replace(resp, "'", "\"", -1)
		resp = strings.Replace(resp, "None", "\"\"", -1)
		resp = strings.Replace(resp, "u\"", "\"", -1)
		resp = strings.Replace(resp, "VARCHAR()", "\"string\"", -1)
		resp = strings.Replace(resp, "TIMESTAMP()", "\"time\"", -1)
		resp = strings.Replace(resp, "BIGINT()", "\"int\"", -1)
		resp = strings.Replace(resp, "False", "false", -1)
		resp = strings.Replace(resp, "True", "true", -1)
		vs := regexp.MustCompile(`next.*\)`).FindAllString(resp, -1)
		for _, v := range vs {
			v2 := strings.Replace(v, "\"", "'", -1)
			resp = strings.Replace(resp, v, v2, -1)
		}
		return resp
	}

	getTableNamesFunc := metaModule.GetAttrString("get_columns")
	bArgs := packParam(tablename, drivername, database, username, password, host, port)
	rsp := getTableNamesFunc.Call(bArgs, python.Py_None)
	resp := formatJson(rsp)

	infoes := make([]TableInfo, 0)
	utee.Chk(json.Unmarshal([]byte(resp), &infoes))
	return infoes
}
