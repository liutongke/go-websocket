package Etcds

import (
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"testing"
	"time"
)

// 阻塞，一直等待到获取到锁
func TestLock(t *testing.T) {
	config.InitTest()
	lock := NewEtcdLock()
	// 初始化 etcd 客户端
	client, session := lock.InitializeEtcdClient()
	if client == nil || session == nil {
		log.Println("初始化失败")
		return
	}
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer func(client *clientv3.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("client关闭失败:%v\n", err)
		}
	}(client)
	defer func(session *concurrency.Session) {
		err := session.Close()
		if err != nil {
			log.Fatalf("Session关闭失败:%v\n", err)
		}
	}(session)
	// 创建一个 etcd 分布式锁会话
	//session, err := concurrency.NewSession(client)
	//if err != nil {
	//	log.Fatal("Failed to create session:", err)
	//}
	//defer session.Close()

	// 创建一个分布式锁
	lockKey := "/lock/key"
	mutex := concurrency.NewMutex(session, lockKey)

	// 尝试获取锁
	err := lock.AcquireLock(mutex)
	if err != nil {
		log.Fatal(err)
	}
	defer lock.ReleaseLock(mutex)

	// 在这里执行需要保护的代码，即只有获得锁的客户端才能执行的部分

	fmt.Println("Lock acquired. Executing protected code...")

	// 模拟执行一些操作
	//time.Sleep(5 * time.Second)

	fmt.Println("Protected code executed. Lock released.")
}

// 不阻塞，获取不到立马返回
func TestTryLock(t *testing.T) {
	config.InitTest()
	tryLock := NewEtcdTryLock()
	client, session := tryLock.InitializeClient()
	if client == nil || session == nil {
		log.Println("初始化失败")
		return
	}
	defer client.Close()
	defer func(session *concurrency.Session) {
		err := session.Close()
		if err != nil {
			log.Fatalf("Session关闭失败:%v\n", err)
		}
	}(session)

	mutex := concurrency.NewMutex(session, "my-lock")

	err := tryLock.AcquireLock(mutex)
	if err != nil {
		log.Fatalf("加锁失败立即返回:%v\n", err)
		return
	}

	log.Println("加锁成功开始执行业务")
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		log.Printf("执行 %%%d ...", i*10)
	}

	err = tryLock.ReleaseLock(mutex)
	if err != nil {
		log.Fatalf("释放锁失败:%v\n", err)
		return
	}
	log.Println("释放锁完成")
}
