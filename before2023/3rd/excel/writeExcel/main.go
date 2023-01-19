package main

import (
	"encoding/csv"
	"fmt"
	"github.com/icrowley/fake"
	"os"
)

func main() {
	f, err := os.Create("./test.xls")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	w := csv.NewWriter(f)
	w.Write([]string{"编号", "姓名", "语言", "所在地"})

	for i := 0; i < 100*10000; i++ {
		w.Write([]string{
			fmt.Sprint(i),
			fake.FullName(),
			fake.Language(),
			fake.City(),
		})
		if i%10000 == 0 {
			fmt.Println("Handle @i:", i)
			fmt.Println("Flush @i:",i)
			w.Flush()
		}
	}
	w.Flush()
}
