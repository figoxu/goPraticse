package main

import (
	"testing"
	"fmt"
)

func TestChartDefine_GenSql(t *testing.T) {
	cd := ChartDefine{
		Name:      "FigoChart",
		ChartType: "LineChart",
		Dimensions: []DbColumn{
			DbColumn{
				Column:   "province",
				ShowName: "薪水",
				Operator: "",
			}, DbColumn{
				Column:   "sex",
				ShowName: "性别",
				Operator: "",
			},
		},
		Measurements: []DbColumn{DbColumn{
			Column:   "salary",
			ShowName: "薪水",
			Operator: "sum",
		},},
		Table: "test_table",
	}
	fmt.Println(cd.GenSql())
}
