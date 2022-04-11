package initialize

import (
	"SESS/cmd/grpc_svr_user/global"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"io/ioutil"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	config2 "SESS/cmd/grpc_svr_user/config"
)

func ConnNac(fileName string) config_client.IConfigClient {
	pathStr := fileName
	config, err := ioutil.ReadFile(pathStr)
	if err != nil {
		zap.S().Panic(err)
	}
	var this config2.NacOS
	err = yaml.Unmarshal(config, &this)
	zap.S().Info(this)

	// 使用配置中心
	sc := []constant.ServerConfig{
		{
			IpAddr: this.IP,
			Port:   uint64(this.Port),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         this.NamespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端的另一种方式 (推荐)
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		zap.S().Panic(err)
	}

	return configClient
}

func getConfigInfo(configClient config_client.IConfigClient, DataId, Group string) string {
	DbInfo, err := configClient.GetConfig(vo.ConfigParam{
		DataId: DataId,
		Group:  Group,
	})
	if err != nil {
		return ""
	}
	return DbInfo
}

func InitConfig(fileName string) {

	// 1-链接Nac
	configClient := ConnNac(fileName)
	// 2-获取数据信息
	if str := getConfigInfo(configClient, "dbconfig.yaml", "dev"); str != "" {
		if err := yaml.Unmarshal([]byte(str), &global.DBCfg); err != nil {
			zap.S().Panic(global.DBCfg)
			return
		}
		zap.S().Info(global.DBCfg)
	}
	// 3-获取服务注册中心信息
	if str := getConfigInfo(configClient, "consul.yaml", "dev"); str != "" {
		if err := yaml.Unmarshal([]byte(str), &global.Consul); err != nil {
			zap.S().Panic(global.Consul)
			return
		}
		zap.S().Info(global.Consul)
	}

	//// 检测配置文件改变
	//configClient.ListenConfig(vo.ConfigParam{
	//	DataId: "dbconfig.yaml",
	//	Group:  "dev",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件发生更改")
	//		fmt.Printf("namespace is %s\n, group is %s\n, dataId is %s\n, data is %s\n",
	//			namespace, group, dataId, data)
	//	},
	//})
}
