package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"github.com/quexer/utee"
)

func main() {

	excelFileName := "/home/figo/develop/env/GOPATH/src/github.com/figoxu/goPraticse/files/excel/read/test.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	utee.Chk(err)
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				s,err := cell.String()
				utee.Chk(err)
				fmt.Printf("%s\n", s)
			}
		}
	}
}
