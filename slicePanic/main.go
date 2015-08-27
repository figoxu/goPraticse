package main
import "fmt"



func main(){
	fmt.Println("start")
	d,err := testInfo()
	if err != nil {
		fmt.Println("@error:",err)
	}else{
		fmt.Println("@data:",d)
	}
	fmt.Println("end")
}


func testInfo()([]interface{},error){
	data := []interface{}{}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err ," (recover)")    //这里的err其实就是panic传入的内容，55
		}
	}()
	fmt.Println("hello")
	data = append(data, "text1")
	data = append(data, "text2")
	data = append(data, "text3")
	data = append(data, "text4")
	data = append(data, "text5")
	d:=data[:1000]
	fmt.Println(d)
	return data,nil
}


