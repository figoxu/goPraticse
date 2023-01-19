package main

import (
	"github.com/figoxu/Figo"
	"log"
	"github.com/quexer/utee"
	"encoding/hex"
	"encoding/base64"
)

// 可通过openssl产生
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey4Test = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDTNQybk1sCt4yFb+l8HQX9nBd3QkCUTenx+zzsyQycuJvGmX13
/b/03La8OvxVm9a1WZTk/lnhA0erVedLzC1Lp7hIfnfZhHRoEZzj9afpMa8B47k5
5Lh5s4GgRwx0zw0IEKFzxwN8O3IJTpeeeZgABoWcfJhVSiujzxbEpw2EaQIDAQAB
AoGBAK9/4EcSLcjXLive71uTXlv7LUCKy9Cv4VqSknCLKzC68a4X8rsXHj9ge3Nz
bCPSx5mPo3qYo6SmrhH/4p8IPQWi+aBJ/y1GwQuR/iVzlCxJt8rJVyUl3Rji4UIb
sYX/0Mwu+NxMQ43bvv9YT/8zifpd3hEdGCg9CuEN45TPJF+1AkEA7q+3uy/nIknB
mTJNkZV438XAtWH5jNK+eYE1jnkgntpX7sQU21oUzQ0Wbr/aEprEu9KuAHH1roaJ
Ae6IwATR7wJBAOKHDr2m3M+kFqwxq7+nM36rKEUgbmJtJypT9Set/UtQZxvUGDLi
+d+/yiJqivHXKxygpk93TbcX4MrtCUisBycCQFIdzBUvRtKqE1v0TXF/viUmcMU2
XtePDY7Z4CYTECD2t3fip9ZLaIqfLQ+PG6R48KQ5uDlY+5A+otYyTYPaZKsCQFcM
Tz3RwUiJZa0F6Vncho1GeFMYA1MPXt2FJc/5rDwkyXqIJkRntF2m9aYECyCj7o0x
rrcawWJ6aoeQTuD+OkECQQDCxqWrCqNvEEaTGclfqAzK1Ba0mEbqfpg1Be6oDFRh
e06BZuCvT4sGR14Wpygioe4ZXkKPNIf096MvO8Tv4ikB
-----END RSA PRIVATE KEY-----
`

//openssl
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey4Test = `
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDTNQybk1sCt4yFb+l8HQX9nBd3
QkCUTenx+zzsyQycuJvGmX13/b/03La8OvxVm9a1WZTk/lnhA0erVedLzC1Lp7hI
fnfZhHRoEZzj9afpMa8B47k55Lh5s4GgRwx0zw0IEKFzxwN8O3IJTpeeeZgABoWc
fJhVSiujzxbEpw2EaQIDAQAB
-----END PUBLIC KEY-----
`


func main() {
	rsaHelp := Figo.NewRsaHelp(publicKey4Test, privateKey4Test)

	bs, err := rsaHelp.PriEnc([]byte("hello figo"))
	utee.Chk(err)
	result, err := rsaHelp.PubDec(bs)
	utee.Chk(err)
	log.Println(string(result))

	bs, err = rsaHelp.PubEnc([]byte("nice world"))
	utee.Chk(err)
	result, err = rsaHelp.PriDec(bs)
	utee.Chk(err)
	log.Println(string(result))

	bs,err=hex.DecodeString("51255C409348D84F403E55B148A7D9BF046E22A533849D6FD9A859B2057D4D6E33F16A8ECD0B2336EC6148C957993175A01D7817B5AB46CF9C56933EC1B33C5F072A312580181DC9A65D2518638458553EBEC26B5BD7953AFDCE8E36E6441A3FF97E9150AA11F0AE97DE6E7D3BFEA0FB6103ED735461E501E70BDF58BE2904FB")
	utee.Chk(err)
	bs,err=rsaHelp.PriDec(bs)
	utee.Chk(err)
	log.Println(string(bs))

	bs,err=base64.StdEncoding.DecodeString(`lC+HpruUsfDiGNFJMGzrmjZqFJoS7+/6DjmNijdBBdLKoL0EMRbRuV/vpLsDGBiTMtfZplHpehrnkwieT7Ud4HRwYY9q2K5+FUxkCROtJyesJfvLZP6/xszCCoQy/fT6MOOrhHVy3xky5Dz80IG8gPP8H6AziOl11+Scxc/d8/g=`)
	utee.Chk(err)
	bs,err=rsaHelp.PriDec(bs)
	utee.Chk(err)
	log.Println(string(bs))

}