package Etcds

import (
	"context"
	"go-websocket/tools"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"time"
)

// 心跳检测
func startHeartbeat() {
	client, err := clientv3.New(EtcdConfig())
	if err != nil {
		log.Println("Failed to create etcd client:", err)
		return
	}
	defer client.Close()

	// 使用Get请求检查连接可用性
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err = client.Get(ctx, "test-key")
	cancel()

	if err != nil {
		//log.Println("Etcd服务器连接失败:", err)
		tools.EchoError("Etcd服务器连接失败,服务器终止运行")
		os.Exit(1)
	} else {
		log.Println("Etcd连接可用")
	}
}
