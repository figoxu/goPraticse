package main

import (
	"github.com/gin-gonic/gin"
	"github.com/flosch/pongo2"
	"net/http"
)

func h_db_info_index(c *gin.Context) {
	c.HTML(200, "dbinfo.html", pongo2.Context{})
}

func h_db_info_save(c *gin.Context) {
	dbInfo := DbInfo{}
	c.BindJSON(&dbInfo)
	dbInfoDao := NewDbInfoDao(sqlite_db)
	tableInfoDao := NewTableInfoDao(sqlite_db)
	dbInfoDao.Save(&dbInfo)
	tableNames := getTableNames(dbInfo)
	for _, tableName := range tableNames {
		tableInfoes := getColumn(tableName, dbInfo)
		for _, tableInfo := range tableInfoes {
			tableInfo.DbId = dbInfo.Id
			tableInfo.TableName = tableName
			tableInfoDao.Save(&tableInfo)
		}
	}
}

func h_db_info_list(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	ph := env.ph
	pg, size := ph.Int("pg"), ph.Int("size")
	count := &TCount{}
	sqlite_db.Raw("SELECT count(*) as count FROM db_info").Scan(count)
	totalPg := (count.Count + size - 1) / size
	if pg <= 0 {
		pg = 1
	} else if pg > totalPg {
		pg = totalPg
	}
	start, limit := (pg-1)*size, size
	confs := make([]DbInfo, 0)
	sqlite_db.Raw("SELECT * FROM db_info ORDER BY id DESC LIMIT ?,?", start, limit).Scan(&confs)

	result := make(map[string]interface{})
	result["total"] = count.Count
	result["totalPg"] = totalPg
	result["curPg"] = pg
	result["data"] = confs
	c.JSON(http.StatusOK, result)
}
