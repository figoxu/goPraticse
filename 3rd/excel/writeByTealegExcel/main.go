package main

import (
	"51baibao.com/meishi/ut"
	"bytes"
	"fmt"
	"github.com/icrowley/fake"

	//"github.com/quexer/utee"
	"log"
	//"os"
)

func main() {
	log.Println("hello")

	//fileBuffer
	//file, err := os.OpenFile("file.xlsx", os.O_CREATE|os.O_WRONLY, 0600)
	//utee.Chk(err)
	//writer := ut.NewExcelWriter(file)


	//ioBuffer
	buffer := bytes.NewBuffer(make([]byte, 0))
	writer := ut.NewExcelWriter(buffer)

	recordKey := []string{"编号", "姓名", "语言", "所在地"}
	writer.Write(recordKey)

	for i := 0; i < 10000*10000; i++ {
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
