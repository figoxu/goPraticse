package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Unknwon/goconfig"
	"github.com/lunny/log"
	"github.com/syndtr/goleveldb/leveldb"
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

	var auth = &LDBAuth{db}
	var perm Perm
	if cfg.MustValue("perm", "type") == "leveldb" {
		perm = NewLDBPerm(db, "root", "root", os.ModePerm)
	} else {
		perm = NewSimplePerm("root", "root")
	}

	rootPath, _ := cfg.GetValue("file", "rootpath")
	_, err = os.Lstat(rootPath)
	if os.IsNotExist(err) {
		os.MkdirAll(rootPath, os.ModePerm)
	} else if err != nil {
		fmt.Println(err)
		return
	}
	factory := &FileDriverFactory{
		rootPath,
		perm,
	}


	ftpName, _ := cfg.GetValue("server", "name")
	opt := &ServerOpts{
		Name:    ftpName,
		Factory: factory,
		Port:    port,
		Auth:    auth,
	}

	// start ftp server
	ftpServer := NewServer(opt)
	err = ftpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
