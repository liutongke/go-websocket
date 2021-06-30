package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var clientHub *Hub

//获取hub实例
func GetClientHub() *Hub {
	return clientHub
}

// 获取用户LoginUsers下标key
func GetUserKey(userId int) (key int) {
	return userId
	//key = fmt.Sprintf("%s", userId)
	//return key
}

//根据clients发送全服消息
func (h *Hub) SendGlobalServMsg(msg []byte, clients []*Client) {
	for _, client := range clients {
		client.SendMsg(msg)
	}
}

//获取全服的用户登录
func (h *Hub) GetGlobalServClient() []*Client {
	var clients []*Client
	for _, v := range h.LoginUsers {
		clients = append(clients, v)
	}
	return clients
}

//用户登录
func (h *Hub) LoginUser(client *Client) {
	//将自己跟服务器绑定起来
	h.AddUser(client)
	h.AddClients(client)
	//h.addClient2Group(client.LoginZone, client)
	//models.NewEverClient().Login(client.UserId, client.LoginZone)          //初始化下数据库中的数据
	//client2.BindUidAndService(client.Uid, client.UserId, client.LoginZone) //登录服务器跟用户绑定一下
	b, _ := json.Marshal(Response{
		Id:   0,
		Err:  http.StatusOK,
		Msg:  "login succ",
		Data: nil,
	})
	client.SendMsg(b)
}

//添加用户
func (h *Hub) AddUser(client *Client) {
	h.LoginUsersLock.Lock()
	defer h.LoginUsersLock.Unlock()
	loginUserKey := GetUserKey(client.UserId)
	if _, ok := h.LoginUsers[loginUserKey]; !ok {
		h.LoginUsers[loginUserKey] = client
	}
	return
}

//添加clients
func (h *Hub) AddClients(client *Client) {
	h.ClientsLock.Lock()
	defer h.ClientsLock.Unlock()
	h.Clients[client] = true
	return
}

//用户退出登录
func (h *Hub) LoginOutUser(clients *Client) {
	fmt.Println("注销登录LoginOutUser", clients.UserId)
	loginUserKey := GetUserKey(clients.UserId)
	if client, ok := h.LoginUsers[loginUserKey]; ok {
		//models.NewEverClient().LogOut(client.UserId, client.LoginZone) //将用户数据存入数据库中去
		//client2.DelBindUidAndService(client.Uid)                       //解绑和服务器的绑定
		//h.delGroupClient(client.LoginZone, client)                     //移除区服分组
		delete(h.Clients, client)
		delete(h.LoginUsers, loginUserKey)
		close(client.Send)
	}
}

//通过UserId获取clients
func (h *Hub) GetClientByUserId(userId int) *Client {
	h.ClientsLock.RLock()
	defer h.ClientsLock.RUnlock()
	loginUserKey := GetUserKey(userId)
	if client, ok := h.LoginUsers[loginUserKey]; ok {
		return client
	}
	return nil
}

// GetClients
func (h *Hub) GetClients() (clients map[*Client]bool) {

	clients = make(map[*Client]bool)

	h.ClientsRange(func(client *Client, value bool) (result bool) {
		clients[client] = value

		return true
	})

	return
}

// 遍历
func (h *Hub) ClientsRange(f func(client *Client, value bool) (result bool)) {

	h.ClientsLock.RLock()
	defer h.ClientsLock.RUnlock()

	for key, value := range h.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}

	return
}

// 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())
	clients := GetClientHub().GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			fmt.Println("心跳时间超时 关闭连接", client.UserId, client.LoginTime, client.HeartbeatTime)

			client.Ws.Close()
		}
	}
}
