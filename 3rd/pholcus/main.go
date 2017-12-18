package main

import (
	"github.com/henrylee2cn/pholcus/exec"
	_ "github.com/henrylee2cn/pholcus_lib" // 此为公开维护的spider规则库
)

func main() {
	exec.DefaultRun("web")
}

