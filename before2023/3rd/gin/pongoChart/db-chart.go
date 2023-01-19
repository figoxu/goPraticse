package main

import (
	"fmt"
)

type DbColumn struct {
	Column   string
	ShowName string
	Operator string
}

type ChartDefine struct {
	Name        string
	ChartType   string
	Dimensions   []DbColumn
	Measurements []DbColumn
	Table       string
}

func (p *ChartDefine) GenSql() string {
	return p.GenSqlWithWhere("")
}

func (p *ChartDefine) GenSqlWithWhere(where string) string {
	sqlStr:=fmt.Sprint("SELECT ")
	for i,dimension:=range p.Dimensions {
		if i>0 {
			sqlStr = fmt.Sprint(sqlStr,",")
		}
		sqlStr=fmt.Sprint(sqlStr," ",dimension.Column)
	}
	for i,measurement:=range p.Measurements {
		if i+len(p.Measurements)==0 {
			sqlStr = fmt.Sprint(sqlStr,",")
		}
		sqlStr=fmt.Sprint(sqlStr,",",measurement.Operator,"(",measurement.Column,")")
	}
	sqlStr=fmt.Sprint(sqlStr," FROM ",p.Table," ",where)
	for i,dimension:=range p.Dimensions {
		if i==0 {
			sqlStr = fmt.Sprint(sqlStr," GROUP BY ")
		}else{
			sqlStr = fmt.Sprint(sqlStr," , ")
		}
		sqlStr=fmt.Sprint(sqlStr," ",dimension.Column)
	}
	return sqlStr
}
