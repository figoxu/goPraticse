package main

import (
	"strconv"
	"strings"

	"github.com/vladimirvivien/automi/api/tuple"
	"github.com/vladimirvivien/automi/sinks/csv"
	"github.com/vladimirvivien/automi/stream"
)

func main() {
	src := stream.NewSliceSource("Hello World", "Hello Milkyway", "Hello Universe")
	snk := csv.New().WithFile("./wc.out")

	stream := stream.New().From(src)

	stream.FlatMap(func(line string) []string {
		return strings.Split(line, " ")
	})
	stream.Map(func(data string) tuple.KV {
		return tuple.KV{data, 1}
	})

	stream.GroupBy(0).ReStream()

	stream.Map(func(m tuple.KV) []string {
		key := m[0].(string)
		sum := sum(m[1].([]interface{}))
		return []string{key, strconv.Itoa(sum)}
	})

	stream.To(snk)
	<-stream.Open()
}

func sum(slice []interface{}) (count int) {
	for _, v := range slice {
		count += v.(int)
	}
	return
}