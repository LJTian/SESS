package global

import (
	"SESS/cmd/grpc_svr_user/config"
	"github.com/hashicorp/consul/api"
	"gorm.io/gorm"
)

var (
	ServerInfo config.ServerConfig
	CfgInfo    config.Config
)

var (
	DB      *gorm.DB
	GClient *api.Client // consul 服务中心 链接句柄
)
