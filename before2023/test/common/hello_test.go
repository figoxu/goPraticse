package common
import (
	"testing"
	"fmt"
)
//go test -run Hello
func TestHello(t *testing.T) {
	hello()
	fmt.Println("hello world 1")
}
//go test -run World
func TestWorld(t *testing.T) {
	fmt.Println("hello world 2")
}


//go test -test.bench Hello
func BenchmarkHello(b *testing.B) {
	fmt.Println("Benchmark Hello")
	for n := 0; n < b.N; n++ {
		hello()
	}
}

//go test -test.bench World
func BenchmarkWorld(b *testing.B) {
	fmt.Println("Benchmark World")
	for n := 0; n < b.N; n++ {
		hello()
		fmt.Println("world")
	}
}


//mix test
//go test -test.bench Hello -test.run World