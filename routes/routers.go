package routers

import (
	"github.com/gin-gonic/gin"
	"go-websocket/app/Middleware"
	"go-websocket/app/services/websocket"
	"go-websocket/config"
	"go-websocket/tools/Tools"
)

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	var r *gin.Engine
	if Tools.IsDebug() { //开发模式
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
		r.GET("/ws", websocket.StartClientHub().StartWs)
	}
}

// 中间件
func middleware(r *gin.Engine) {
	r.Use(Middleware.GinBodyLogMiddleware) //打印日志
	r.Use(Middleware.Cors())               //开启中间件 允许使用跨域请求
	r.Use(Middleware.Recovers())           //捕获错误
}
