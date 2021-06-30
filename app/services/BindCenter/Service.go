package BindCenter

import (
	"errors"
	"fmt"
	"go-websocket/config"
	"go-websocket/global"
	"go-websocket/utils"
	"strings"
)

type Service struct {
	Ip      string
	Rpcport string
}

func init() {
	global.Gservip = utils.GetLocalIp()
	global.Rpcport = config.GetConfClient().Server.RpcPort
}

//获取服务器的信息
func GetService() *Service {
	return &Service{
		Ip:      global.Gservip,
		Rpcport: global.Rpcport,
	}
}

//获取ip+端口组合体字符串
func GetServiceToStr() (str string) {
	s := GetService()
	str = fmt.Sprintf("%s:%s", s.Ip, s.Rpcport)
	return
}

//将ip+端口组合体字符串拆解
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
