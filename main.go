package main

import (
	"encoding/json"
	"fmt"
	"go-websocket/app/services/bindCenter"
	"go-websocket/app/services/grpc"
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
	clientv3 "go.etcd.io/etcd/client/v3"
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
	go Etcds.NewEtcdDiscovery(map[string]Etcds.FunDiscovery{"put": EventPut, "del": EventDel}).EtcdStartDiscovery([]string{"/prefix1", "/prefix2", "/net", "go-nat-x", Etcds.ETCD_SERVER_LIST, Etcds.ETCD_PREFIX_ACCOUNT_INFO})
	go Etcds.NewEtcdRegister().EtcdStartRegister(RegisterServer)
	mkdir()             //启动项目时判断下文件夹
	Logger.InitLogger() //初始化日志
	routers.InitWsRouters()
	validates.InitValidate() //初始化参数验证
	socket.StartTcp()        //启动tcp
	DbLine.InitDbLine()
	RdLine.InitRdLine()

	if config.GetConf().Grpc.IsOpenRpc {
		go grpc.InitGrpcServer() //启动grpc
	}
	bindCenter.NewService().InitSetServer() //初始化绑定
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

func EventPut(event *clientv3.Event) {
	log.Printf("watch put test---------->key:%q val:%q", event.Kv.Key, event.Kv.Value)
}
func EventDel(event *clientv3.Event) {
	log.Printf("watch del test---------->key:%q val:%q", event.Kv.Key, event.Kv.Value)
}

type ServerInfo struct {
	ServerIp string `json:"server-ip"`
	Rpcport  string `json:"rpc-port"`
	Tm       string `json:"tm"`
}

// RegisterServer 注册主机
func RegisterServer(e *Etcds.EtcdRegister) {
	key := fmt.Sprintf("%s%s:nat-x", Etcds.ETCD_SERVER_LIST, Tools.GetLocalIp())

	info := ServerInfo{
		ServerIp: Tools.GetLocalIp(),
		Rpcport:  "nat-x",
		Tm:       Timer.GetNowStr(),
	}
	// 将Person对象转换为JSON字符串
	val, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	e.PutKey(key, string(val))
}
