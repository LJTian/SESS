package consul

import "testing"

func TestRegister(t *testing.T) {

	// 1、建立链接
	GClient := Connet("10.211.55.3", 8500)
	// 2、注册服务
	Register(GClient, "10.211.55.3", 8080, "test", []string{"test"}, "test3", "GRPC")
	Register(GClient, "10.211.55.3", 8080, "test", []string{"test"}, "test4", "GRPC")
	// 3、摘掉服务
	//UnRegister("test")
}
