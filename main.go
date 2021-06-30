package main

import (
	"fmt"
	"go-websocket/app/services/grpcService"
	"go-websocket/app/services/task"
	"go-websocket/config"
	routers "go-websocket/routes"
	services "go-websocket/utils"
	"go-websocket/utils/Dir"
	"go-websocket/utils/Logger"
	"go-websocket/utils/Timer"
	"go-websocket/utils/validates"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //全局随机种子
	Init()
	r := routers.SetupRouter()
	task.StartTask() //开启异步定时任务
	formatNow := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(fmt.Sprintf("服务器时间：%s,北京时间：%s", formatNow, Timer.NowStr()))
	fmt.Println(fmt.Sprintf("本机IP地址GetServIp方式获得：%s；GetLocalIp方式获得：%s", services.GetServIp(), services.GetLocalIp()))
	fmt.Println(fmt.Sprintf("开启的RPC端口：%s", config.GetConfClient().Server.RpcPort))
	r.Run(":" + config.GetConfClient().Server.Port) // listen and serve on 0.0.0.0:8080
}

//初始化配置
func Init() {
	mkdir()             //启动项目时判断下文件夹
	Logger.InitLogger() //初始化日志
	routers.InitWsRouters()
	validates.InitValidate()        //初始化参数验证
	config.Init()                   //初始化配置文件
	go grpcService.InitGrpcServer() //启动grpc
}

//启动创建文件夹
func mkdir() {
	list := [...]string{"runtime", "config"}
	for _, v := range list {
		if Dir.Mkdir(v) == "" {
			fmt.Println(fmt.Sprintf("文件夹%s已经存在！", v))
		} else {
			fmt.Println(fmt.Sprintf("创建文件夹%s成功！", v))
		}
	}
}
