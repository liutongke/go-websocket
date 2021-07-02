package routers

import (
	"go-websocket/app/Controller"
	"go-websocket/app/services/websocket"
	"net/http"
	"time"
)

func InitWsRouters() {
	websocket.Register("/ping", Ping)
	websocket.Register("/heartbeat", Controller.Heartbeat)
	websocket.Register("/CreateGroup", Controller.CreateGroup)
}

//获取信息 {"id":123,"path":"/ping","ver":"1.0.0","data":""}
func Ping(c *websocket.Client, msg string) websocket.Response {
	data := make(map[string]interface{}, 10)
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["userId"] = c.UserId
	return websocket.Response{
		Err:  http.StatusOK,
		Msg:  "获取用户信息",
		Data: data,
	}
}
