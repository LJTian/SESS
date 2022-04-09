package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	UserProto "SESS/api/rpc/proto"
	"SESS/cmd/grpc_svr_user/handler"
	"SESS/cmd/grpc_svr_user/initialize"
	"SESS/pkg/tools"
)

func main() {

	// 通过命令行获取参数信息
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口号")

	// 加载配置信息
	// 1-加载日志配置
	initialize.InitLoger()
	zap.S().Info("加载日志文件配置成功")
	// 2-链接配置中心
	initialize.InitConfig()
	zap.S().Info("链接配置中心成功")
	// 3-初始化数据库
	initialize.InitDB()
	zap.S().Info("初始化数据库成功")

	// 4-注册GRPC服务
	flag.Parse()
	zap.S().Info("ip: ", *IP)
	if *Port == 0 {
		*Port, _ = tools.GetFreePort()
	}
	zap.S().Info("port: ", *Port)
	server := grpc.NewServer()
	UserProto.RegisterUserGrpcServiceServer(server, &handler.UserServer{})
	_, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	zap.S().Info("启动grpc服务成功")

	// 5-向注册中心进行注册

	// 6-获取信号，优雅退出
	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//if err = client.Agent().ServiceDeregister(serviceID); err != nil {
	//	zap.S().Info("注销失败")
	//}
	zap.S().Info("注销成功")

}
