package config

import "time"

type TomlConfig struct {
	Server     server
	Mysql      mysql
	Redis      redis
	Etcd       etcd
	Grpc       grpc
	CommonConf CommonConf
	WebSocket  WebSocket
}

// 服务器配置
type server struct {
	Port    string
	Mode    int
	Cluster bool
	TcpPort string
}

// redis连接配置
type redis struct {
	MaxIdle     int           //最大闲置连接数量
	MaxActive   int           //最大活动连接数
	IdleTimeout time.Duration //闲置过期时间 在get函数中会有逻辑 删除过期的连接
	Addr        string
	Password    string
	DB          int
}
type mysql struct {
	Addr               string
	SetMaxIdleConn     int
	SetMaxOpenConn     int
	SetConnMaxLifetime time.Duration
	LogLevel           int
	DateFormat         int
	LogFolder          string
	Cmd                bool
	SlowThreshold      int
}

type etcd struct {
	Addr                 string
	LogLevel             int
	LogFolder            string
	Username             string
	Password             string
	MaxCallSendMsgSize   int
	MaxCallRecvMsgSize   int
	AutoSyncInterval     time.Duration
	DialTimeout          time.Duration
	DialKeepAliveTime    time.Duration
	DialKeepAliveTimeout time.Duration
	OutputPaths          string
	ErrorOutputPaths     string
}

type grpc struct {
	RpcPort   string
	IsOpenRpc bool
	Tsl       bool
}

// 公共文件配置
type CommonConf struct {
	IsOpenWebsocket bool   //false关闭true开启
	MgCk            string //敏感词路径
}

type WebSocket struct {
	CleanConnection         bool //true
	HeartbeatExpirationTime uint64
}
