package websocket

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map/v2"
	"sync"

	"github.com/tidwall/gjson"
)

// Hub 管理所有用户
type Hub struct {
	LoginUsers cmap.ConcurrentMap[string, *Client] // 已经链接上来的用户

	Broadcast  chan []byte  // 全局待发送的消息，所有区服
	Register   chan *Client // 管道通信的连接上来的客户端
	Unregister chan *Client // 管道通信需要注销的客户端

	GroupLock sync.RWMutex      // 组锁
	Groups    map[int][]*Client // 组内客户端列表
	GroupMsg  chan string       // 对应区内的信息
}

// newClientHub 创建一个新的 Hub 实例
func newClientHub() *Hub {
	return &Hub{
		LoginUsers: cmap.New[*Client](),
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Groups:     make(map[int][]*Client),
		GroupMsg:   make(chan string),
	}
}

// StartClientHub 启动客户端 Hub
func StartClientHub() *Hub {
	clientHub = newClientHub()
	go clientHub.run()
	return clientHub
}

// run 启动 Hub 运行循环
func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.LoginUser(client)
		case client := <-h.Unregister:
			h.LoginOutUser(client)
		case message := <-h.Broadcast:
			// 使用IterBuffered迭代并发Map
			for tuple := range h.LoginUsers.IterBuffered() {
				fmt.Printf("键：%s 值：%d\n", tuple.Key, tuple.Val.OpenId)
				select {
				case tuple.Val.Send <- message:
				default:
					close(tuple.Val.Send)
				}
			}
		case zoneMsg := <-h.GroupMsg:
			zone := gjson.Get(zoneMsg, "zone").Int()
			msg := gjson.Get(zoneMsg, "msg").String()
			if zone > 0 && len(msg) > 0 {
				//h.SendGroupMsg(cast.ToInt(zone), []byte(msg))
			}
		}
	}
}
