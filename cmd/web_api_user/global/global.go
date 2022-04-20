package global

import (
	UserProto "SESS/api/rpc/proto"
	ut "github.com/go-playground/universal-translator"
	"github.com/hashicorp/consul/api"

	"SESS/cmd/web_api_user/config"
)

var (
	CfgInfo      config.Config
	ServerConfig config.ServerConfig
)

var (
	Trans   ut.Translator
	GClient *api.Client // consul 服务中心 链接句柄
)

var (
	UserSrvClient UserProto.UserGrpcServiceClient
)
