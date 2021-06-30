package utils

import (
	"fmt"
	"runtime/debug"
	"time"
)

//开始任务
func StartTask() {
	//var d = 5 * time.Second
	PingTimer(cleanConnection, "", 300*time.Second)
}

// 清理超时连接
func cleanConnection(param interface{}) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()

	fmt.Println("定时任务，清理超时连接", param)
	//websocket.ClearTimeoutConnections()

	return
}
