package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main(){
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	// Define your handlers
	r.GET("/", func(c *gin.Context){
		c.String(200, "Hello World!")
	})
	r.GET("/ping", func(c *gin.Context){
		c.String(200, "pong")
	})

	// Handle all requests using net/http
	http.Handle("/", r)
	http.ListenAndServe(":9090", nil)

}
