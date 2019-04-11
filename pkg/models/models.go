package models

import (
	"fmt"
	"github.com/blog/pkg/setting"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup() {
	db, err := gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))
	if err != nil {
		panic("sql connect error")
	}
	defer db.Close()
}
