package main

import (
	mapset "github.com/deckarep/golang-set"
	"log"
)

func main() {
	a:=mapset.NewSet(1, 2, 3, 4, 5, 6, 7)
	b:=mapset.NewSet(2, 4, 6, 8, 9, 10)
	log.Println(a.Difference(b))
	log.Println(b.Difference(a))
	log.Println(b.Intersect(a))
	log.Println(b.SymmetricDifference(a))
}
