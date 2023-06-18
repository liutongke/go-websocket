package config

import (
	"github.com/BurntSushi/toml"
	"go-websocket/tools/Dir"
	"go-websocket/tools/Tools"
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
	once.Do(func() {
		var filePath string
		if Tools.IsDebug() {
			filePath = Dir.GetAbsDirPath("./config/config.toml")
		} else {
			filePath = Dir.GetAbsDirPath("./config/config_line.toml")
		}
		if _, err := toml.DecodeFile(filePath, &conf); err != nil {
			panic(err)
		}
	})
	return conf
}
