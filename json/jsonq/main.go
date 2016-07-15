package main

import (
	"encoding/json"
	"fmt"
	"github.com/figoxu/utee"
	"github.com/jmoiron/jsonq"
	"log"
	"strings"
)

const (
	jsonstring = `{
		    "foo": 1,
		    "bar": 2,
		    "test": "Hello, world!",
		    "baz": 123.1,
		    "array": [
			{"foo": 1},
			{"bar": 2},
			{"baz": 3}
		    ],
		    "subobj": {
			"foo": 1,
			"subarray": [1,2,3],
			"subsubobj": {
			    "bar": 2,
			    "baz": 3,
			    "array": ["hello", "world"]
			}
		    },
		    "bool": true
		}`

	msgstring = `
  [{
    "device": {
      "hardwareConfig": {
        "manufacture": "Huawei",
        "brand": "google",
        "model": "Nexus 6P"
      },
      "deviceId": {
        "tid": "3513034472c990179016cba46a6c8b83e",
        "imeis": [
          "867981020673938"
        ],
        "wifiMacs": [
          "02:00:00:00:00:00"
        ],
        "androidId": "efdcf5074552a62f",
        "serialNo": "ENU7N15A28003794"
      }
    },
    "app": {
      "name": "Demo",
      "globalId": "com.tendcloud.demo",
      "versionName": "1.0",
      "versionCode": 1,
      "installTime": 1459922545801,
      "updateTime": 1459924397419,
      "cert": "3082030d308201f5a00302010202040f6bc62b300d06092a864886f70d01010b05003037310b30090603550406130255533110300e060355040a1307416e64726f6964311630140603550403130d416e64726f6964204465627567301e170d3136303232363031353835355a170d3436303231383031353835355a3037310b30090603550406130255533110300e060355040a1307416e64726f6964311630140603550403130d416e64726f696420446562756730820122300d06092a864886f70d01010105000382010f003082010a0282010100b620bde90feb8394b5c6dda6306e4b8331fd52ebe3ab9441789d4271609c0e9ac787934ef45ed9e21b0a3ba3ff7cfbbc8481072fb61df6667fa6f68367f8f21e10468d3347a37f89403a59299a776c3c0e33092ffea3412b5d3d7c4fd96749eeda72e692fc503ee7b692cced5c5c125b56883c720beef4c8cb18a3cb9db875a28b6a3ea2de864c717d37f98b34af3006dca7794ff017c86d734dfc87fd2e4f772c124ddae3e6a4719a945ff6491a9c71315c9e0f497ebcea429528d12ae496c822b609bea2a6f186b961663807686be538c105901b5b7e44766e5be32a0aa77ef18eeaee20c96d55788a5e187c212735adf5c2dd26baa974ff267441de5919710203010001a321301f301d0603551d0e04160414b8ee6820071e8a2e5fa2eb59a81e0a2ac280aacd300d06092a864886f70d01010b050003820101007a85dd0456d3cbdd150512b2ac86cc13185e1ad96efe958024badf83e13bbf2a6d87e0ce096df5fdc21bda2ae72d34213ad753db1ef7f9111dcca9c531c7f891c5827eb0943523c2845520655856404e4f0dcdac31a87124d63011804015f1040ede580e613b34070a9baedd33327ab15eabeb5826f5d17783b65d53f8b226162a452ffbe0739dc92f73e7ee7c69ff7ec5caed0d380d8c018cb13c8da3754693f7f3f0eaa248b26feadeece879a43044283eecd8bf4b61ee52686efafd55276a12e101b25ff50dbb4f83ff574ff42cd40478cf42e5b363482883429fdfa24216225efb92c29928d55726846f91908cee40e6e8a0f8ca43e867de5690eb80970d",
      "appKey": [
        {
          "appKey": "2ABD1A108697BD992CFD1AF115FC669B",
          "service": "app"
        }
      ]
    },
    "sdk": {
      "features": [

      ],
      "minorVersion": 0,
      "build": 0,
      "platform": "Android",
      "version": 1.2
    },
    "appContext": {
      "account": {
        "accountId": "-1"
      },
      "push": [{"3rdAppId":"2ABD1A108697BD992CFD1AF115FC669B","channel":"nick","dt":"3513034472c990179016cba46a6c8b83e"}]
    },
    "ts": 1459924476160,
    "action": {
      "domain": "push",
      "name": "message",
      "data": {
        "action": 1,
        "msgSign": "1214"
      }
    }
  },{"test":"cool"}]`
)

func main() {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(jsonstring))
	dec.Decode(&data)
	jq := jsonq.NewQuery(data)

	// data["foo"] -> 1
	i, err := jq.Int("foo")
	utee.Chk(err)
	log.Println("data['foo'] -> 1 : ", i)
	i, err = jq.Int("subobj", "subarray", "1")
	utee.Chk(err)
	log.Println("data['subobj']['subarray'][1] -> 2: ", i)
	s, err := jq.String("subobj", "subsubobj", "array", "0")
	utee.Chk(err)
	log.Println("data['subobj']['subarray']['array'][0] -> 'hello' ", s)

	// data["subobj"] -> map[string]interface{}{"subobj": ...}
	obj, err := jq.Object("subobj")
	utee.Chk(err)
	log.Println("jq.Object('subobj') : ", obj)

	type item map[string]interface{}

	log.Println("====================>")
	jsa := wrapAsArray(msgstring)
	log.Println("@jsa:   ", jsa)
	data2 := map[string]interface{}{}
	dec = json.NewDecoder(strings.NewReader(jsa))
	dec.Decode(&data2)
	jq = jsonq.NewQuery(data2)

	action, err := jq.Int("items", "0", "action", "data", "action")
	utee.Chk(err)
	log.Println("@action:", action)
	msgSign, err := jq.String("items", "0", "action", "data", "msgSign")
	utee.Chk(err)
	log.Println("@msgSign:", msgSign)

	arr, err := jq.Array("items")
	utee.Chk(err)
	log.Println("len of arr is :", len(arr))
}

func wrapAsArray(jsonTxt string) string {
	v := fmt.Sprint("{ \"items\":", jsonTxt, "}")
	return v
}
