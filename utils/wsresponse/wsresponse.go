package wsresponse

import (
	"go-websocket/app/services/websocket"
	"net/http"
)

func ReturnJson(errCode int, msg string, data interface{}) websocket.Response {
	return websocket.Response{
		Err:  errCode,
		Msg:  msg,
		Data: data,
	}
}

// 直接返回成功
func Success(msg string, data interface{}) websocket.Response {
	return ReturnJson(http.StatusOK, msg, data)
}

// 失败的业务逻辑
func Fail(msg string, data interface{}) websocket.Response {
	return ReturnJson(http.StatusBadRequest, msg, data)
}
