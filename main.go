package main

import (
	"fmt"
	"go-websocket/app/services/grpcService"
	"go-websocket/app/services/socket"
	"go-websocket/app/services/task"
	"go-websocket/config"
	routers "go-websocket/routes"
	"go-websocket/tools/DbLine"
	"go-websocket/tools/Dir"
	"go-websocket/tools/Logger"
	"go-websocket/tools/RdLine"
	"go-websocket/tools/Timer"
	"go-websocket/tools/Tools"
	"go-websocket/tools/validates"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //全局随机种子
	Init()
	r := routers.SetupRouter()
	task.StartTask() //开启异步定时任务
	formatNow := time.Now().Format("2006-01-02 15:04:05")

	fmt.Printf("服务器时间：%s,北京时间：%s \n", formatNow, Timer.GetNowStr())
	fmt.Printf("本机IP地址GetServIp方式获得：%s；GetLocalIp方式获得：%s \n", Tools.GetServIp(), Tools.GetLocalIp())
	fmt.Printf("开启的服务端口: %s \n", config.GetConfClient().Server.Port)
	r.Run(":" + config.GetConfClient().Server.Port) // listen and serve on 0.0.0.0:8080
}

// 初始化配置
func Init() {
	p := socket.NewPool()
	p.StartPool()
	mkdir()             //启动项目时判断下文件夹
	Logger.InitLogger() //初始化日志
	routers.InitWsRouters()
	validates.InitValidate() //初始化参数验证
	config.Init()            //初始化配置文件
	socket.StartTcp()        //启动tcp
	DbLine.InitDbLine()
	RdLine.InitRdLine()
	if config.GetConfClient().CommonConf.IsOpenRpc {
		go grpcService.InitGrpcServer() //启动grpc
	}
}

// 启动创建文件夹
func mkdir() {
	list := [...]string{"runtime", "config"}
	for _, v := range list {
		if Dir.MkDirAll(v) {
			fmt.Println(fmt.Sprintf("文件夹%s已经存在！", v))
		} else {
			fmt.Println(fmt.Sprintf("创建文件夹%s成功！", v))
		}
	}
}
