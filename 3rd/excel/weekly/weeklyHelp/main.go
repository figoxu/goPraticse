package main

import (
	"fmt"
	"github.com/figoxu/figo"
	"github.com/quexer/utee"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"path/filepath"
)

var (
	m map[string][]string
)

func appendResult(category, content string) {
	if m == nil {
		m = make(map[string][]string)
	}
	if m[category] == nil {
		m[category] = []string{}
	}
	m[category] = append(m[category], content)
}

func main() {
	fnames := []string{}
	filepath.Walk("./", func(path string, fi os.FileInfo, err error) error {
		fnames = append(fnames, path)
		return nil
	})
	for _, b := range fnames {
		read(b)
	}
	exportExcel()
}

func exportExcel() {
	taskId := 1
	for k, vs := range m {
		fmt.Println("进行中 ", taskId, ">", k)
		for idx, s := range vs {
			rs := []rune(s)
			numId := fmt.Sprint(taskId, ".", idx, ">")
			fmt.Println(" ", string(rs[0:4]), numId, string(rs[4:len(rs)]))
		}
		taskId++
	}
}

func read(fname string) {
	defer Figo.Catch()
	xlFile, err := xlsx.OpenFile(fname)
	utee.Chk(err)
	for _, sheet := range xlFile.Sheets {
		for rowIdex, row := range sheet.Rows {
			if rowIdex < 2 {
				continue
			}
			category, status, info := "", "", ""
			for i, cell := range row.Cells {
				v, err := cell.String()
				utee.Chk(err)
				if i == 1 {
					info = v
				} else if i == 2 {
					status = v
				} else if i == 4 {
					category = v
				}
			}
			if category == "" && status == "" && info == "" {
				continue
			}
			appendResult(category, fmt.Sprint(status, " ", info))
		}
	}
}
