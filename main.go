package main

import (
	"fmt"
	"go-websocket/app/services"
	"go-websocket/config"
	routers "go-websocket/routes"
	"go-websocket/tools"
	"go-websocket/tools/Timer"
	"go-websocket/tools/Tools"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //全局随机种子

	services.InitializeService()
	r := routers.SetupRouter()

	tools.EchoSuccess(fmt.Sprintf("服务器时间：%s", Timer.GetNowStr()))
	tools.EchoSuccess(fmt.Sprintf("本机IP地址GetServIp方式获得：%s；GetLocalIp方式获得：%s", Tools.GetServIp(), Tools.GetLocalIp()))
	tools.EchoSuccess(fmt.Sprintf("开启的服务HTTP/WS端口: %s GRPC端口: %s", config.GetConf().Server.Port, config.GetConf().Grpc.RpcPort))

	r.Run(":" + config.GetConf().Server.Port) // listen and serve on 0.0.0.0:8080
}
