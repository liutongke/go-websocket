package config

import (
	"github.com/BurntSushi/toml"
	"go-websocket/utils"
	"go-websocket/utils/Dir"
	"sync"
	"time"
)

type TomlConfig struct {
	Server     server
	Mysql      MySQL
	Redis      Redis
	CommonConf CommonConf
}

//服务器配置
type server struct {
	Port    string
	RpcPort string
	Mode    int
	Cluster bool
	TcpPort string
}

//redis连接配置
type MySQL struct {
	Addr        string
	TablePrefix string
}

//redis连接配置
type Redis struct {
	MaxIdle     int           //最大闲置连接数量
	MaxActive   int           //最大活动连接数
	IdleTimeout time.Duration //闲置过期时间 在get函数中会有逻辑 删除过期的连接
	Addr        string
	Password    string
	DB          int
}

//公共文件配置
type CommonConf struct {
	IsOpenWebsocket bool   //false关闭true开启
	MgCk            string //敏感词路径
}

var (
	conf *TomlConfig
	once sync.Once
	lock sync.Mutex
)

//获取配置
func GetConfClient() *TomlConfig {
	return conf
}

func Init() *TomlConfig {
	once.Do(func() {
		var filePath string
		if utils.IsDebug() {
			filePath = Dir.GetAbsolutePath("/config/config.toml")
		} else {
			filePath = Dir.GetAbsolutePath("/config/config_line.toml")
		}
		if _, err := toml.DecodeFile(filePath, &conf); err != nil {
			panic(err)
		}
	})
	return conf
}
