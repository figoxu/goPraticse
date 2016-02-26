package main
import (
	"log"
	"regexp"
)


func main(){
	REG := regexp.MustCompile(`^[0-9a-zA-Z]{16}$`)
	if ta := "asdfasdf"; REG.MatchString(ta) {
		log.Panic("@val:",ta," mustn't be match")
	}
	if ta := "0123456789abcdef";!REG.MatchString(ta) {
		log.Panic("@val:",ta," must match")
	}
	if ta := "0123456789abcdefgh"; REG.MatchString(ta) {
		log.Panic("@val:",ta," mustn.t match")
	}
	REG2  := regexp.MustCompile(`^[0-9a-zA-Z]{16,}$`)
	if ta := "0123456789abcdefgh";!REG2.MatchString(ta) {
		log.Panic("@val:",ta," must be match")
	}
	if ta := "asdfasdf";REG2.MatchString(ta) {
		log.Panic("@val:",ta," mustn't be match")
	}
	log.Println("all test pass")
}
