package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"SESS/cmd/web_api_user/global"
	"SESS/cmd/web_api_user/initialize"
	"SESS/pkg/consul"
	"SESS/pkg/tools"
)

var (
	NacFilePath = "/Users/ljtian/data/git/github.com/LJTian/SESS/cmd/web_api_user/config.yaml"
)

func main() {

	// 获取基础信息
	serverId := uuid.New().String()

	// 获取一个没有占用的端口
	Port, _ := tools.GetFreePort()
	Port = 60001

	// 1-加载日志配置
	initialize.InitLoger()

	// 2-链接配置中心
	initialize.InitConfig(NacFilePath)

	// 3-初始化路由
	initialize.InitRouters(Port)

	// 4-服务注册
	zap.S().Infof("服务IP: [%s] port: [%d] ServerName: [%s] ServerId: [%s]",
		global.CfgInfo.LocalIp, Port, global.CfgInfo.DataId, serverId)
	global.GClient = consul.Connet(global.ServerConfig.RC.IP, global.ServerConfig.RC.Port)
	consul.Register(global.GClient,
		global.CfgInfo.LocalIp,
		Port,
		global.CfgInfo.DataId,
		global.ServerConfig.Tags,
		serverId,
		consul.CheckHttp,
	)
	zap.S().Info("向注册中心进行注册成功")

	// 5-初始化svr链接
	initialize.InitSvrConn()

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	consul.UnRegister(global.GClient,
		serverId,
	)
	zap.S().Info("注销成功")
}
