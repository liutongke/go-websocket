package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"go-websocket/tools/fileutil"
	"go-websocket/tools/utils"
	"sync"
)

var (
	conf *TomlConfig
	once sync.Once
	lock sync.Mutex
)

// 获取配置
func GetConf() *TomlConfig {
	return conf
}

func Init() *TomlConfig {
	once.Do(loadConfig)
	return conf
}

func loadConfig() {
	var configFilePath string

	// 根据环境选择配置文件路径
	switch {
	case utils.IsRsapi():
		configFilePath = "./config/config_raspi.toml"
	case utils.IsDebug():
		configFilePath = "./config/config.toml"
	default:
		configFilePath = "./config/config_line.toml"
	}

	// 打印加载的配置文件信息
	printLoadedConfig(configFilePath)

	// 解析配置文件
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		panic(err)
	}
}

func printLoadedConfig(filePath string) {
	// 设置终端输出颜色为绿色
	fmt.Printf("\033[1;32;40m%s\033[0m\n", "加载配置文件："+filePath)
}

func InitTest() *TomlConfig {
	once.Do(func() {
		var filePath string
		if utils.IsDebug() {
			filePath = fileutil.GetAbsDirPath("../../config/config.toml")
		} else {
			filePath = fileutil.GetAbsDirPath("../../config/config_line.toml")
		}

		if _, err := toml.DecodeFile(filePath, &conf); err != nil {
			panic(err)
		}
	})
	return conf
}
