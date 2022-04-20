package initialize

import (
	"io/ioutil"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	cfg "SESS/cmd/web_api_user/config"
	"SESS/cmd/web_api_user/global"
)

func GetConfigByFile(fileName string) cfg.Config {
	pathStr := fileName
	config, err := ioutil.ReadFile(pathStr)
	if err != nil {
		zap.S().Panic(err)
	}
	var this cfg.Config
	err = yaml.Unmarshal(config, &this)
	zap.S().Info(this)
	return this
}

func ConnNac(cfg cfg.Config) config_client.IConfigClient {

	// 使用配置中心
	sc := []constant.ServerConfig{
		{
			IpAddr: cfg.IP,
			Port:   uint64(cfg.Port),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         cfg.NamespaceId, //namespace id
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

	// 1 - 链接Nac
	global.CfgInfo = GetConfigByFile(fileName)
	configClient := ConnNac(global.CfgInfo)
	// 2-获取数据信息
	if str := getConfigInfo(configClient, global.CfgInfo.DataId, global.CfgInfo.Group); str != "" {
		if err := yaml.Unmarshal([]byte(str), &global.ServerConfig); err != nil {
			zap.S().Panic(global.ServerConfig)
			return
		}
		zap.S().Info(global.ServerConfig)
	}
}
