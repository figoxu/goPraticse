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
)

var (
	sqlite_db *gorm.DB
)

func init() {
	driver := "sqlite3"
	dbLoc := "./db/test.db"
	db, err := gorm.Open(driver, dbLoc)
	utee.Chk(err)
	db.DB().SetConnMaxLifetime(time.Minute * 5)
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(5)
	db.SingularTable(true)
	db.Debug().AutoMigrate(&DbInfo{}, &TableInfo{})
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
