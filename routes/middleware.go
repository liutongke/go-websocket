package routers

import (
	"github.com/gin-gonic/gin"
	"go-websocket/middleware"
)

// 中间件
func middlewares(r *gin.Engine) {
	r.Use(middleware.GinBodyLogMiddleware) //打印日志
	r.Use(middleware.Cors())               //开启中间件 允许使用跨域请求
	r.Use(middleware.Recovers())           //捕获错误
}
