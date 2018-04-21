package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	r := gin.Default()
	r.GET("/hello/world/json", h_helloWorldJson)
	r.GET("/hello/name/:name", h_helloByName)
	r.GET("/hello/redirect/:url", h_helloRedirect)
	r.Run(":8080")
}

func h_helloWorldJson(c *gin.Context) {
	c.SecureJSON(http.StatusOK, map[string]string{
		"foo":   "bar",
		"hello": "world",
		"figo":  "xu",
	})
}

func h_helloByName(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

// http://localhost:8080/hello/redirect/www.baidu.com
func h_helloRedirect(c *gin.Context){
	url:=c.Param("url")
	c.Redirect(http.StatusMovedPermanently,fmt.Sprintf("http://%s",url))
}
