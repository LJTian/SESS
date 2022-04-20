package router

import (
	api "SESS/api/http/http_user_api"
	"SESS/cmd/web_api_user/handler"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userGroup := Router.Group("user")
	{
		//userGroup.GET("", middlewares.JWTAuth())
		userGroup.POST(api.PostPwdLogin, handler.PwdLogin)
		userGroup.POST(api.PostRegister, handler.Register)
	}
}
