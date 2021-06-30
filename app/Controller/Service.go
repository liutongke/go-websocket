package Controller

import (
	"github.com/gin-gonic/gin"
)

//发送全服信息
func SendGlobalServMsg(c *gin.Context) {
	//hub := websocket.GetClientHub()
	//clients := hub.GetGlobalServClient()
	//b, _ := json.Marshal(websocket.Response{
	//	Id:   0,
	//	Err:  http.StatusOK,
	//	Msg:  "SendGlobalServMsg",
	//	Data: "这是一条全服信息",
	//})
	//hub.SendGlobalServMsg(b, clients)
}
