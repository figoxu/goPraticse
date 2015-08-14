package main
import (
	"time"
	"fmt"
)


func main(){
	dateFormat := "2006-01-02 15:04:05"
//	当前时间戳
	fmt.Println(time.Now().Unix())
//	当前格式化时间
	fmt.Println(time.Now().Format(dateFormat))
//	时间戳转str格式化时间
	str_time := time.Unix(1389058332, 0).Format(dateFormat)
	fmt.Println(str_time)

//	str格式化时间转时间戳--way 01
	the_time := time.Date(2014, 1, 7, 5, 50, 4, 0, time.Local)
	unix_time := the_time.Unix()
	fmt.Println(unix_time)

//	str格式化时间转时间戳--way 02
	the_time, err := time.Parse(dateFormat, "2014-01-08 09:04:41")
	if err == nil {
		unix_time := the_time.Unix()
		fmt.Println(unix_time)
	}

//	日期换算
	baseTime := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
	date := baseTime.Add(1722*7*24*time.Hour + 24*time.Hour + 66355*time.Second)
	fmt.Println(date)
}
