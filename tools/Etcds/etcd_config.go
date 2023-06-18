package Etcds

import (
	"github.com/coreos/etcd/clientv3"
	"go-websocket/config"
	"go-websocket/tools/Dir"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// Endpoints：etcd 服务的地址列表，以 URL 形式表示。
// AutoSyncInterval：更新端点的最新成员的间隔时间。设置为 0 表示禁用自动同步。
// DialTimeout：建立连接的超时时间。
// DialKeepAliveTime：客户端向服务器发送保活探测的时间间隔。
// DialKeepAliveTimeout：客户端等待保活探测的响应的超时时间。如果在此时间内未收到响应，则关闭连接。
// MaxCallSendMsgSize：客户端发送请求的大小限制（以字节为单位）。如果设置为 0，则默认为 2.0 MiB。
// MaxCallRecvMsgSize：客户端接收响应的大小限制（以字节为单位）。如果设置为 0，则默认为 "math.MaxInt32"，因为范围响应可能超过请求发送限制。
// TLS：客户端的安全凭证，如果有的话，使用 tls.Config 类型。
// Username：用于身份验证的用户名。
// Password：用于身份验证的密码。
// RejectOldCluster：当设置为 true 时，拒绝与过时的集群创建客户端。
// DialOptions：用于 gRPC 客户端的拨号选项列表，可以添加拦截器等。
// LogConfig：配置客户端端的日志记录器。如果为 nil，则使用默认日志记录器。
// Context：默认的客户端上下文，可用于取消 gRPC 的拨号和其他操作。
// PermitWithoutStream：当设置为 true 时，允许客户端在没有任何活动流（RPC）的情况下向服务器发送保活探测。
func EtcdConfig() clientv3.Config {
	etcdConfig := config.GetConf().Etcd
	
	return clientv3.Config{
		Endpoints:            []string{etcdConfig.Addr},
		AutoSyncInterval:     etcdConfig.AutoSyncInterval * time.Second,
		DialTimeout:          etcdConfig.DialTimeout * time.Second,
		DialKeepAliveTime:    etcdConfig.DialKeepAliveTime * time.Second,
		DialKeepAliveTimeout: etcdConfig.DialKeepAliveTimeout * time.Second,
		MaxCallSendMsgSize:   etcdConfig.MaxCallSendMsgSize * 1024 * 1024,
		MaxCallRecvMsgSize:   etcdConfig.MaxCallRecvMsgSize * 1024 * 1024,
		//TLS:                  nil,
		Username: etcdConfig.Username,
		Password: etcdConfig.Password,
		//RejectOldCluster: false,
		//DialOptions:      nil,
		LogConfig: getLogConfig(),
		//Context:              nil,
		//PermitWithoutStream:  false,
	}
}

func getLogConfig() *zap.Config {
	return &zap.Config{
		Level:            zap.NewAtomicLevelAt(getLogLevel()),
		Encoding:         "json",
		OutputPaths:      []string{Dir.GetAbsDirPath(config.GetConf().Etcd.OutputPaths)},
		ErrorOutputPaths: []string{Dir.GetAbsDirPath(config.GetConf().Etcd.ErrorOutputPaths)},
		Development:      false,
	}
}

// #0 DebugLevel: 调试级别日志，通常在生产环境中被禁用。
// #1 InfoLevel: 默认日志级别，用于普通信息的日志记录。
// #2 WarnLevel: 警告级别日志，比普通信息更重要，但不需要人工逐个审查。
// #3 ErrorLevel: 错误级别日志，高优先级。如果应用程序正常运行，不应生成任何错误级别的日志。
// #4 DPanicLevel: 重要错误级别日志，用于开发环境。在写入消息后，记录器会触发 panic。
// #5 PanicLevel: 日志一条消息，然后触发 panic。
// #6 FatalLevel: 日志一条消息，然后调用 os.Exit(1) 终止程序。
func getLogLevel() zapcore.Level {
	switch config.GetConf().Etcd.LogLevel {
	case 0:
		return zapcore.DebugLevel
	case 1:
		return zapcore.InfoLevel
	case 2:
		return zapcore.WarnLevel
	case 3:
		return zapcore.ErrorLevel
	case 4:
		return zapcore.DPanicLevel
	case 5:
		return zapcore.PanicLevel
	case 6:
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}
