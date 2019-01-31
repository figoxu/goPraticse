package main

import (
	"github.com/quexer/utee"
	"github.com/tealeg/xlsx"
	"io"
)

type ExcelWriter struct {
	file  *xlsx.File
	sheet *xlsx.Sheet
	w     io.Writer
}

func NewExcelWriter(w io.Writer) *ExcelWriter {
	xw := &ExcelWriter{
		file: xlsx.NewFile(),
		w:    w,
	}
	sheet, err := xw.file.AddSheet("Sheet1")
	utee.Chk(err)
	xw.sheet = sheet

	return xw
}

func (p *ExcelWriter) Write(record []string) {
	row := p.sheet.AddRow()
	for _, s := range record {
		cell := row.AddCell()
		cell.Value = s
	}
}

func (p *ExcelWriter) Flush() error {
	return p.file.Write(p.w)
}

func (p *ExcelWriter) WriteWithMergeCell(record []string, indexToMerge []int, hcells, vcells int) {
	row := p.sheet.AddRow()
	for i, s := range record {
		cell := row.AddCell()
		if utee.ContainsInt(indexToMerge, i) && (hcells > 0 || vcells > 0) {
			cell.Merge(hcells, vcells)
		}
		cell.Value = s
	}
}

