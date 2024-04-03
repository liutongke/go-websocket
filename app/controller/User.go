package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-websocket/tools/jwt"
	"go-websocket/tools/timer"
	"net/http"
)

// 获取用户信息
func GetUserInfo(c *gin.Context) {
	data := make(map[string]interface{}, 3)
	//根据实际业务修改
	userId := timer.GetNowUnix()
	openId := uuid.NewString()
	data["userId"] = userId
	data["openId"] = openId
	token, _ := jwt.GenerateToken("AppName", "appid", "module", "ver", openId, int(userId))
	data["token"] = token
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "success",
		"data": data,
	})
	return
}
