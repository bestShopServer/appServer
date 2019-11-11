package setting

import (
	"github.com/go-ini/ini"
	"os"
	"shop/pkg/logging"
	"time"
)

type app struct {
	PageSize int
	Name     string
}

//本服务配置
type server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

//数据库配置
type database struct {
	DbType string
	DbUser string
	DbPass string
	DbHost string
	DbPort int
	DbName string
}

var (
	AppCfg    = &app{}
	ServerCfg = &server{}
	Dbcfg     = &database{}
	cfg       *ini.File
)

// Setup initialize the configuration instance
func Setup() {
	var err error
	if RunMode := os.Getenv("RunMode"); RunMode == "release" {
		if cfg, err = ini.Load("conf/app_release.ini"); err != nil {
			logging.Fatal(err)
		}
	} else {
		if cfg, err = ini.Load("conf/app_debug.ini"); err != nil {
			logging.Fatal(err)
		}
	}

	mapTo("app", AppCfg)
	mapTo("server", ServerCfg)
	mapTo("database", Dbcfg)

	ServerCfg.ReadTimeout = ServerCfg.ReadTimeout * time.Second
	ServerCfg.WriteTimeout = ServerCfg.WriteTimeout * time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		logging.Fatal(err)
	}
}
