package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

type DatabaseConf struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &DatabaseConf{}

type AppConf struct {
	Env          string
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

var AppSetting = &AppConf{}
var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("conf/main.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	mapTo("database", DatabaseSetting)
	mapTo("app", AppSetting)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo  err: %v", err)
	}
}
