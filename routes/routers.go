package routers

import (
	"github.com/gin-gonic/gin"
	"go-websocket/tools/utils"
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
	middlewares(r) //调用中间件
	apiRouter(r)   //api路由
	wsRouter(r)    //ws路由
	return r
}
