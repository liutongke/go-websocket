package websocket

const (
	heartbeatExpirationTime = 10 * 60 // 用户连接超时时间
)

// 用户心跳
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime

	return
}

// 心跳超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}

	return
}

//给对应群组发送消息
func SendMsgToGroup() {

}

//给所有用户发送消息
func SendMsgToALl() {

}

//给用户发送信息需要判断是否在本机
func SendMsgToUser(toUserId, toUid, zone int, data map[string]interface{}) {
	//toClient := GetClientHub().GetClientByUserId(toUserId)
	//b, _ := json.Marshal(Response{
	//	Err:  http.StatusOK,
	//	Msg:  "C2C friend msg",
	//	Data: data,
	//})
	//if toClient != nil { //本机上发送
	//	if zone != toClient.LoginZone {
	//		fmt.Println("跨区聊天了")
	//		return
	//	}
	//	toClient.SendMsg(b)
	//}
	//bindInfo := client.GetBindInfo(toUid) //不在本机上则调用GRPC
	//if bindInfo == (client.BindUserInfo{}) {
	//	//用户未登录
	//	fmt.Println("用户未登录")
	//	return
	//}
	//if bindInfo.LoginZone != zone {
	//	fmt.Println("用户不在同一个区")
	//	return
	//}
	//fmt.Println("开始调用grpc去发送消息")
	//grpcClient.SendToUserMsg(bindInfo.Addr, toUserId, toUid, zone, b)
}

//发送给本机用户
func SendToUserMsgLocal(toUserId, toUid, zone int, data []byte) {
	//toClient := GetClientHub().GetClientByUserId(toUserId)
	//if toClient == nil || zone != toClient.LoginZone { //本机上发送
	//	fmt.Println("SendToUserMsgLocal跨区聊天了", "toClient", toClient, "zone:", zone, "LoginZone", toClient.LoginZone)
	//	return
	//}
	//toClient.SendMsg(data)
}
