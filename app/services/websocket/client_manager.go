package websocket

import (
	"encoding/json"
	"fmt"
	"go-websocket/app/services/bind_center"
	"log"
	"net/http"
	"time"
)

var clientHub *Hub

// GetClientHub 获取hub实例
func GetClientHub() *Hub {
	return clientHub
}

// GetUserKey 获取用户LoginUsers下标key
func GetUserKey(userId int) (key string) {
	return fmt.Sprintf("%d", userId)
}

// SendGlobalServMsg 根据clients发送全服消息
func (h *Hub) SendGlobalServMsg(msg []byte, clients []*Client) {
	for _, client := range clients {
		client.SendMsg(msg)
	}
}

// GetGlobalServClient 获取全服的用户登录
func (h *Hub) GetGlobalServClient() []*Client {
	var clients []*Client

	for tuple := range h.LoginUsers.IterBuffered() {
		clients = append(clients, tuple.Val)
	}

	return clients
}

// LoginUser 用户登录
func (h *Hub) LoginUser(client *Client) {
	//将自己跟服务器绑定起来
	h.AddUser(client)
	//h.addClient2Group(client.LoginZone, client)
	//models.NewEverClient().Login(client.UserId, client.LoginZone)          //初始化下数据库中的数据
	bind_center.BindUidAndService(client.UserId) //登录服务器跟用户绑定一下
	//client2.BindUidAndService(client.Uid, client.UserId, client.LoginZone) //登录服务器跟用户绑定一下
	b, _ := json.Marshal(Response{
		Id:   0,
		Code: http.StatusOK,
		Msg:  "login success",
		Data: nil,
	})
	client.SendMsg(b)
}

// AddUser 添加用户
func (h *Hub) AddUser(client *Client) {
	loginUserKey := GetUserKey(client.UserId)
	h.LoginUsers.Set(loginUserKey, client)
	return
}

// LoginOutUser 用户退出登录
func (h *Hub) LoginOutUser(clients *Client) {
	log.Printf("注销%d用户登录", clients.UserId)
	loginUserKey := GetUserKey(clients.UserId)

	if client, ok := h.LoginUsers.Get(loginUserKey); ok {
		//models.NewEverClient().LogOut(client.UserId, client.LoginZone) //将用户数据存入数据库中去
		bind_center.DelBindUidAndService(client.UserId) //解绑和服务器的绑定
		//h.delGroupClient(client.LoginZone, client)      //移除区服分组
		h.LoginUsers.Remove(loginUserKey)
		close(client.Send)
	}
}

// GetClientByUserId 通过UserId获取clients
func (h *Hub) GetClientByUserId(userId int) *Client {

	loginUserKey := GetUserKey(userId)
	if client, ok := h.LoginUsers.Get(loginUserKey); ok {
		return client
	}

	return nil
}

// GetClients 获取client列表
func (h *Hub) GetClients() (clients map[*Client]bool) {

	clients = make(map[*Client]bool)

	for tuple := range h.LoginUsers.IterBuffered() {
		clients[tuple.Val] = true
	}

	return
}

// ClearTimeoutConnections 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())
	clients := GetClientHub().GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			log.Println("心跳时间超时 关闭连接", client.UserId, client.LoginTime, client.HeartbeatTime)

			client.Ws.Close()
		}
	}
}
