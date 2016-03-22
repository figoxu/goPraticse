package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"github.com/figoxu/utee"
)

func main() {
	excelFileName := "/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/3rd/excel/readExcel/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	utee.Chk(err)
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				v,err:=cell.String()
				utee.Chk(err)
				fmt.Printf("%s\n", v)
			}
		}
	}
}
