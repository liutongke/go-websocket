package etcds

import (
	"fmt"
	"go-websocket/config"
	"log"
	"testing"
	"time"
)

// 指定时间内获取不到锁，直接安徽
func TestLock(t *testing.T) {
	config.InitTest()
	locker, err := NewEtcdLocker("/lock/key")
	if err != nil {
		log.Fatal(err)
	}
	defer locker.Close()

	err = locker.AcquireLock()
	if err != nil {
		log.Fatal(err)
	}

	// 在这里执行需要保护的代码，即只有获得锁的客户端才能执行的部分
	fmt.Println("Lock acquired. Executing protected code...")

	// 模拟执行一些操作
	//time.Sleep(5 * time.Second)

	fmt.Println("Protected code executed. Lock released.")

	err = locker.ReleaseLock()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lock released.")
}

// 阻塞，一直等待到获取到锁
func TestLockWait(t *testing.T) {
	config.InitTest()

	// 初始化 etcd 客户端
	locker, err := NewEtcdWaitLock("/lock/key")
	if err != nil {
		log.Fatal(err)
	}
	defer locker.Close()

	// 尝试获取锁
	err = locker.AcquireLock()
	if err != nil {
		log.Fatal(err)
	}

	// 在这里执行需要保护的代码，即只有获得锁的客户端才能执行的部分
	fmt.Println("Lock acquired. Executing protected code...")

	// 模拟执行一些操作
	//time.Sleep(5 * time.Second)

	fmt.Println("Protected code executed. Lock released.")
	err = locker.ReleaseLock()
	if err != nil {
		log.Fatalf("释放锁错误:%v", err)
	}
}

// 不阻塞，获取不到立马返回
func TestTryLock(t *testing.T) {
	config.InitTest()
	locker, err := NewEtcdTryLock("/lock/key")
	if err != nil {
		log.Fatal(err)
	}
	defer locker.Close()

	err = locker.AcquireLock()
	if err != nil {
		log.Fatalf("加锁失败立即返回:%v\n", err)
		return
	}

	log.Println("加锁成功开始执行业务")
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Second)
		log.Printf("执行 %%%d ...", i*10)
	}

	err = locker.ReleaseLock()
	if err != nil {
		log.Fatalf("释放锁失败:%v\n", err)
		return
	}
	log.Println("释放锁完成")
}
