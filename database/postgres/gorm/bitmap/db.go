package main

import (
	"github.com/jinzhu/gorm"
	"time"
	"github.com/figoxu/Figo"
	"github.com/quexer/utee"
	_ "github.com/lib/pq"
)

var (
	pg_plsql *gorm.DB
)

func initDB() {
	pg_db_plsql, err := gorm.Open("postgres", "user=figo password=123456 dbname=plpgsql host=127.0.0.1 port=5432 sslmode=disable application_name=bitmapfaker")
	pg_db_plsql.DB().SetConnMaxLifetime(time.Minute * 5)
	pg_db_plsql.DB().SetMaxIdleConns(0)
	pg_db_plsql.DB().SetMaxOpenConns(5)
	pg_db_plsql.SetLogger(&Figo.GormLog{})
	utee.Chk(err)
	pg_db_plsql.LogMode(true)
	pg_db_plsql.SingularTable(true)
	pg_db_plsql.Debug().AutoMigrate(&BitUser{}, &BitUserFriend{}, BitUserFriendMap{})
	pg_plsql = pg_db_plsql
}

type BitUser struct {
	Id   int
	Name string
	Age  int
	Sex  string
}

type BitUserDao struct {
	db *gorm.DB
}

func NewBitUserDao(db *gorm.DB) BitUserDao {
	return BitUserDao{
		db: db,
	}
}

func (p *BitUserDao) Save(bitUser *BitUser) {
	p.db.Save(bitUser)
}

type BitUserFriend struct {
	Id     int
	Uid    int
	Friend int
}

type BitUserFriendDao struct {
	db *gorm.DB
}

func (p *BitUserFriendDao) Save(bitUserFriend *BitUserFriend) {
	p.db.Save(bitUserFriend)
}

type BitUserFriendMap struct {
	Id    int
	Uid   int
	BM8KW string `sql:"type:bit varying"`
}
