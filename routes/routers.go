package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-websocket/app/Controller"
	middleware2 "go-websocket/app/middleware"
	"go-websocket/app/services/websocket"
	"go-websocket/config"
	"go-websocket/utils"
	"net/http"
	"time"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	var r *gin.Engine
	if utils.IsDebug() { //开发模式
		r = gin.Default()
	} else { //生产模式
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	middleware(r) //调用中间件
	apiRouter(r)  //api路由
	wsRouter(r)   //ws路由
	return r
}

func wsRouter(r *gin.Engine) {
	if config.GetConfClient().CommonConf.IsOpenWebsocket { //是否开启websocket                                    //websocket连接
		r.GET("/ws/:token", websocket.StartClientHub().StartWs)
	}
}

//路由
func apiRouter(r *gin.Engine) {
	r.MaxMultipartMemory = 8 << 20                // 限制上传文件最大8 MiB
	r.GET("User/GetInfo", Controller.GetUserInfo) //获取用户信息
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

//中间件
func middleware(r *gin.Engine) {
	r.Use(middleware2.GinBodyLogMiddleware) //打印日志
	r.Use(middleware2.Cors())               //开启中间件 允许使用跨域请求
	r.Use(middleware2.Recovers())           //捕获错误
}
