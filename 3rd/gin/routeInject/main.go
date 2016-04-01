package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"regexp"
)

type InjectObj struct {
	val string
}

func (p *InjectObj) Println() {
	log.Println("i am injector with value :", p.val)
}

func main() {
	api := gin.Default()
	//	api.GET("/sample/cool", A1MiddleWare, A2MiddleWare,AHanlder)
	//	api.GET("/sample/you", BMiddleWare, BHandler)

	api.GET("/sample/:param", midSmart, handleWithSmart)

	api.Run(":5000")
}

type SmartHandler map[string]*gin.HandlersChain

func midSmart(c *gin.Context) {
	smart := make(SmartHandler)
	smart["/cool"] = &gin.HandlersChain{A1MiddleWare, A2MiddleWare, AHanlder}
	smart["/you"] = &gin.HandlersChain{BMiddleWare, BHandler}
	c.Set("s",smart)
}

func handleWithSmart(c *gin.Context) {
	smart, ok := c.Keys["s"].(SmartHandler)
	if !ok {
		log.Panic(" inject error ")
	}
	requestUri := c.Request.RequestURI
	for k, h := range smart {
		if( regexp.MustCompile(k).FindString(requestUri) != "" ){
			log.Println("found match @k:", k, " @h:", h)
			ginExecute(c,*h)
			return
		}
	}
	log.Println("found nothing")
}

func ginExecute(c *gin.Context,handlerChain gin.HandlersChain){
	c.Set("handlerChain",handlerChain)
	c.Set("gin_next_idx",int8(0))
	handlerChain[0](c)
	ginIdx,_:=c.Keys["gin_next_idx"].(int8)
	s := int8(len(handlerChain))
	for ; ginIdx < s; ginIdx++ {
		handlerChain[ginIdx](c)
	}
}


func ginNext(c *gin.Context){
	handlerChain, ok := c.Keys["handlerChain"].(gin.HandlersChain)
	if !ok {
		c.Next()
		return
	}
	ginIdx,ok:=c.Keys["gin_next_idx"].(int8)
	if !ok {
		log.Println("before reset :",ginIdx)
		ginIdx=0
		log.Println("after reset :",ginIdx)
	}
	ginIdx++
	s := int8(len(handlerChain))
	for ; ginIdx < s; ginIdx++ {
		handlerChain[ginIdx](c)
	}
}




