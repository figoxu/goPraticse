package grace

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/facebookgo/grace/gracehttp"
)


func main() {
	log.Println("STARTED")
	r := gin.Default()
	r.GET("/hello", h_hello)
	gracehttp.Serve(
		&http.Server{Addr: ":10062", Handler: r},
	)
}

func h_hello(c *gin.Context){
	c.String(http.StatusOK,"world")
}
