package services

import (
	"go-websocket/app/services/bind_center"
	"go-websocket/app/services/grpc_server"
	"go-websocket/app/services/socket"
	"go-websocket/app/services/task"
	"go-websocket/config"
	routers "go-websocket/routes"
	"go-websocket/tools/DbLine"
	"go-websocket/tools/Dir"
	"go-websocket/tools/Etcds"
	"go-websocket/tools/Logger"
	"go-websocket/tools/RdLine"
	"go-websocket/tools/validates"
	"log"
)

// 初始化配置
func InitializeService() {
	validates.InitValidate() //初始化参数验证
	config.Init()            //初始化配置文件
	Logger.InitLogger()      //初始化日志
	DbLine.InitDbLine()
	RdLine.InitRdLine()
	bind_center.NewService().InitSetServer() //初始化绑定
	
	if config.GetConf().Etcd.Open {
		//初始化发现服务
		go Etcds.NewEtcdDiscovery(map[string]Etcds.FunDiscovery{"put": bind_center.EventPut, "del": bind_center.EventDel}).EtcdStartDiscovery([]string{"/prefix1", "/prefix2", "/net", "go-nat-x", Etcds.ETCD_SERVER_LIST, Etcds.ETCD_PREFIX_ACCOUNT_INFO})
		//初始化注册服务
		go Etcds.NewEtcdRegister().EtcdStartRegister(bind_center.RegisterServer)
	}
	mkdir() //启动项目时判断下文件夹

	routers.InitWsRouters()

	socket.StartTcp() //启动tcp

	if config.GetConf().Grpc.IsOpenRpc {
		go grpc_server.InitGrpcServer() //启动grpc
	}
	task.StartTask() //开启异步定时任务
}

// 启动创建文件夹
func mkdir() {
	list := [...]string{"config", "log"}
	for _, v := range list {
		if Dir.MkDirAll(v) {
			log.Printf("文件夹%s已经存在！", v)
		} else {
			log.Printf("创建文件夹%s成功！", v)
		}
	}
}
