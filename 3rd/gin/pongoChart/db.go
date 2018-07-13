package main

import "github.com/jinzhu/gorm"

type DbInfo struct {
	Id         int
	Name       string
	Drivername string
	Database   string
	Username   string
	Password   string
	Host       string
	Port       string
	TableCount int `gorm:"-"`
}

type DbInfoDao struct {
	db *gorm.DB
}

func NewDbInfoDao(db *gorm.DB) DbInfoDao {
	return DbInfoDao{
		db: db,
	}
}

func (p *DbInfoDao) Save(dbInfo *DbInfo) {
	p.db.Save(dbInfo)
}

type TableInfo struct {
	Id            int
	Comment       string `json:"comment"`
	Name          string `json:"name"`
	Nullable      bool   `json:"nullable"`
	Default       string `json:"default"`
	Autoincrement bool   `json:"autoincrement"`
	Type          string `json:"type"`
}

type TableInfoDao struct {
	db *gorm.DB
}

func NewTableInfoDao(db *gorm.DB) TableInfoDao {
	return TableInfoDao{
		db: db,
	}
}

func (p *TableInfoDao) Save(tableInfo *TableInfo){
	p.db.Save(tableInfo)
}
