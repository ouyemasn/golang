package main

import (
	"blog.com/pkg/models"
	"blog.com/pkg/plugin"
	"blog.com/pkg/setting"
	"blog.com/route"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin" // 引入适配器，必须引入，如若不引入，则需要自己定义
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/themes/adminlte" // 必须引入，不然报错

	"net/http"
	"time"
)

func init() {
	setting.Setting()
	models.Setup()
	plugin.Setup()
}

func main() {
	router := route.InitRoute()
	// 实例化一个GoAdmin引擎对象
	eng := engine.Default()
	// GoAdmin全局配置，也可以写成一个json，通过json引入

	// 这里引入你需要管理的业务表配置
	// 关于Generators，详见 https://github.com/GoAdminGroup/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// 增加配置与插件，使用Use方法挂载到Web框架中

	_ = eng.AddConfig(plugin.Conf).AddPlugins(adminPlugin).Use(router)
	s := &http.Server{
		Addr:           setting.AppSetting.Port,
		Handler:        router,
		ReadTimeout:    setting.AppSetting.ReadTimeout * time.Second,
		WriteTimeout:   setting.AppSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
