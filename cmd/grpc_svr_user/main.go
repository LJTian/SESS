package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"SESS/cmd/grpc_svr_user/global"
	"SESS/cmd/grpc_svr_user/initialize"
	"SESS/pkg/consul"
	"SESS/pkg/tools"
)

var (
	CfgFilePath = "/Users/ljtian/data/git/github.com/LJTian/SESS/cmd/grpc_svr_user/config_dev.yaml"
)

func main() {

	Port := flag.Int("port", 0, "端口号")
	flag.Parse()

	// 1-加载日志配置
	initialize.InitLoger()
	zap.S().Info("加载日志文件配置成功")

	// 2-链接配置中心
	initialize.InitConfig(CfgFilePath)
	zap.S().Info("链接配置中心成功")

	// 3-初始化数据库
	initialize.InitDB()
	zap.S().Info("初始化数据库成功")

	// 4-注册GRPC服务
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	initialize.InitRegisterGrpcServer(global.CfgInfo.LocalIp, *Port)
	zap.S().Info("启动grpc服务成功")

	// 5-向注册中心进行注册
	serverId := uuid.New().String()
	global.GClient = consul.Connet(global.ServerInfo.RC.IP, global.ServerInfo.RC.Port)
	consul.Register(global.GClient,
		global.CfgInfo.LocalIp,
		*Port,
		global.ServerInfo.Name,
		global.ServerInfo.Tags,
		serverId,
		consul.CheckGrpc,
	)
	zap.S().Info("向注册中心进行注册成功")
	zap.S().Infof("\n 服务IP: [%s]\n port: [%d]\n ServerName: [%s]\n Tags: [%s]\n ServerId: [%s]\n",
		global.CfgInfo.LocalIp, *Port, global.ServerInfo.Name, global.ServerInfo.Tags, serverId)

	// 6-获取信号，优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	consul.UnRegister(global.GClient,
		serverId,
	)
	zap.S().Info("注销成功")
}
