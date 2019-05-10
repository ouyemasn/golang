package route

import (
	"blog.com/route/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	apiv1 := route.Group("/")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}
	return route
}
