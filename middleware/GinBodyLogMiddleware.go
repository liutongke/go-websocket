package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-websocket/tools/logger"
	"go-websocket/tools/timer"
)

// 打印返回信息
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func GinBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	if gin.IsDebugging() {
		debugInfo := "[GIN-HTTP] " + timer.GetNowStr() + " -->Request path:" + c.Request.URL.Path + "\n-->Response body: " + blw.body.String()
		fmt.Println(debugInfo)
		logger.Info(debugInfo)
	}
}
