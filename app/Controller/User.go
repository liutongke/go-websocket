package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-websocket/utils/Jwt"
	"go-websocket/utils/Timer"
	"net/http"
)

//获取用户信息
func GetUserInfo(c *gin.Context) {
	data := make(map[string]interface{}, 3)
	//根据实际业务修改
	userId := Timer.NowUnix()
	openId := uuid.NewString()
	data["userId"] = userId
	data["openId"] = openId
	token, _ := Jwt.GenerateToken(userId, openId)
	data["token"] = token
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "success",
		"data": data,
	})
	return
}
