package Etcds

import (
	"context"
	"fmt"
	"go-websocket/tools"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// 心跳检测
func startHeartbeat() {
	client, err := clientv3.New(EtcdConfig())
	if err != nil {
		tools.EchoErrorExit(fmt.Sprintf("Failed to create etcd client:%v", err))
		return
	}
	defer client.Close()

	// 使用Get请求检查连接可用性
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = client.Get(ctx, "test-key")
	cancel()

	if err != nil {
		tools.EchoErrorExit(fmt.Sprintf("Etcd server connection failed, server terminated:%v", err))
	}
}
