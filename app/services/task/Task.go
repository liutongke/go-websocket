package task

import (
	"fmt"
	"go-websocket/app/services/websocket"
	"go-websocket/config"
	"runtime/debug"
	"time"
)

// 开始任务
func StartTask() {
	PingTimer(cleanConnection, "", 300*time.Second)
}

// 清理超时连接任务
func cleanConnection(param interface{}) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()

	if config.GetConf().WebSocket.CleanConnection {
		fmt.Println("定时任务，清理超时连接", param)
		websocket.ClearTimeoutConnections()
	}

	return
}
