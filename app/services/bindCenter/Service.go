package bindCenter

import (
	"errors"
	"fmt"
	"go-websocket/global"
	"go-websocket/tools/RdLine"
	"go-websocket/tools/Timer"
	"strings"
)

type Service struct {
	Ip      string
	Rpcport string
}

func init() {
	//global.Gservip = Tools.GetLocalIp()
	//global.Rpcport = config.GetConfClient().Server.RpcPort
	global.Gservip = "192.168.1.106"
	global.Rpcport = "8972"
}

// 获取服务器的信息
func GetService() *Service {
	return &Service{
		Ip:      global.Gservip,
		Rpcport: global.Rpcport,
	}
}

// 获取ip+端口组合体字符串
func GetServiceToStr() (str string) {
	s := GetService()
	str = fmt.Sprintf("%s:%s", s.Ip, s.Rpcport)
	return
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
func GetAllService() []string {
	RdLine := RdLine.GetRedisClient()
	defer RdLine.CloseRedisClient()
	list, err := RdLine.Strings(RdLine.Exec("ZREVRANGEBYSCORE", getServiceKey(), Timer.GetNowStr(), Timer.GetOffsetUnix(-(F5ServiceTm * 3))))
	if err != nil {
		return nil
	}
	return list
}
