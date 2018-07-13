package main

import "github.com/jinzhu/gorm"

type DbInfo struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Drivername string `json:"drivername"`
	Database   string `json:"database"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	TableCount int    `gorm:"-" json:"tableCount"`
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

func (p *DbInfoDao) GetAll() []DbInfo {
	dbInfos := make([]DbInfo, 0)
	p.db.Raw("SELECT * FROM db_info").Scan(&dbInfos)
	return dbInfos
}

type TableInfo struct {
	Id            int    `json:"id"`
	TableName     string `json:"table_name"`
	DbId          int    `json:"db_id"`
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

func (p *TableInfoDao) Save(tableInfo *TableInfo) {
	p.db.Save(tableInfo)
}

func (p *TableInfoDao) GetAll() []TableInfo {
	tableInfoes := make([]TableInfo, 0)
	p.db.Raw("SELECT * FROM table_info").Scan(&tableInfoes)
	return tableInfoes
}

type TCount struct {
	Count int
}
