package websocket

import (
	"encoding/json"
	"go-websocket/app/services/bind_center"
	"go-websocket/app/services/grpc_client"
	"log"
	"net/http"
	"strings"
)

// SendMsgALlUser 给所有用户发送消息
func SendMsgALlUser(data map[string]interface{}) {
	serviceList := bind_center.GetAllService()

	b, _ := json.Marshal(Response{
		Code: http.StatusOK,
		Msg:  "SendMsgALl msg",
		Data: data,
	})

	NewRpcService().SendMsgALlLocal(b) //本机的用户来一波

	if len(serviceList) <= 0 {
		log.Println("服务器ip空")
		return
	}

	localAddr := bind_center.GetServiceToStr()
	for _, serverInfo := range serviceList {
		if strings.Compare(localAddr, serverInfo.Addr) != 0 {
			grpc_client.SendMsgALl(serverInfo.Addr, b) //全服发送
		}
	}
	return
}

type RpcService struct {
}

func NewRpcService() *RpcService {
	return &RpcService{}
}

// SendMsgALlLocal 给本地所有用户发送消息
func (r *RpcService) SendMsgALlLocal(data []byte) {
	GetClientHub().Broadcast <- data
}

// SendMsgByUserId 给用户发送信息需要判断是否在本机
func SendMsgByUserId(toUserId int, data map[string]interface{}) {
	toClient := GetClientHub().GetClientByUserId(toUserId)

	b, _ := json.Marshal(Response{
		Code: http.StatusOK,
		Msg:  "C2C friend msg",
		Data: data,
	})
	if toClient != nil { //本机上发送
		toClient.SendMsg(b)
		return
	}

	bindInfo := bind_center.GetBindInfo(toUserId) //不在本机上则调用GRPC
	if bindInfo == (bind_center.BindUserInfo{}) {
		//用户未登录
		log.Println("用户未登录")
		return
	}
	log.Println("开始调用grpc去发送消息")
	grpc_client.SendUserMsg(bindInfo.RpcAddr, toUserId, b)
}

// SendUserMsgLocal 发送给本机用户
func (r *RpcService) SendUserMsgLocal(toUserId int, data []byte) {
	toClient := GetClientHub().GetClientByUserId(toUserId)
	if toClient == nil { //本机上发送
		log.Printf("SendToUserMsgLocal跨区聊天了,toClient:%v", toClient)
		return
	}
	toClient.SendMsg(data)
}

// SendMsgToGroup 给分组发送消息
func SendMsgToGroup(groupId int, data map[string]interface{}) {
	groupClientList := GetClientHub().GetGroupClientList(groupId)

	b, _ := json.Marshal(Response{
		Code: http.StatusOK,
		Msg:  "SendGroupMsg msg",
		Data: data,
	})

	if len(groupClientList) > 0 { //本机上发送
		GetClientHub().SendGroupMsg(groupClientList, b)
	}

	serviceList := bind_center.GetAllService()

	if len(serviceList) <= 0 {
		log.Println("服务器ip空")
		return
	}

	localAddr := bind_center.GetServiceToStr()
	for _, serverInfo := range serviceList {
		if strings.Compare(localAddr, serverInfo.Addr) != 0 {
			grpc_client.SendGroupMsgToLocal(serverInfo.Addr, groupId, b) //给对应的群发送消息
		}
	}
}

// SendToGroupMsgToLocal 给本地分组发送消息
func (r *RpcService) SendToGroupMsgToLocal(groupId int, data []byte) {
	toGroupClientList := GetClientHub().GetGroupClientList(groupId)
	if len(toGroupClientList) <= 0 {
		log.Println("群不存在") //分组不存在
		return
	}
	GetClientHub().SendGroupMsg(toGroupClientList, data) //发送消息
}
