package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
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
func Register(GClient *api.Client, address string, port int, name string, tags []string, id string) error {

	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", address, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
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

func main() {

	// 1、建立链接
	GClient := Connet("10.211.55.3", 8500)
	// 2、注册服务
	Register(GClient, "10.211.55.3", 8080, "test", []string{"test"}, "test3")
	Register(GClient, "10.211.55.3", 8080, "test", []string{"test"}, "test4")
	// 3、摘掉服务
	//UnRegister("test")
}
