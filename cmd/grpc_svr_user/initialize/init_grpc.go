package initialize

import (
	UserProto "SESS/api/rpc/proto"
	"SESS/cmd/grpc_svr_user/handler"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

func InitRegisterGrpcServer(ip string, port int) {

	zap.S().Infof("Ip: %s port: %d ", ip, port)
	server := grpc.NewServer()
	UserProto.RegisterUserGrpcServiceServer(server, &handler.UserServer{})
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()
}
