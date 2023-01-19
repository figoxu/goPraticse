package main

import (
	"context"
	"fmt"
	"github.com/figoxu/Figo"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/quexer/utee"
	log "github.com/sirupsen/logrus"
	"time"
)

var db *pg.DB

func init() {
	option := &pg.Options{
		User:         "figo",
		Password:     "xujianhui0915",
		Addr:         "127.0.0.1:5432",
		Database:     "figo",
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 120,
		PoolSize:     10,
	}
	db = pg.Connect(option)
	db.AddQueryHook(&Hooker{})

}

type Hooker struct {
}

func (p *Hooker) BeforeQuery(event *pg.QueryEvent) {
	event.Data["ct"] = utee.Tick()
}

func (p *Hooker) AfterQuery(event *pg.QueryEvent) {
	query, err := event.FormattedQuery()
	if err != nil {
		log.Warn("pg event format err", err)
		return
	}
	log.WithField("mod", "pg").WithField("sql", query).Warnln("sql log")
}

type Email struct {
	Base
	Email  string
	UserID int
}

type Base struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (p *Email) BeforeInsert(c context.Context, db orm.DB) error {
	log.Println("before insert")
	return nil
}

func (p *Email) BeforeUpdate(c context.Context, db orm.DB) error {
	log.Println("before update")
	return nil
}

func (p *Email) AfterInsert(c context.Context, db orm.DB) error {
	log.Println(Figo.JsonString(p))
	log.Println("after insert")
	return nil
}

func (p *Email) AfterUpdate(c context.Context, db orm.DB) error {
	log.Println("after update")
	return nil
}

func (p *Email) AfterDelete(c context.Context, db orm.DB) error {
	log.Println("after delete")
	return nil
}

func main() {
	db.Insert(&Email{
		Email:  "xujh945@qq.com",
		UserID: 1,
	})
	email := Email{}
	db.Model(&email).First()
	fmt.Println(Figo.JsonString(email))
}
