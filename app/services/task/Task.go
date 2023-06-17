package task

import (
	"fmt"
	"go-websocket/app/services/bindCenter"
	"go-websocket/app/services/websocket"
	"go-websocket/config"
	"runtime/debug"
	"time"
)

// 开始任务
func StartTask() {
	//var d = 5 * time.Second
	PingTimer(cleanConnection, "", 300*time.Second)
	PingTimer(initServiceCenter, "", bindCenter.F5ServiceTm*time.Second)
	PingTimer(delServiceCenter, "", bindCenter.F5ServiceTm*3*time.Second)
}

// 清理超时连接任务
func cleanConnection(param interface{}) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ClearTimeoutConnections stop", r, string(debug.Stack()))
		}
	}()

	if config.GetConfClient().WebSocket.CleanConnection {
		fmt.Println("定时任务，清理超时连接", param)
		websocket.ClearTimeoutConnections()
	}

	return
}

// 服务中心注册
func initServiceCenter(param interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("服务器注册中心失败", r, string(debug.Stack()))
		}
	}()

	if config.GetConfClient().CommonConf.IsOpenRpc {
		//fmt.Println("定时任务，服务器注册中心", param)
		bindCenter.SetService()
	}

	return
}

// 清理超时的服务中心
func delServiceCenter(param interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("清理服务器注册中心失败", r, string(debug.Stack()))
		}
	}()
	if config.GetConfClient().CommonConf.IsOpenRpc {
		//fmt.Println("定时任务，清理服务器注册中心", param)
		bindCenter.DelTimeoutService()
	}
	return
}
