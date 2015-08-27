package main

import (
	"strconv"
	"time"
)

const (
	DATE_FORMAT_YmD   = "20060102"
	DATE_FORMAT_YmDH   = "2006010215"
	DATE_FORMAT_YmDHMS = "2006 01 02 15 04 05"
)

func getDay(day int, format string) int {
	t := time.Now()
	t = t.Add(time.Hour * 24 * time.Duration(day))
	i, _ := strconv.Atoi(t.Format(format))
	return i
}
