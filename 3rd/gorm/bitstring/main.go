package main

import (
	"bytes"
	"database/sql/driver"
	"time"

	"github.com/RoaringBitmap/roaring"
	"github.com/ahmetb/go-linq/v3"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

var env struct {
	db *gorm.DB
}

func init() {
	pgdb, err := gorm.Open("postgres", "user=figo password=xujianhui0915 dbname=figo host=127.0.0.1 port=5432 sslmode=disable application_name=praticse")
	if err != nil {
		panic(err)
	}
	pgdb.DB().SetConnMaxLifetime(time.Minute * 5)
	pgdb.DB().SetMaxIdleConns(0)
	pgdb.DB().SetMaxOpenConns(5)
	pgdb.LogMode(true)
	pgdb.Debug().AutoMigrate(&BitInfo{})
	env.db = pgdb
}

func main() {
	write()
	read(1)
}

func write() {
	info := &BitInfo{
		BitData: &BitString{
			Bm: roaring.NewBitmap(),
		},
	}
	bitmap := info.BitData.Bm
	bitmap.Add(1)
	bitmap.Add(3)
	bitmap.Add(5)
	bitmap.Add(7)
	env.db.Save(&info)
}

func read(id int) {
	var info BitInfo
	err := env.db.Raw(`select * from bit_infos where id=?`, id).
		Scan(&info).Error
	if err != nil {
		panic(err)
	}
	logrus.WithField("info", info).Println("done")
}

type BitInfo struct {
	ID      int
	BitData *BitString `gorm:"type:BIT VARYING(10000000);"`
}

type BitString struct {
	Bm *roaring.Bitmap
}

func (p *BitString) Scan(src interface{}) error {
	bs := src.([]byte)
	logrus.WithField("src", string(bs)).
		Println("done")

	p.Bm = roaring.NewBitmap()
	for i, v := range bs {
		if v == 49 {
			p.Bm.Add(uint32(i))
		}
	}

	return nil
}

func (p *BitString) Value() (driver.Value, error) {
	query := linq.From(p.Bm.ToArray())
	sb := bytes.NewBufferString(``)
	for i := 0; i <= int(p.Bm.Maximum()); i++ {
		existFlag := query.Contains(uint32(i))
		v := "0"
		if existFlag {
			v = "1"
		}
		sb.WriteString(v)
	}
	return sb.String(), nil
}
