package Etcds

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
	"sync"
)

type EtcdRegister struct {
	Client        *clientv3.Client                        //etcd client
	LeaseID       clientv3.LeaseID                        //租约ID
	KeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse //租约keepalieve相应chan
	rwMutex       sync.RWMutex
	UserList      map[string]string
}

var etcdRegister *EtcdRegister

func GetEtcdRegister() *EtcdRegister {
	return etcdRegister
}
func NewEtcdRegister() *EtcdRegister {
	etcdRegister = &EtcdRegister{
		UserList: make(map[string]string),
	}
	return etcdRegister
}

type FunRegister func(*EtcdRegister) // 声明了一个函数类型

func (e *EtcdRegister) EtcdStartRegister(fun FunRegister) {
	// 创建 etcd 客户端连接

	client, err := clientv3.New(EtcdConfig())
	if err != nil {
		log.Println("Failed to create etcd client:", err)
		return
	}
	defer client.Close()

	// 为键设置租约，并获取租约 ID
	resp, err := client.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	e.Client = client
	e.LeaseID = resp.ID

	//fmt.Println("resp.ID:", resp.ID)
	// 将键与租约关联
	//_, err = client.Put(context.TODO(), "go-nat-x", "bar1", clientv3.WithLease(resp.ID))
	//if err != nil {
	//	log.Println("Failed to put key-value pair:", err)
	//	return
	//}

	// 续期键的租约
	keepAliveChan, err := client.KeepAlive(context.TODO(), resp.ID)
	if err != nil {
		log.Fatal(err)
	}
	e.KeepAliveChan = keepAliveChan
	go listen(keepAliveChan)
	//e.RegisterServer() //注册本机
	fun(e) //注册本机
	log.Println("注册服务启动成功")
	log.Println("租约lease ID:", resp.ID)
	select {}
}

func listen(keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse) {
	for {
		select {
		case leaseKeepResp := <-keepAliveChan:
			if leaseKeepResp == nil {
				log.Println("续约关闭")
			} else {
				// 续约成功
				//log.Println("续约成功")
			}
		}
	}
}

// 租期添加
func (e *EtcdRegister) PutKey(key, val string) (*clientv3.PutResponse, error) {
	e.rwMutex.Lock()         // 加写锁
	defer e.rwMutex.Unlock() // 释放写锁
	put, err := e.Client.Put(context.TODO(), key, val, clientv3.WithLease(e.LeaseID))
	if err != nil {
		log.Println("Failed to put key-value pair:", err)
		return put, err
	}
	return put, err
}
func (e *EtcdRegister) DelKey(key string) int {
	e.rwMutex.Lock()         // 加写锁
	defer e.rwMutex.Unlock() // 释放写锁
	resp, err := e.Client.Delete(context.TODO(), key)

	if err != nil {
		log.Println("Failed to put key-value pair:", err)
		return 0
	}
	return int(resp.Deleted)
}
