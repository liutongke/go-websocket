package main

import (
	"fmt"
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
	"go-websocket/tools/Timer"
	"go-websocket/tools/Tools"
	"go-websocket/tools/validates"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //全局随机种子
	Init()
	r := routers.SetupRouter()
	task.StartTask() //开启异步定时任务

	fmt.Printf("服务器时间：%s \n", Timer.GetNowStr())
	fmt.Printf("本机IP地址GetServIp方式获得：%s；GetLocalIp方式获得：%s \n", Tools.GetServIp(), Tools.GetLocalIp())
	fmt.Printf("开启的服务HTTP/WS端口: %s GRPC端口: %s \n", config.GetConf().Server.Port, config.GetConf().Grpc.RpcPort)
	r.Run(":" + config.GetConf().Server.Port) // listen and serve on 0.0.0.0:8080
}

// 初始化配置
func Init() {
	p := socket.NewPool()
	p.StartPool()
	config.Init() //初始化配置文件
	go Etcds.NewEtcdDiscovery(map[string]Etcds.FunDiscovery{"put": bind_center.EventPut, "del": bind_center.EventDel}).EtcdStartDiscovery([]string{"/prefix1", "/prefix2", "/net", "go-nat-x", Etcds.ETCD_SERVER_LIST, Etcds.ETCD_PREFIX_ACCOUNT_INFO})
	go Etcds.NewEtcdRegister().EtcdStartRegister(bind_center.RegisterServer)
	mkdir()             //启动项目时判断下文件夹
	Logger.InitLogger() //初始化日志
	routers.InitWsRouters()
	validates.InitValidate() //初始化参数验证
	socket.StartTcp()        //启动tcp
	DbLine.InitDbLine()
	RdLine.InitRdLine()

	if config.GetConf().Grpc.IsOpenRpc {
		go grpc_server.InitGrpcServer() //启动grpc
	}
	bind_center.NewService().InitSetServer() //初始化绑定
}

// 启动创建文件夹
func mkdir() {
	list := [...]string{"config"}
	for _, v := range list {
		if Dir.MkDirAll(v) {
			log.Printf("文件夹%s已经存在！", v)
		} else {
			log.Printf("创建文件夹%s成功！", v)
		}
	}
}
