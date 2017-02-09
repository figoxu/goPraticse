package main

import (
	"log"
	"github.com/gogap/aop"
)

const (
	pointcut_name    = "pointcut_1"
	pointcut_method  = "Login()"
	aspect_name_auth = "auth"
	aspect_name_foo  = "foo"
)

func initAop() *aop.AOP {
	gogapAop := aop.NewAOP()
	pointcut := aop.NewPointcut(pointcut_name).Execution(pointcut_method)
	pointcut.Execution(pointcut_method)

	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean(aspect_name_auth, new(Auth))
	beanFactory.RegisterBean(aspect_name_foo, new(Foo))
	gogapAop.SetBeanFactory(beanFactory)

	aspect := aop.NewAspect("aspect_1", aspect_name_auth)
	aspect.SetBeanFactory(beanFactory)
	aspect.AddPointcut(pointcut)
	aspect.AddAdvice(&aop.Advice{Ordering: aop.Before, Method: "Before", PointcutRefID: pointcut_name})
	aspect.AddAdvice(&aop.Advice{Ordering: aop.After, Method: "After", PointcutRefID: pointcut_name})
	aspect.AddAdvice(&aop.Advice{Ordering: aop.Around, Method: "Around", PointcutRefID: pointcut_name})

	aspectFoo := aop.NewAspect("aspect_2", aspect_name_foo)
	aspectFoo.SetBeanFactory(beanFactory)
	aspectFoo.AddPointcut(pointcut)
	aspectFoo.AddAdvice(&aop.Advice{Ordering: aop.AfterReturning, Method: "Bar", PointcutRefID: pointcut_name})

	gogapAop.AddAspect(aspect)
	gogapAop.AddAspect(aspectFoo)
	return gogapAop
}

func main() {
	gogapAop := initAop()

	var err error
	var proxy *aop.Proxy
	// Get proxy
	if proxy, err = gogapAop.GetProxy(aspect_name_auth); err != nil {
		log.Println("get proxy failed", err)
		return
	}

	// start trace for debug
	aop.StartTrace()
	log.Println("Invoke By Method")
	resultByMethod := proxy.Method(new(Auth).Login).(func(string, string) bool)("zeal", "gogap")
	log.Println("Invoke By Proxy")
	if err = proxy.Invoke(new(Auth).Login, "zeal", "errorpassword").End(
		func(resultByInvoke bool) {
			resultByMethod = resultByInvoke
		}); err != nil {
		log.Println("invoke proxy func error", err)
	} else {
		log.Println("Login result:", resultByMethod)
	}

	log.Println("====[StopTrace]====>")
	printTrace()
}

func printTrace() {
	t, _ := aop.StopTrace()
	for _, item := range t.Items() {
		log.Println(item.ID, item.InvokeID, item.BeanRefID, item.Pointcut, item.Method)
	}
}
