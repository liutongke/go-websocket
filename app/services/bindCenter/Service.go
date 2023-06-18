package bindCenter

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-websocket/config"
	"go-websocket/tools/Etcds"
	"go-websocket/tools/Tools"
	"log"
	"strings"
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
		}
		fmt.Printf("Key: %s, Value: %s\n", item.Key, item.Value)
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
	log.Println("SetServer:", data, string(val))
	s.ServiceList[data.ServerIp] = &Service{
		Ip:      data.ServerIp,
		Rpcport: data.Rpcport,
	}
}

// 下线
func (s *Server) DelServer(val []byte) {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	var data ServerInfo
	err := json.Unmarshal(val, &data)
	if err != nil {
		fmt.Println("Failed to unmarshal JSON:", err)
		return
	}
	log.Println("DelServer:", data, string(val))
	delete(s.ServiceList, data.ServerIp)
}

// 清除超时的服务
func DelTimeoutService() {

}

// 获取服务器的信息
func GetService() *Service {
	return &Service{
		Ip:      Tools.GetLocalIp(),
		Rpcport: config.GetConf().Server.RpcPort,
	}
}

// 获取ip+端口组合体字符串
func GetServiceToStr() (str string) {
	return fmt.Sprintf("%s:%s", Tools.GetLocalIp(), config.GetConf().Server.RpcPort)
}

// 将ip+端口组合体字符串拆解
func GetStrToService(str string) (server *Service, err error) {
	list := strings.Split(str, ":")
	if len(list) != 2 {

		return nil, errors.New("err")
	}
	server = &Service{
		Ip:      list[0],
		Rpcport: list[1],
	}
	return
}

// 获取所有的服务器
func GetAllService() map[string]*Service {
	return GetServiceConf().ServiceList
}
