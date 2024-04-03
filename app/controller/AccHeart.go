package controller

import (
	"go-websocket/app/services/websocket"
	"go-websocket/tools/timer"
	"net/http"
	"time"
)

func Heartbeat(c *websocket.Client, msg string) websocket.Response {
	//if strings.Compare("/heartbeat", gjson.Get(msg, "path").String()) != 0 {
	//	return websocket.Response{
	//		Err:  http.StatusBadRequest,
	//		Msg:  "心跳",
	//		Data: currentTime,
	//	}
	//}
	//if v, ok := c.UserId; !ok {
	//
	//}
	currentTime := uint64(time.Now().Unix())
	c.Heartbeat(currentTime)
	data := make(map[string]interface{})
	data["tm"] = currentTime
	data["strTm"] = timer.GetNowStr()

	return websocket.Response{
		Code: http.StatusOK,
		Msg:  "心跳biubiubiu---",
		Data: data,
	}
}
