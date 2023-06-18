package Etcds

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"sync"
)

type EtcdDiscovery struct {
	Client        *clientv3.Client                        //etcd client
	LeaseID       clientv3.LeaseID                        //租约ID
	KeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse //租约keepalieve相应chan
	rwMutex       sync.RWMutex
	UserList      map[string]string
	WatchEvents   map[string]FunDiscovery
}
type FunDiscovery func(*clientv3.Event) // 声明了一个函数类型

var etcdDiscovery *EtcdDiscovery

func GetEtcdDiscovery() *EtcdDiscovery {
	return etcdDiscovery
}
func NewEtcdDiscovery(fun map[string]FunDiscovery) *EtcdDiscovery {
	etcdDiscovery = &EtcdDiscovery{
		UserList:    make(map[string]string),
		WatchEvents: fun,
	}
	return etcdDiscovery
}

func (e *EtcdDiscovery) EtcdStartDiscovery(keyPrefixes []string) {
	// 创建 Etcd 客户端
	client, err := clientv3.New(EtcdConfig())
	if err != nil {
		log.Println("Failed to create etcd client:", err)
		return
	}
	e.Client = client
	defer client.Close()

	// 监听某个键的变化
	//keyPrefixes := []string{"/prefix1", "/prefix2", "/net", "go-nat-x", ETCD_SERVER_LIST, ETCD_PREFIX_ACCOUNT_INFO}

	// 创建一个用于取消监听的context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 监听多个键前缀的所有值
	for _, prefix := range keyPrefixes {
		go e.watchPrefix(ctx, client, prefix)
	}
	log.Println("发现服务启动成功")

	select {} // 阻塞主线程，等待程序退出
}

func (e *EtcdDiscovery) watchPrefix(ctx context.Context, client *clientv3.Client, prefix string) {

	watcher := clientv3.NewWatcher(client)                  // 创建一个Watcher
	watchKeyPrefix := clientv3.WithPrefix()                 // 设置要监听的前缀
	watchChan := watcher.Watch(ctx, prefix, watchKeyPrefix) // 启动Watch

	// 处理Watch事件
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			if event.Type == clientv3.EventTypePut {
				e.WatchEvents["put"](event)
			} else if event.Type == clientv3.EventTypeDelete {
				e.WatchEvents["del"](event)
			}
		}
	}
}
