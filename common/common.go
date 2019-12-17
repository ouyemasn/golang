package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnJson(c *gin.Context, status int, message string, data interface{}) {
	code := http.StatusOK
	if status == 0 {
		code = http.StatusBadRequest
	}
	c.JSON(code, gin.H{
		"code": code,
		"msg":  message,
		"data": data,
	})
}
