package handler

import (
	"context"

	UserProto "SESS/api/rpc/proto"
	"SESS/cmd/web_api_user/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"SESS/cmd/web_api_user/forms"
)

// PwdLogin 密码登录
func PwdLogin(c *gin.Context) {

	pwdLogin := forms.PwdLoginForm{}
	err := c.ShouldBind(&pwdLogin)
	if err != nil {
		return
	}

}

// Register 用户注册
func Register(c *gin.Context) {

	//pwdLogin := forms.PwdLoginForm{}
	//err := c.ShouldBind(&pwdLogin)
	//if err != nil {
	//	zap.S().Error(err)
	//	return
	//}

	response, err := global.UserSrvClient.CreateUser(context.Background(), &UserProto.CreateUserInfo{
		NickName: "ljtian",
		PassWord: "123321",
		Mobile:   "17612133321",
	})
	if err != nil {
		zap.S().Error(err)
		return
	}

	zap.S().Info(response.Id)
}
