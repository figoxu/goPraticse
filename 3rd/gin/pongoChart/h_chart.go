package main

import (
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	"net/http"
	"github.com/figoxu/goPraticse/3rd/gin/pongoChart/common/db"
	"github.com/figoxu/goPraticse/3rd/gin/pongoChart/common/pg"
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
	dbInfoDao, tableInfoDao := db.NewDbInfoDao(sqlite_db), db.NewTableInfoDao(sqlite_db)
	dbInfoes, tableInfoes := dbInfoDao.GetAll(), tableInfoDao.GetAll()
	c.HTML(200, "chart.html", pongo2.Context{
		"tagMap":      Figo.JsonString(tagMap),
		"dbInfoes":    Figo.JsonString(dbInfoes),
		"tableInfoes": Figo.JsonString(tableInfoes),
		"opMap":       Figo.JsonString(operatorMap),
	})
}

type ChartRow map[string]interface{}

type ChartQueryResult struct {
	Columns []string   `json:"columns"`
	Rows    []ChartRow `json:"rows"`
}

func h_chart_query(c *gin.Context) {
	dbInfoDao, tableInfoDao := db.NewDbInfoDao(sqlite_db), db.NewTableInfoDao(sqlite_db)
	env := c.MustGet("env").(*Env)
	fh := env.fh
	v_db, v_table, v_op := fh.Int("v_db"), fh.String("v_table"), fh.String("v_op")
	v_dimensions, v_measurements := fh.IntArr("v_dimension", ","), fh.IntArr("v_measurement", ",")

	dbInfo := dbInfoDao.GetByKey(v_db)

	dimensions, measurements, columns := make([]DbColumn, 0), make([]DbColumn, 0), make([]string, 0)
	for _, dimension := range v_dimensions {
		tableInfo := tableInfoDao.GetByKey(dimension)
		dimensions = append(dimensions, DbColumn{
			Column:   tableInfo.Name,
			ShowName: tableInfo.Name,
			Operator: "",
		})
	}
	for _, measurement := range v_measurements {
		tableInfo := tableInfoDao.GetByKey(measurement)
		measurements = append(measurements, DbColumn{
			Column:   tableInfo.Name,
			ShowName: tableInfo.Name,
			Operator: v_op,
		})
	}
	cd := ChartDefine{
		Dimensions:   dimensions,
		Measurements: measurements,
		Table:        v_table,
	}
	chartSql := cd.GenSql()
	results := make([]ChartRow, 0)
	con := pg.BuildDbCon(dbInfo)
	data := pg.Query(con, chartSql)
	for _, d := range data {
		row := ChartRow{}
		for k, v := range d {
			row[k] = v
			columns = append(columns,k)
		}
		results = append(results, row)
	}
	columns = utee.UniqueStr(columns)
	chartQueryResult := ChartQueryResult{
		Columns: columns,
		Rows:    results,
	}
	c.JSON(http.StatusOK, chartQueryResult)
}
