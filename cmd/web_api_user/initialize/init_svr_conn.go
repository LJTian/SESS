package initialize

import (
	UserProto "SESS/api/rpc/proto"
	"SESS/cmd/web_api_user/global"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSvrConn() {

	var userSrvHost string
	var userSrvPort int

	data, err := global.GClient.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo))
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	for _, value := range data {
		userSrvHost = value.Address
		userSrvPort = value.Port
		break
	}
	if userSrvHost == "" {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
		return
	}

	zap.S().Infof("userSrvHost is [%s] userSrvPort is [%d], serverName is [%s]",
		userSrvHost, userSrvPort, global.ServerConfig.UserSrvInfo)

	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}
	userSrvClient := UserProto.NewUserGrpcServiceClient(userConn)
	global.UserSrvClient = userSrvClient
}
