package etcds

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type TryLock struct {
	Client  *clientv3.Client
	Session *concurrency.Session
	Mutex   *concurrency.Mutex
}

// 初始化客户端
func NewEtcdTryLock(lockKey string) (*TryLock, error) {
	//log.Println("客户端初始化")
	client, err := clientv3.New(EtcdConfig())
	if err != nil {
		return nil, fmt.Errorf("客户端初始化失败:%v\n", err)
	}

	// 创建一个session并设置默认租期30s，即锁默认超过30s会自动释放(内部会自动续期Etcd KeepAlive)
	//log.Println("Session初始化")
	session, err := concurrency.NewSession(client, concurrency.WithTTL(30))
	if err != nil {
		//log.Fatalf("Session初始化失败:%v\n", err)
		return nil, fmt.Errorf("Session初始化失败:%v\n", err)
	}
	mutex := concurrency.NewMutex(session, lockKey)
	return &TryLock{
		Client:  client,
		Session: session,
		Mutex:   mutex,
	}, nil
}

// 获取锁
func (t *TryLock) AcquireLock() error {
	//log.Println("TryLock加锁失败不会等待")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	err := t.Mutex.TryLock(ctx)
	if err != nil {
		return err
	}

	return nil
}

// 释放锁
func (t *TryLock) ReleaseLock() error {
	err := t.Mutex.Unlock(context.TODO())
	if err != nil {
		return err
	}

	return nil
}
func (t *TryLock) Close() {
	err := t.Client.Close()
	if err != nil {
		fmt.Println("client close err")
	}
	err = t.Session.Close()
	if err != nil {
		fmt.Println("Session close err")
	}
}
