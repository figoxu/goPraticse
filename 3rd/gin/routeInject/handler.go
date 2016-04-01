package main
import (
	"github.com/gin-gonic/gin"
	"log"
)



func A1MiddleWare(c *gin.Context){
	obj := &InjectObj{
		val : "cool u,handsome me",
	}
	c.Set("iObj",obj)
	log.Println(" before  next ")
	c.Next()
	log.Println(" after  next ")

}

func A2MiddleWare(c *gin.Context){
	obj, ok := c.Keys["iObj"].(*InjectObj)
	if !ok {
		log.Panic(" inject error ")
	}
	obj.Println()
}

func AHanlder(c *gin.Context){
	log.Println("....AHandler invoke....")
}

func BMiddleWare(c *gin.Context){
	obj := &InjectObj{
		val : "i am rocker Mr'B",
	}
	c.Set("iObj",obj)
}

func BHandler(c *gin.Context){
	obj, ok := c.Keys["iObj"].(*InjectObj)
	if !ok {
		log.Panic(" inject error ")
	}
	obj.Println()
}