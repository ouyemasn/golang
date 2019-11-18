package plugin

import (
	"blog.com/pkg/setting"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
)

var Conf config.Config

func Setup() {
	Conf = config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:       setting.DatabaseSetting.Host,
				Port:       setting.DatabaseSetting.Port,
				User:       setting.DatabaseSetting.User,
				Pwd:        setting.DatabaseSetting.Password,
				Name:       setting.DatabaseSetting.Name,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     config.DriverMysql,
			},
		},
		UrlPrefix: "admin", // 访问网站的前缀
		// Store 必须设置且保证有写权限，否则增加不了新的管理员用户
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.CN,
	}
	fmt.Printf("Fail to read file: %v", Conf)

}
