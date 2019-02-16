package ormsample

import (
	"github.com/figoxu/Figo"
	"github.com/jinzhu/gorm"
	"github.com/quexer/utee"
	"time"
)

var env struct {
	db *gorm.DB
}

func init() {
	pgdb, err := gorm.Open("postgres", "user=figo password=xujianhui0915 dbname=figo host=127.0.0.1 port=5432 sslmode=disable application_name=praticse")
	pgdb.DB().SetConnMaxLifetime(time.Minute * 5)
	pgdb.DB().SetMaxIdleConns(0)
	pgdb.DB().SetMaxOpenConns(5)
	pgdb.SetLogger(&Figo.GormLog{})
	utee.Chk(err)
	pgdb.LogMode(true)
	pgdb.Debug().AutoMigrate(&User{}, &CreditCard{}, &Email{})
	env.db = pgdb
}
