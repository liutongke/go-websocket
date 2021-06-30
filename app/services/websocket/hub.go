package websocket

import (
	"github.com/tidwall/gjson"
	"sync"
)

//所有用户管理
type Hub struct {
	Clients     map[*Client]bool //注册客户端
	ClientsLock sync.RWMutex     // 读写锁

	LoginUsers     map[int]*Client //已经链接上来的用户
	LoginUsersLock sync.RWMutex    // 读写锁

	Broadcast chan []byte //全局待发送的消息，所有区服

	Register   chan *Client //管道通信的连接上来的客户端
	Unregister chan *Client //管道通信需要注销的客户端

	GroupLock sync.RWMutex
	Groups    map[int][]*Client
	GroupMsg  chan string //对应区内的信息
}

func newClientHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		LoginUsers: make(map[int]*Client),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Groups:     make(map[int][]*Client),
		GroupMsg:   make(chan string),
	}
}

//启动hub
func StartClientHub() *Hub {
	clientHub = newClientHub()
	go clientHub.run()
	return clientHub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register: //记录登录用户
			h.LoginUser(client) //记录登录时间
		case client := <-h.Unregister: //注销登录用户
			if _, ok := h.Clients[client]; ok {
				h.LoginOutUser(client) //注销登录
			}
		case message := <-h.Broadcast: //全服务器的信息
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		case zoneMsg := <-h.GroupMsg: //单服内的全部消息
			zone := gjson.Get(zoneMsg, "zone").Int()
			msg := gjson.Get(zoneMsg, "msg").String()
			if zone > 0 && len(msg) > 0 {
				//h.SendGroupMsg(cast.ToInt(zone), []byte(msg))
			}
		}
	}
}
