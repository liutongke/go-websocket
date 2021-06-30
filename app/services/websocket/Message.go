package websocket

import (
	"fmt"
)

//给对应群组发送消息
func SendMsgToGroup() {

}

//给所有用户发送消息
func SendMsgToALl() {

}

//给用户发送信息需要判断是否在本机
func SendMsgToUser(toUserId int, data map[string]interface{}) {
	//toClient := GetClientHub().GetClientByUserId(toUserId)
	//b, _ := json.Marshal(Response{
	//	Err:  http.StatusOK,
	//	Msg:  "C2C friend msg",
	//	Data: data,
	//})
	//if toClient != nil { //本机上发送
	//	toClient.SendMsg(b)
	//}
	//
	//bindInfo := Center.GetBindInfo(toUserId) //不在本机上则调用GRPC
	//if bindInfo == (Center.BindUserInfo{}) {
	//	//用户未登录
	//	fmt.Println("用户未登录")
	//	return
	//}
	//fmt.Println("开始调用grpc去发送消息")
	//grpcClient.SendToUserMsg(bindInfo.RpcAddr, toUserId, b)
}

//发送给本机用户
func SendToUserMsgLocal(toUserId int, data []byte) {
	toClient := GetClientHub().GetClientByUserId(toUserId)
	if toClient == nil { //本机上发送
		fmt.Println("SendToUserMsgLocal跨区聊天了", "toClient", toClient)
		return
	}
	toClient.SendMsg(data)
}

//给分组发送消息
func SendToGroupMsg(groupId int, data []byte) {
	//获取所有的激活rpc
}

//给本地分组发送消息
func SendToGroupMsgToLocal(groupId int, data []byte) {
	toGroupClientList := GetClientHub().GetGroupClientList(groupId)
	if len(toGroupClientList) <= 0 {
		//分组不存在
		return
	}
	GetClientHub().SendGroupMsg(toGroupClientList, data) //发送消息
}
