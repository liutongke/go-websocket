package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go-websocket/app/services/websocket"
	"net/http"
	"runtime"
	"time"
)

//发送全服信息
func SendGlobalServMsg(c *gin.Context) {
	data := make(map[string]interface{}, 10)
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["uuid"] = uuid.NewString()
	websocket.SendMsgALl(data)
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendGlobalServMsg success",
		"data": data,
	})
	return
}

//给单个用户发送消息
func SendUserMsg(c *gin.Context) {
	data := make(map[string]interface{})
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["uuid"] = uuid.NewString()
	userId := cast.ToInt(c.Param("userId"))
	websocket.SendUserMsg(userId, data)
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendUserMsg success",
		"data": data,
	})
	return
}

//发送分组消息
func SendGroupMsg(c *gin.Context) {
	data := make(map[string]interface{})
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["uuid"] = uuid.NewString()
	groupId := cast.ToInt(c.Param("groupId"))
	websocket.SendMsgToGroup(groupId, data)
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendGroupMsg success",
		"data": data,
	})
	return
}

// 查询系统状态
func Status(c *gin.Context) {
	data := make(map[string]interface{})
	NumGoroutine := runtime.NumGoroutine()
	NumCPU := runtime.NumCPU()
	data["NumGoroutine"] = NumGoroutine // goroutine数量
	data["NumCPU"] = NumCPU
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendGroupMsg success",
		"data": data,
	})
	return
}
