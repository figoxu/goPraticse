package main

import (
	"errors"
	"github.com/jinzhu/gorm"
	"reflect"
)

type BaseDao struct {
}

func (p *BaseDao) Insert(db *gorm.DB, value interface{}) error {
	return db.Model(value).Save(value).Error
}

func (p *BaseDao) GetById(db *gorm.DB, id int, value interface{}) error {
	return db.Model(value).Where("id=?", id).Scan(value).Error
}

func (p *BaseDao) Update(db *gorm.DB, v interface{}, fields ...string) error {
	val := reflect.ValueOf(v).Elem()
	base := val.FieldByName("Base").Interface().(Base)
	if base.IsPersistent() {
		return db.Debug().Model(v).Select(fields).Update(v).Error
	}
	return errors.New("Can't Update All Table, Id Is Empty")
}
