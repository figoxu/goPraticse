package main

import (
	snk "github.com/vladimirvivien/automi/sinks/csv"
	src "github.com/vladimirvivien/automi/sources/csv"
	"github.com/vladimirvivien/automi/stream"
	"log"
	"strconv"
)

type scientist struct {
	FirstName string
	LastName  string
	Title     string
	BornYear  int
}

func main() {
	in := src.New().WithFile("C:/GoPath/src/github.com/figoxu/goPraticse/3rd/stream/automi/csvStreamProcess/data.txt")
	out := snk.New().WithFile("C:/GoPath/src/github.com/figoxu/goPraticse/3rd/stream/automi/csvStreamProcess/result.txt")
	s:=stream.New().From(in)
	s.Map(translate2Struct).Filter(filterByYear).Map(mapOutput).To(out)
	log.Println("hello")
	<-s.Open()
}

func translate2Struct(cs []string) scientist {
	bornYear, _ := strconv.Atoi(cs[3])
	return scientist{
		FirstName: cs[1],
		LastName:  cs[0],
		Title:     cs[2],
		BornYear:  bornYear,
	}
}

func filterByYear(s scientist) bool{
	if s.BornYear > 1930 {
		return true
	}
	return false
}

func mapOutput(s scientist) []string {
	return []string{s.FirstName, s.LastName, s.Title}
}
