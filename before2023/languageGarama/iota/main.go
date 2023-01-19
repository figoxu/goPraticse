package main
import "log"

const(
	i = iota
	j
	k
	l
	m
	n
)

func main(){
	log.Println("hello ",i)
	log.Println("hello ",j)
	log.Println("hello ",k)
	log.Println("hello ",l)
	log.Println("hello ",m)
	log.Println("hello ",n)
}
