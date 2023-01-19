package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/tok"
	"github.com/quexer/utee"
	"log"
	"net/http"
	"startrip.vip/sdz-solution/sdz-connector/netware/arm/tokV2"
	"time"
	"github.com/figoxu/Figo"
	"github.com/go-martini/martini"
)

var (
	g_hub *tok.Hub
)

func main() {
	log.Println("Hello")
	port := ":16666"
	hub, h_ws := initActor(port)
	g_hub = hub

	m := Figo.NewMartini(1, 10, "")
	m.Use(martini.Static("./dist", martini.StaticOptions{SkipLogging: true}))
	http.Handle("/ws", h_ws)
	utee.Chk(http.ListenAndServe(port, nil))
}

func initActor(port string) (*tok.Hub, http.Handler) {
	initTokCfg := func() {
		cfg, err := config.NewConfig("ini", "conf.ini")
		timeoutRead, err := cfg.Int("tok::timeout_read")
		utee.Chk(err)
		timeoutWrite, err := cfg.Int("tok::timeout_write")
		utee.Chk(err)
		timeoutAuth, err := cfg.Int("tok::timeout_auth")
		utee.Chk(err)
		timeoutPing, err := cfg.Int("tok::interval_ping")
		tokV2.READ_TIMEOUT = time.Second * time.Duration(timeoutRead)
		tokV2.WRITE_TIMEOUT = time.Second * time.Duration(timeoutWrite)
		tokV2.AUTH_TIMEOUT = time.Second * time.Duration(timeoutAuth)
		tokV2.SERVER_PING_INTERVAL = time.Second * time.Duration(timeoutPing)
		tokV2.SERVER_PING_FLAG, err = cfg.Bool("tok::server_ping")
	}
	initTokCfg()
	hubCfg := &tok.HubConfig{
		Actor: &WsActor{},
		Sso:   true,
	}

	return tok.CreateWsHandler(nil, hubCfg, true, func(*http.Request) (*tok.Device, error) {
		dv := tok.CreateDevice("figo", "xu")
		return dv, nil
	})
}
