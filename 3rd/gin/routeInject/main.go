package main
import (
	"log"
	"github.com/gin-gonic/gin"
)

type InjectObj struct {
	val string
}

func (p *InjectObj) Println(){
	log.Println("i am injector with value :",p.val)
}

func main(){
	api := gin.Default()
	api.GET("/sample/cool", A1MiddleWare, A2MiddleWare,AHanlder)
	api.GET("/sample/you", BMiddleWare, BHandler)
	api.Run(":5000")
}

