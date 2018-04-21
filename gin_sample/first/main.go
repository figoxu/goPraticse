package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello/world/json", h_helloWorldJson)
	r.Run(":8080")
}

func h_helloWorldJson(c *gin.Context) {
	c.SecureJSON(http.StatusOK, map[string]string{
		"foo":"bar",
		"hello":"world",
		"figo":"xu",
	})
}