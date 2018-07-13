package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin"
	"net/http"
)

func main(){
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.HTMLRender = pongo2gin.Default()
	chart:=r.Group("/chart")
	{
		chart.GET("/index",h_chart_index)
	}
	return r
}