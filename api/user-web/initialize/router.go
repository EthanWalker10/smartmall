package initialize

import (
	"net/http"

	"github.com/EthanWalker10/smartmall/api/user-web/middlewares"
	"github.com/EthanWalker10/smartmall/api/user-web/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	//配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return Router
}
