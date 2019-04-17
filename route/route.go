package route

import "github.com/gin-gonic/gin"

func initRoute() {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	api := route.Group("/")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
