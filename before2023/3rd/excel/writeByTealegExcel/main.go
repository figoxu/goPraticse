package main

import (
	"51baibao.com/meishi/ut"
	"fmt"
	"github.com/icrowley/fake"
	"github.com/quexer/utee"
	"os"

	//"github.com/quexer/utee"
	"log"
	//"os"
)

func main() {
	log.Println("hello")
	file, err := os.OpenFile("file.xlsx", os.O_CREATE|os.O_WRONLY, 0600)
	utee.Chk(err)
	writer := ut.NewExcelWriter(file)

	recordKey := []string{"编号", "姓名", "语言", "所在地"}
	writer.Write(recordKey)

	for i := 0; i < 100*10000; i++ {
		writer.Write([]string{
			fmt.Sprint(i),
			fake.FullName(),
			fake.Language(),
			fake.City(),
		})
		if i%10000 == 0 {
			fmt.Println("Handle @i:", i)
			fmt.Println("Flush @i:",i)
			writer.Flush()
		}
	}

	log.Println("done")
	writer.Flush()
	log.Println("finish")

}
