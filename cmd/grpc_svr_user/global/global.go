package global

import (
	"SESS/cmd/grpc_svr_user/config"
	"gorm.io/gorm"
)

var (
	DBCfg config.DBInfo
	DB    *gorm.DB
)
