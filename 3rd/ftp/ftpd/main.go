package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Unknwon/goconfig"
	"github.com/goftp/file-driver"
	"github.com/goftp/ftpd/web"
	"github.com/goftp/leveldb-auth"
	"github.com/goftp/leveldb-perm"
	"github.com/goftp/qiniu-driver"
	"github.com/goftp/server"
	"github.com/lunny/log"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	version = "v0.1.1104"
)

var (
	cfg        *goconfig.ConfigFile
	cfgPath    string
	customPath string
)

func main() {
	flag.StringVar(&cfgPath, "config", "config.ini",
		"config file path, default is config.ini and custom.ini")
	flag.Parse()

	if len(cfgPath) <= 0 {
		cfgPath = "config.ini"
		customPath = "custom.ini"
	} else {
		f, _ := filepath.Abs(cfgPath)
		customPath = filepath.Join(filepath.Dir(f), "custom.ini")
	}

	var err error
	cfg, err = goconfig.LoadConfigFile(cfgPath, customPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info("Loaded config files:", cfgPath, customPath)

	port, _ := cfg.Int("server", "port")
	db, err := leveldb.OpenFile("./authperm.db", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	var auth = &ldbauth.LDBAuth{db}
	var perm server.Perm
	if cfg.MustValue("perm", "type") == "leveldb" {
		perm = ldbperm.NewLDBPerm(db, "root", "root", os.ModePerm)
	} else {
		perm = server.NewSimplePerm("root", "root")
	}

	typ, _ := cfg.GetValue("driver", "type")
	var factory server.DriverFactory
	if typ == "file" {
		rootPath, _ := cfg.GetValue("file", "rootpath")
		_, err = os.Lstat(rootPath)
		if os.IsNotExist(err) {
			os.MkdirAll(rootPath, os.ModePerm)
		} else if err != nil {
			fmt.Println(err)
			return
		}
		factory = &filedriver.FileDriverFactory{
			rootPath,
			perm,
		}
	} else if typ == "qiniu" {
		accessKey, _ := cfg.GetValue("qiniu", "accessKey")
		secretKey, _ := cfg.GetValue("qiniu", "secretKey")
		bucket, _ := cfg.GetValue("qiniu", "bucket")
		factory = qiniudriver.NewQiniuDriverFactory(accessKey,
			secretKey, bucket)
	} else {
		log.Fatal("no driver type input")
	}

	// start web manage UI
	useweb, _ := cfg.Bool("web", "enable")
	if useweb {
		web.DB = auth
		web.Perm = perm
		web.Factory = factory
		weblisten, _ := cfg.GetValue("web", "listen")
		admin, _ := cfg.GetValue("admin", "user")
		pass, _ := cfg.GetValue("admin", "pass")
		tls, _ := cfg.Bool("web", "tls")
		certFile, _ := cfg.GetValue("web", "certFile")
		keyFile, _ := cfg.GetValue("web", "keyFile")

		go web.Web(weblisten, "static", "templates", admin, pass, tls, certFile, keyFile)
	}

	ftpName, _ := cfg.GetValue("server", "name")
	opt := &server.ServerOpts{
		Name:    ftpName,
		Factory: factory,
		Port:    port,
		Auth:    auth,
	}

	// start ftp server
	ftpServer := server.NewServer(opt)
	log.Info("FTP Server", version)
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
