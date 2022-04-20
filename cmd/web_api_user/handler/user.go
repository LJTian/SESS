package handler

import (
	"SESS/cmd/web_api_user/middlewares"
	"SESS/cmd/web_api_user/models"
	"context"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"

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

	register := forms.RegisterForm{}
	err := c.ShouldBind(&register)
	if err != nil {
		zap.S().Error(err)
		return
	}

	response, err := global.UserSrvClient.CreateUser(context.Background(), &UserProto.CreateUserInfo{
		NickName: register.Mobile,
		PassWord: register.PassWord,
		Mobile:   register.Mobile,
	})
	if err != nil {
		zap.S().Error(err)
		return
	}
	zap.S().Info(response.Id)
	zap.S().Info("用户注册成功")

	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(response.Id),
		NickName:    response.NickName,
		AuthorityId: response.Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "imooc",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         response.Id,
		"nick_name":  response.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})

}
