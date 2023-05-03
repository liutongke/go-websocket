package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-websocket/app/Controller"
	"net/http"
	"time"
)

// 路由
func apiRouter(r *gin.Engine) {
	r.MaxMultipartMemory = 8 << 20                               // 限制上传文件最大8 MiB
	r.GET("User/SendMsgALl", Controller.SendGlobalServMsg)       //发送全服信息
	r.GET("User/SendUserMsg/:userId", Controller.SendUserMsg)    //给单个用户发送消息
	r.GET("User/SendGroupMsg/:groupId", Controller.SendGroupMsg) //给单个用户发送消息
	r.GET("Status", Controller.Status)                           //服务器健康
	r.GET("User/GetInfo", Controller.GetUserInfo)                //获取用户信息
	r.GET("ping", func(c *gin.Context) {
		data := make(map[string]interface{})
		data["test"] = "test1data"
		c.JSON(http.StatusOK, gin.H{
			"msg":  "success",
			"data": uuid.NewString(),
			"tm:":  time.Now().Format("2006-01-02 15:04:05"),
			"err":  200,
		})
	})
}
