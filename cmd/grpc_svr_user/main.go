package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"SESS/cmd/grpc_svr_user/global"
	"SESS/cmd/grpc_svr_user/initialize"
	"SESS/pkg/consul"
	"SESS/pkg/tools"
)

var (
	ServerName  = "grpc_svr_user"
	NacFilePath = "/Users/ljtian/data/git/github.com/LJTian/SESS/cmd/grpc_svr_user/Nac.yaml"
)

func main() {

	// 通过命令行获取参数信息
	IP := flag.String("ip", "192.168.124.5", "ip地址")
	Port := flag.Int("port", 0, "端口号")
	flag.Parse()

	// 1-加载日志配置
	initialize.InitLoger()
	zap.S().Info("加载日志文件配置成功")

	// 2-链接配置中心
	initialize.InitConfig(NacFilePath)
	zap.S().Info("链接配置中心成功")

	// 3-初始化数据库
	initialize.InitDB()
	zap.S().Info("初始化数据库成功")

	// 4-注册GRPC服务
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	initialize.InitRegisterGrpcServer(*IP, *Port)
	zap.S().Info("启动grpc服务成功")

	// 5-向注册中心进行注册
	global.GClient = consul.Connet(global.Consul.IP, global.Consul.Port)
	consul.Register(global.GClient,
		"192.168.124.5",
		*Port,
		ServerName,
		[]string{""},
		ServerName,
	)
	zap.S().Info("向注册中心进行注册成功")

	// 6-获取信号，优雅退出
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	consul.UnRegister(global.GClient,
		ServerName,
	)
	zap.S().Info("注销成功")
}
