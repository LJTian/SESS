package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

var (
	CheckHttp = "HTTP"
	CheckGrpc = "GRPC"
)

// 建立链接
func Connet(addr string, port int) *api.Client {

	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", addr, port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return client
}

// 注册
func Register(GClient *api.Client,
	address string, port int, name string, tags []string, id string, checkType string) error {

	var check *api.AgentServiceCheck

	if checkType == CheckHttp {
		//生成对应的检查对象
		check = &api.AgentServiceCheck{
			HTTP:                           fmt.Sprintf("http://%s:%d/health", address, port),
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "150s",
		}
		zap.S().Infof("健康检查地址:[%s]", check.HTTP)
	} else {
		//生成对应的检查对象
		check = &api.AgentServiceCheck{
			GRPC:                           fmt.Sprintf("%s:%d", address, port),
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "150s",
		}
		zap.S().Infof("健康检查地址:[%s]", check.GRPC)
	}

	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err := GClient.Agent().ServiceRegister(registration)
	//client.Agent().ServiceDeregister()
	if err != nil {
		panic(err)
	}
	return nil
}

// 注销服务
func UnRegister(GClient *api.Client, serverId string) {
	err := GClient.Agent().ServiceDeregister(serverId)
	if err != nil {
		panic(err)
	}
}
