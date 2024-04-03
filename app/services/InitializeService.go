package services

import (
	"go-websocket/app/services/bind_center"
	"go-websocket/app/services/grpc_server"
	"go-websocket/app/services/socket"
	"go-websocket/app/services/task"
	"go-websocket/config"
	routers "go-websocket/routes"
	"go-websocket/tools/dbutil"
	"go-websocket/tools/etcds"
	"go-websocket/tools/logger"
	"go-websocket/tools/redisutil"
	"go-websocket/tools/utils"
	"go-websocket/tools/validates"
)

// 初始化配置
func InitializeService() {
	mkdir()                  //启动项目时判断下文件夹
	validates.InitValidate() //初始化参数验证
	config.Init()            //初始化配置文件
	logger.InitLogger()      //初始化日志
	dbutil.InitDbLine()
	redisutil.InitRdLine()
	bind_center.NewService().InitSetServer() //初始化绑定

	if config.GetConf().Etcd.Open {
		//初始化发现服务
		go etcds.NewEtcdDiscovery(map[string]etcds.FunDiscovery{"put": bind_center.EventPut, "del": bind_center.EventDel}).EtcdStartDiscovery([]string{"/prefix1", "/prefix2", "/net", "go-nat-x", etcds.ETCD_SERVER_LIST, etcds.ETCD_PREFIX_ACCOUNT_INFO})
		//初始化注册服务
		go etcds.NewEtcdRegister().EtcdStartRegister(bind_center.RegisterServer)
	}

	routers.InitWsRouters()

	socket.StartTcp() //启动tcp

	if config.GetConf().Grpc.IsOpenRpc {
		go grpc_server.InitGrpcServer() //启动grpc
	}
	task.StartTask() //开启异步定时任务
}

// 启动创建文件夹
func mkdir() {
	utils.CreateDirectoryIfNotExist("config") //初始化一些文件
	utils.CreateDirectoryIfNotExist("log")    //初始化一些文件
}
