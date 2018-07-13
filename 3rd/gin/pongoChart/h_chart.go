package main

import (
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2"
	"github.com/figoxu/Figo"
)

var (
	tagMap = map[string]string{
		"折线图": "ve-line",
		"柱状图": "ve-histogram",
		"条形图": "ve-bar",
		"饼图":  "ve-pie",
	}
	operatorMap = map[string]string{
		"计次":  "count",
		"总数":  "sum",
		"平均数": "avg",
	}
)

func h_chart_index(c *gin.Context) {
	c.HTML(200, "index.html", pongo2.Context{})
}

func h_chart_define_index(c *gin.Context) {
	dbInfoDao, tableInfoDao := NewDbInfoDao(sqlite_db), NewTableInfoDao(sqlite_db)
	dbInfoes, tableInfoes := dbInfoDao.GetAll(), tableInfoDao.GetAll()
	c.HTML(200, "chart.html", pongo2.Context{
		"tagMap":      Figo.JsonString(tagMap),
		"dbInfoes":    Figo.JsonString(dbInfoes),
		"tableInfoes": Figo.JsonString(tableInfoes),
		"opMap":       Figo.JsonString(operatorMap),
	})
}
