package main

import (
	"github.com/quexer/utee"
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile("http://xxiongdi.iteye.com", qrcode.Medium, 256, "qr.png")
	utee.Chk(err)
}
