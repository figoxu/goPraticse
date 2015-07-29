package main

import (
	"github.com/quexer/utee"
//	"net/url"
	"net/url"
	"fmt"
)

func main() {
	address := "http://sendcloud.sohu.com/webapi/mail.send.json"
	apiUser := "xujianhui_test_4TyfVp"
	apiKey := "3p03OLGHQ1W4cCnv"

	q :=  url.Values{}
	q.Add("api_user", apiUser)
	q.Add("api_key", apiKey)

	q.Add("from", "service@sendcloud.im")
	q.Add("fromname", "SendCloud测试邮件")
	q.Add("to", "xujianhui@mrocker.com")
	q.Add("subject", "FigoTest")
	q.Add("html", "你太 棒了！你已成功的从SendCloud发送了一封测试邮件，接下来快登录前台去完善账户信息吧！")
	q.Add("resp_email_id", "true")



	b ,err := utee.HttpPost(address, q)
	if err!=nil{
		fmt.Print("err:",err)
	}
	fmt.Print("@rsp:"+string(b))
}