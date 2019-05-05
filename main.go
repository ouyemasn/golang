package main

import (
	"blog.com/pkg/models"
	"blog.com/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	setting.Setting()
	models.Setup()
}

func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           setting.AppSetting.Port,
		Handler:        router,
		ReadTimeout:    setting.AppSetting.ReadTimeout * time.Second,
		WriteTimeout:   setting.AppSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
