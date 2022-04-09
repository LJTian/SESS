package initialize

import (
	"SESS/cmd/grpc_svr_user/global"
	"io/ioutil"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	config2 "SESS/cmd/grpc_svr_user/config"
)

func InitConfig() {

	pathStr := "/Users/ljtian/data/git/github.com/LJTian/SESS/cmd/grpc_svr_user/Nac.yaml"
	config, err := ioutil.ReadFile(pathStr)
	if err != nil {
		return
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
	DbInfo, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "dbconfig.yaml",
		Group:  "dev",
	})
	if err != nil {
		return
	}

	err = yaml.Unmarshal([]byte(DbInfo), &global.DBCfg)
	zap.S().Info(global.DBCfg)

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
