package redisutil

import (
	"fmt"
	"go-websocket/config"
	"sync"
	"testing"
	"time"
)

func TestMultLock(t *testing.T) {
	config.InitTest()
	InitRdLine()
	pool := GetRedisPool()
	// 创建 Redis 锁实例
	redisLock := NewRedisLock(pool, 10*time.Second)

	// 创建一个等待组
	var wg sync.WaitGroup

	// 启动 10 个 Goroutine
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 尝试获取锁
			lockKey := "mylock"
			locked, err := redisLock.TryLock(lockKey)
			if err != nil {
				fmt.Printf("协程 %d: 获取锁失败: %s\n", id, err)
				return
			}

			if locked {
				fmt.Printf("协程 %d: 锁获取成功\n", id)

				// 模拟业务操作
				time.Sleep(10 * time.Second)

				// 释放锁
				err := redisLock.Unlock(lockKey)
				if err != nil {
					fmt.Printf("协程 %d: 释放锁失败: %s\n", id, err)
					return
				}

				fmt.Printf("协程 %d: 锁释放成功\n", id)
			} else {
				fmt.Printf("协程 %d: 获取锁失败: 锁已被其他进程持有\n", id)
			}
		}(i)
	}

	// 等待所有 Goroutine 完成
	wg.Wait()
}
func TestLock(t *testing.T) {
	config.InitTest()
	InitRdLine()
	pool := GetRedisPool()
	// 创建 Redis 锁实例
	redisLock := NewRedisLock(pool, 10*time.Second)

	// 尝试获取锁
	lockKey := "mylock"
	locked, err := redisLock.TryLock(lockKey)
	if err != nil {
		fmt.Println("Failed to acquire lock:", err)
		return
	}

	if locked {
		fmt.Println("Lock acquired successfully")

		// 模拟业务操作
		time.Sleep(5 * time.Second)

		// 释放锁
		err := redisLock.Unlock(lockKey)
		if err != nil {
			fmt.Println("Failed to release lock:", err)
			return
		}

		fmt.Println("Lock released successfully")
	} else {
		fmt.Println("Failed to acquire lock: Lock is held by another process")
	}
}
func TestPipe(t *testing.T) {
	config.InitTest()
	InitRdLine()
	pipeClient := GetPipeClient()
	defer pipeClient.ClosePipeClient()

	// 使用 Send 方法将命令发送到管道中
	pipeClient.PipeAdd("SET", "keke", "12111111")
	pipeClient.PipeAdd("SET", "haha", "12111111")
	pipeClient.PipeAdd("expire", "keke", 600)
	pipeClient.PipeAdd("expire", "haha", 600)
	pipeClient.PipeAdd("GET", "keke")

	// 使用 Flush 方法将管道中的命令发送到 Redis 服务器执行
	if err := pipeClient.PipeExec(); err != nil {
		fmt.Println("发送管道命令错误：", err)
		return
	}

	// 使用 Receive 方法获取执行结果
	value, err := pipeClient.String(pipeClient.PipeRecv())
	if err != nil {
		fmt.Println("获取命令执行结果错误：", err)
		return
	}

	fmt.Println("命令执行结果：", value)
}
