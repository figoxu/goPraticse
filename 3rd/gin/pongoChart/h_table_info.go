package main

import (
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2"
)

func h_table_info_index(c *gin.Context) {
	c.HTML(200, "tableinfo.html", pongo2.Context{})
}