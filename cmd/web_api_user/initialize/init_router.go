package initialize

import (
	"SESS/cmd/web_api_user/middlewares"
	"SESS/cmd/web_api_user/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouters(port int) {
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
	//router.InitBaseRouter(ApiGroup)

	go func() {
		Router.Run(fmt.Sprintf(":%d", port))
	}()
}
