package main

import (
	"fmt"
	"github.com/gogap/aop"
)

type Auth struct {
}

func (p *Auth) Login(userName, password string) bool {
	fmt.Println("@Login", userName, password)
	if userName == "zeal" && password == "gogap" {
		return true
	}
	return false
}

// use join point to get Args from real method
func (p *Auth) Before(jp aop.JoinPointer) {
	username := ""
	jp.Args().MapTo(func(u, p string) {
		username = u
	})

	fmt.Printf("Before Login: %s\n", username)
}

// the args is same as Login
func (p *Auth) After(username, password string) {
	fmt.Printf("After Login: %s %s\n", username, password)
}

// use join point to around the real func of login
func (p *Auth) Around(pjp aop.ProceedingJoinPointer) {
	fmt.Println("@Around Begin ")

	ret := pjp.Proceed("fakeName", "fakePassword")
	ret.MapTo(func(loginResult bool) {
		fmt.Println("@Proceed Result is", loginResult)
	})

	fmt.Println("@Around End ")
}
