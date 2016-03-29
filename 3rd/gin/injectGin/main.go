package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetDummyEndpoint(c *gin.Context) {
	resp := map[string]string{"hello": "world"}
	c.JSON(200, resp)
}

func main() {
	api := gin.Default()
	api.Use(DummyMiddleware)
	api.GET("/dummy", GetDummyEndpoint)
	api.GET("/dummy2", GetDummyEndpoint)
	api.GET("/dummy3", func(c *gin.Context) {
		if v, ok := c.Get("hello"); ok {
			log.Println("--------------")
			log.Println(v)
		}
	})
	api.Run(":5000")
}

func DummyMiddleware(c *gin.Context) {
	c.Set("hello", "world") // mapping obj using map
	uri := c.Request.RequestURI
	if uri == "/dummy2" { // abort the request using middleware
		resp := map[string]string{"figo": "xu"}
		c.JSON(200, resp)
		c.Abort()
		return
	}
	log.Println("Im a dummy!")
	c.Next()
}
