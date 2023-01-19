package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"github.com/jinzhu/gorm"
	"github.com/quexer/utee"
	"time"
	"github.com/figoxu/gh"
	"github.com/figoxu/goPraticse/3rd/gin/pongoChart/common/db"
)

var (
	sqlite_db *gorm.DB
)

func init() {
	driver := "sqlite3"
	dbLoc := "./db/test.db"
	sqlitedb, err := gorm.Open(driver, dbLoc)
	utee.Chk(err)
	sqlitedb.DB().SetConnMaxLifetime(time.Minute * 5)
	sqlitedb.DB().SetMaxIdleConns(0)
	sqlitedb.DB().SetMaxOpenConns(5)
	sqlitedb.SingularTable(true)
	sqlitedb.Debug().AutoMigrate(&db.DbInfo{}, &db.TableInfo{})
	sqlite_db = sqlitedb
}

func main() {
	engine := mount()
	http.Handle("/", engine)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func mount() *gin.Engine {
	r := gin.Default()
	r.HTMLRender = pongo2gin.Default()
	chart := r.Group("/chart",m_gh)
	{
		chart.GET("/index", h_chart_index)
	}
	admin := r.Group("/admin",m_gh)
	{
		db:=admin.Group("/db")
		{
			db.POST("/save",h_db_info_save)
			db.GET("/index",h_db_info_index)
			db.POST("/list/:size/:pg",h_db_info_list)
		}
		admin.GET("/index",h_table_info_index)
		chart:=admin.Group("/chart")
		{
			chart.GET("/index", h_chart_define_index)
			chart.POST("/query",h_chart_query)
		}
	}
	return r
}

type Env struct {
	fh *gh.FormHelper
	ph *gh.ParamHelper
}

func m_gh(c *gin.Context) {
	c.Set("env", &Env{
		fh: gh.NewFormHelper(c),
		ph: gh.NewParamHelper(c),
	})
	c.Next()
}
