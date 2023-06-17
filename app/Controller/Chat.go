package Controller

import (
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"go-websocket/app/services/websocket"
	"go-websocket/tools/wsresponse"
)

// 创建本地分组 {"id":123,"path":"/CreateGroup","ver":"1.0.0","data":{"groupId":11}}
func CreateGroup(c *websocket.Client, msg string) websocket.Response {
	groupId := cast.ToInt(gjson.Get(msg, "data.groupId").Int())
	websocket.GetClientHub().AddClient2Group(groupId, c)
	var data = make(map[string]interface{})

	return wsresponse.Success("CreateGroup", data)
}

// 1对1聊天 {"id":123,"path":"/Chat/C2C","ver":"1.0.0","data":{"toUid":11,"toMsg":"你好啊"}}
func C2C(c *websocket.Client, msg string) websocket.Response {
	//toUid := cast.ToInt(gjson.Get(msg, "data.toUid").Int()) //接收方的标识2
	//toMsg := gjson.Get(msg, "data.toMsg").String()
	//if Mgck.CheckWord(toMsg) {
	//	return wsresponse.Fail("违规信息", nil)
	//} //明感词过滤
	//userInfo := Models.GetUserInfoByUid(toUid)
	//toUserId := userInfo.UserId
	//if toUserId <= 0 {
	//	return wsresponse.Fail("C2C toUserId ERR OR user not login or not zone", nil)
	//}
	//var toData = make(map[string]interface{})
	//toData["acvMsg"] = toMsg
	//toData["acvUserId"] = c.UserId
	//toData["acvZone"] = c.LoginZone
	//websocket.SendToUserMsg(toUserId, toUid, c.LoginZone, toData) //发送消息 需要通过区查找，防止跨区发送了
	var data = make(map[string]interface{})
	return wsresponse.Success("C2C", data)
}
