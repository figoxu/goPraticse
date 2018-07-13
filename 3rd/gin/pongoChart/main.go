package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin"
	"net/http"
	"github.com/astaxie/beego/orm"
	"github.com/jinzhu/gorm"
	"github.com/quexer/utee"
)

var (
	sqlite_db *gorm.DB
)

func init() {
	orm.RegisterModel(new(DbInfo), new(TableInfo))
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	driver := "sqlite3"
	dbLoc := "./datas/test.db"
	orm.RegisterDataBase("default", driver, dbLoc)
	orm.RunSyncdb("default", false, true)
	db, err := gorm.Open(driver, dbLoc)
	utee.Chk(err)
	sqlite_db = db
}

func main() {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.HTMLRender = pongo2gin.Default()
	chart := r.Group("/chart")
	{
		chart.GET("/index", h_chart_index)
	}
	return r
}
