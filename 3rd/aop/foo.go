package main

import (
	"fmt"
	"github.com/gogap/aop"
)

type Foo struct {
}

// @AfterReturning, the method could have args of aop.Result,
// it will get the result from real func return values
func (p *Foo) Bar(result aop.Result) {
	result.MapTo(func(v bool) {
		fmt.Println("Bar Bar Bar .... Result is:", v)
	})
}
