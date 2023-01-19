package pg

import (
	"github.com/figoxu/goPraticse/3rd/gin/pongoChart/common/db"
	"github.com/quexer/utee"
	"fmt"
	"database/sql"
	"github.com/elgs/gosqljson"
)

func BuildDbCon(dbInfo db.DbInfo) *sql.DB {
	buildPgStr := func() string {
		pgStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable application_name=chart", dbInfo.Username, dbInfo.Password, dbInfo.Database, dbInfo.Host, dbInfo.Port)
		return pgStr
	}
	db, err := sql.Open("postgres", buildPgStr())
	utee.Chk(err)
	return db
}

func Query(db *sql.DB, sqlQuery string) []map[string]string {
	data, err := gosqljson.QueryDbToMap(db, "lower", sqlQuery)
	utee.Chk(err)
	return data
}
