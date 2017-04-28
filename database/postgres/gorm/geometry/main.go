package main

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/figoxu/utee"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	DB, err = gorm.Open("postgres", "user=xxx password=xxx dbname=sdz_dev host=47.93.118.90 port=3432 sslmode=disable")
	utee.Chk(err)
	DB.LogMode(true)
}

type TestPoint struct {
	Location GeoPoint `sql:"type:geometry(Geometry,4326)"`
}

func main() {
	if DB.CreateTable(&TestPoint{}) == nil {
		log.Fatal("Can't create table")
	}
	p := TestPoint{
		Location: GeoPoint{
			Lat: 43.76857094631136,
			Lng: 11.292383687705296,
		},
	}
	if DB.Create(p) == nil {
		log.Fatal("Can't create row")
	}
	var res TestPoint
	DB.First(&res)
	if res.Location.Lat != 43.76857094631136 {
		log.Fatal("Latitude not correct")
	}
	log.Println("00006")
	if res.Location.Lng != 11.292383687705296 {
		log.Fatal("Longitude not correct")
	}
	log.Println("00007")
}

type GeoPoint struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (p *GeoPoint) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lng, p.Lat)
}

func (p *GeoPoint) Scan(val interface{}) error {
	b, err := hex.DecodeString(string(val.([]uint8)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	}

	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}

	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}
