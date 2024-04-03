package etcds

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
)

type LockWait struct {
	Client  *clientv3.Client
	Session *concurrency.Session
	Mutex   *concurrency.Mutex
}

// 初始化 etcd 客户端
func NewEtcdWaitLock(lockKey string) (*LockWait, error) {
	client, err := clientv3.New(EtcdConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to create etcd client: %v", err)
	}
	// 创建一个 etcd 分布式锁会话
	session, err := concurrency.NewSession(client)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %v", err)
	}

	mutex := concurrency.NewMutex(session, lockKey)

	return &LockWait{
		Client:  client,
		Session: session,
		Mutex:   mutex,
	}, nil
}

// 获取分布式锁
func (l *LockWait) AcquireLock() error {
	ctx := context.TODO()
	err := l.Mutex.Lock(ctx)
	if err != nil {
		return fmt.Errorf("failed to acquire lock: %v", err)
	}

	return nil
}

// 释放分布式锁
func (l *LockWait) ReleaseLock() error {
	ctx := context.TODO()
	err := l.Mutex.Unlock(ctx)
	if err != nil {
		return fmt.Errorf("failed to release lock: %v", err)
	}

	return nil
}

func (l *LockWait) Close() {
	err := l.Session.Close()
	if err != nil {
		log.Fatalf("client关闭失败:%v\n", err)
	}
	err = l.Client.Close()
	if err != nil {
		log.Fatalf("Session关闭失败:%v\n", err)
	}
}
