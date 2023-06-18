package websocket

import (
	"encoding/json"
	"fmt"
	"go-websocket/app/services/bindCenter"
	"net/http"
	"strings"
)

// 给所有用户发送消息
func SendMsgALl(data map[string]interface{}) {
	serviceList := bindCenter.GetAllService()

	b, _ := json.Marshal(Response{
		Err:  http.StatusOK,
		Msg:  "SendMsgALl msg",
		Data: data,
	})

	NewRpcService().SendMsgALlLocal(b) //本机的用户来一波

	if len(serviceList) <= 0 {
		fmt.Println("服务器ip空")
		return
	}

	localAddr := bindCenter.GetServiceToStr()
	for _, addr := range serviceList {
		if strings.Compare(localAddr, addr) != 0 {
			//grpcClient.SendMsgALl(addr, b) //全服发送
		}
	}
	return
}

type RpcService struct {
}

func NewRpcService() *RpcService {
	return &RpcService{}
}

// 给本地所有用户发送消息
func (r *RpcService) SendMsgALlLocal(data []byte) {
	GetClientHub().Broadcast <- data
}

// 给用户发送信息需要判断是否在本机
func SendUserMsg(toUserId int, data map[string]interface{}) {
	toClient := GetClientHub().GetClientByUserId(toUserId)
	b, _ := json.Marshal(Response{
		Err:  http.StatusOK,
		Msg:  "C2C friend msg",
		Data: data,
	})
	if toClient != nil { //本机上发送
		toClient.SendMsg(b)
	}

	bindInfo := bindCenter.GetBindInfo(toUserId) //不在本机上则调用GRPC
	if bindInfo == (bindCenter.BindUserInfo{}) {
		//用户未登录
		fmt.Println("用户未登录")
		return
	}
	fmt.Println("开始调用grpc去发送消息")
	//grpcClient.SendUserMsg(bindInfo.RpcAddr, toUserId, b)
}

// 发送给本机用户
func (r *RpcService) SendUserMsgLocal(toUserId int, data []byte) {
	toClient := GetClientHub().GetClientByUserId(toUserId)
	if toClient == nil { //本机上发送
		fmt.Println("SendToUserMsgLocal跨区聊天了", "toClient", toClient)
		return
	}
	toClient.SendMsg(data)
}

// 给分组发送消息
func SendMsgToGroup(groupId int, data map[string]interface{}) {
	groupClientList := GetClientHub().GetGroupClientList(groupId)

	b, _ := json.Marshal(Response{
		Err:  http.StatusOK,
		Msg:  "SendGroupMsg msg",
		Data: data,
	})

	if len(groupClientList) > 0 { //本机上发送
		GetClientHub().SendGroupMsg(groupClientList, b)
	}

	serviceList := bindCenter.GetAllService()

	if len(serviceList) <= 0 {
		fmt.Println("服务器ip空")
		return
	}

	localAddr := bindCenter.GetServiceToStr()
	for _, addr := range serviceList {
		if strings.Compare(localAddr, addr) != 0 {
			//grpcClient.SendGroupMsgToLocal(addr, groupId, b) //给对应的群发送消息
		}
	}
}

// 给本地分组发送消息
func (r *RpcService) SendToGroupMsgToLocal(groupId int, data []byte) {
	toGroupClientList := GetClientHub().GetGroupClientList(groupId)
	if len(toGroupClientList) <= 0 {
		fmt.Println("群不存在") //分组不存在
		return
	}
	GetClientHub().SendGroupMsg(toGroupClientList, data) //发送消息
}
