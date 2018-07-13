package main

import (
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2"
)

func h_chart_index(c *gin.Context) {
	c.HTML(200, "index.html", pongo2.Context{})
}
