package bind_center

import (
	"encoding/json"
	"fmt"
	"go-websocket/config"
	"go-websocket/tools/Etcds"
	"go-websocket/tools/Timer"
	"go-websocket/tools/Tools"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"regexp"
	"sync"
)

var (
	rwMutex sync.RWMutex
	service *Server
)

type Server struct {
	ServiceList map[string]*Service
}
type Service struct {
	Addr    string
	Ip      string
	Rpcport string
}

func GetServiceConf() *Server {
	return service
}
func NewService() *Server {
	service = &Server{
		ServiceList: make(map[string]*Service),
	}
	return service
}

type ServerInfo struct {
	ServerIp string `json:"server-ip"`
	Rpcport  string `json:"rpc-port"`
	Tm       string `json:"tm"`
}

func EventPut(event *clientv3.Event) {
	log.Printf("bind center watch put test---------->key:%q val:%q", event.Kv.Key, event.Kv.Value)
	GetServiceConf().SetServer(event.Kv.Value)
}
func EventDel(event *clientv3.Event) {
	log.Printf("bind center watch del test---------->key:%q val:%q", event.Kv.Key, event.Kv.Value)
	GetServiceConf().DelServer(event.Kv.Key)
}

// RegisterServer 将本机注册到etcd
func RegisterServer(e *Etcds.EtcdRegister) {
	key := fmt.Sprintf("%s%s:go-websocket", Etcds.ETCD_SERVER_LIST, Tools.GetLocalIp())

	info := ServerInfo{
		ServerIp: Tools.GetLocalIp(),
		Rpcport:  config.GetConf().Grpc.RpcPort,
		Tm:       Timer.GetNowStr(),
	}
	// 将Person对象转换为JSON字符串
	val, err := json.Marshal(info)
	if err != nil {
		log.Fatal(err)
	}
	e.PutKey(key, string(val))
}

// InitSetServer 初始化下服务器列表
func (s *Server) InitSetServer() {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	prefix, err := Etcds.GetPrefix(Etcds.ETCD_SERVER_LIST)
	if err != nil {
		return
	}
	log.Println("InitSetServer:", prefix.Kvs)
	for _, item := range prefix.Kvs {
		var data ServerInfo
		err := json.Unmarshal(item.Value, &data)
		if err != nil {
			fmt.Println("Failed to unmarshal JSON:", err)
			return
		}
		s.ServiceList[data.ServerIp] = &Service{
			Ip:      data.ServerIp,
			Rpcport: data.Rpcport,
			Addr:    fmt.Sprintf("%s:%s", data.ServerIp, data.Rpcport),
		}
		//fmt.Printf("Key: %s, Value: %s\n", item.Key, item.Value)
	}
	return
}

// 将服务器信息存入
func (s *Server) SetServer(val []byte) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	var data ServerInfo
	err := json.Unmarshal(val, &data)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}
	log.Println("SetServer:", string(val))
	s.ServiceList[data.ServerIp] = &Service{
		Ip:      data.ServerIp,
		Rpcport: data.Rpcport,
		Addr:    fmt.Sprintf("%s:%s", data.ServerIp, data.Rpcport),
	}
}

// 下线
func (s *Server) DelServer(key []byte) {
	rwMutex.Lock()
	defer rwMutex.Unlock()

	ipRegex := `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})` // 定义匹配IP地址的正则表达式
	re := regexp.MustCompile(ipRegex)                 // 创建正则表达式对象
	ipMatches := re.FindStringSubmatch(string(key))   // 查找匹配的IP地址
	// 提取第一个匹配的IP地址
	if len(ipMatches) > 1 {
		delete(s.ServiceList, ipMatches[1])
		fmt.Println("DelServer:", ipMatches[1])
	} else {
		fmt.Println("No IP address found")
	}
}

// GetServiceToStr 获取ip+端口组合体字符串
func GetServiceToStr() (str string) {
	return fmt.Sprintf("%s:%s", Tools.GetLocalIp(), config.GetConf().Grpc.RpcPort)
}

// GetAllService 获取所有的服务器
func GetAllService() map[string]*Service {
	return GetServiceConf().ServiceList
}
