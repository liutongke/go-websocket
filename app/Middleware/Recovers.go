package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
	"strings"
)

//统一的错误处理
func Recovers() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				DebugStack := ""
				for _, v := range strings.Split(string(debug.Stack()), "\n") {
					DebugStack += v + "\n\n"
				}
				fmt.Println("->>>", DebugStack)
				fmt.Println("->:", err)
				c.JSON(http.StatusOK, gin.H{
					"err": http.StatusInternalServerError,
					"msg": DebugStack,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
