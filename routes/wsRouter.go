package routers

import (
	"github.com/gin-gonic/gin"
	"go-websocket/app/controller"
	"go-websocket/app/services/websocket"
	"go-websocket/config"
	"net/http"
	"time"
)

func wsRouter(r *gin.Engine) {
	if config.GetConf().WebSocket.IsOpenWebsocket { //是否开启websocket                                    //websocket连接
		r.GET("/ws/:token", websocket.StartClientHub().StartWs)
	}
}
func InitWsRouters() {
	websocket.Register("/ping", Ping)
	websocket.Register("/heartbeat", controller.Heartbeat)
	websocket.Register("/CreateGroup", controller.CreateGroup)
}

// 获取信息 {"id":123,"path":"/ping","ver":"1.0.0","data":""}
func Ping(c *websocket.Client, msg string) websocket.Response {
	data := make(map[string]interface{}, 10)
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["userId"] = c.UserId
	return websocket.Response{
		Code: http.StatusOK,
		Msg:  "获取用户信息",
		Data: data,
	}
}
