package models

import (
	"fmt"
	"log"

	"blog.com/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	gorm.Model
	ID int `gorm:"primary_key" json:"id"`
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name))
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	// 设置空闲连接池中的最大连接数
	db.DB().SetMaxIdleConns(10)

	// 设置到数据库的最大打开连接数
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}
