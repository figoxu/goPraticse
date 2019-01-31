package main

import (
	"github.com/quexer/utee"
	"github.com/tealeg/xlsx"
	"io"
)

type ExcelWriter struct {
	file  *xlsx.StreamFile
}

func NewExcelWriter(w io.Writer) *ExcelWriter {
	builder := xlsx.NewStreamFileBuilder(w)
	builder.AddSheet("Sheet1", []string{}, []*xlsx.CellType{})
	file, err := builder.Build()
	utee.Chk(err)
	xw := &ExcelWriter{
		file: file,
	}
	return xw
}

func (p *ExcelWriter) Write(record []string) {
	err:=p.file.Write(record)
	utee.Chk(err)
}

func (p *ExcelWriter) Flush() error {
	p.file.Flush()
	return nil
}
