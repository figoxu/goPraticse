package pg

import (
	"github.com/jinzhu/gorm"
	"github.com/figoxu/goPraticse/3rd/gin/pongoChart/common/db"
	"github.com/quexer/utee"
	"time"
	"fmt"
)

func BuildDbCon(dbInfo db.DbInfo) *gorm.DB {

	buildPgStr := func() string {
		pgStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable application_name=chart", dbInfo.Username, dbInfo.Password, dbInfo.Database, dbInfo.Host, dbInfo.Port)
		return pgStr
	}
	pg, err := gorm.Open("postgres", buildPgStr())
	utee.Chk(err)
	pg.DB().SetConnMaxLifetime(time.Minute * 5)
	pg.DB().SetMaxIdleConns(0)
	pg.DB().SetMaxOpenConns(5)
	pg.LogMode(true)
	pg.SingularTable(true)
	return pg
}
