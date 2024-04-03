package redisutil

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisLock struct {
	pool    *redis.Pool
	timeout time.Duration
}

func NewRedisLock(pool *redis.Pool, timeout time.Duration) *RedisLock {
	return &RedisLock{pool: pool, timeout: timeout}
}

func (rl *RedisLock) TryLock(key string) (bool, error) {
	conn := rl.pool.Get()
	defer conn.Close()

	// 设置锁
	result, err := redis.String(conn.Do("SET", key, "1", "EX", int(rl.timeout.Seconds()), "NX"))
	if err != nil {
		return false, err
	}

	// 判断是否获取到锁
	if result == "OK" {
		return true, nil
	}

	return false, nil
}

func (rl *RedisLock) Unlock(key string) error {
	conn := rl.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

//	func main() {
//		// 创建 Redis 连接池
//		pool := &redis.Pool{
//			MaxIdle:     5,
//			MaxActive:   10,
//			IdleTimeout: 30 * time.Second,
//			Dial: func() (redis.Conn, error) {
//				return redis.Dial("tcp", "localhost:6379")
//			},
//		}
//
//		// 创建 Redis 锁实例
//		redisLock := NewRedisLock(pool, 10*time.Second)
//
//		// 尝试获取锁
//		lockKey := "mylock"
//		locked, err := redisLock.TryLock(lockKey)
//		if err != nil {
//			fmt.Println("Failed to acquire lock:", err)
//			return
//		}
//
//		if locked {
//			fmt.Println("Lock acquired successfully")
//
//			// 模拟业务操作
//			time.Sleep(5 * time.Second)
//
//			// 释放锁
//			err := redisLock.Unlock(lockKey)
//			if err != nil {
//				fmt.Println("Failed to release lock:", err)
//				return
//			}
//
//			fmt.Println("Lock released successfully")
//		} else {
//			fmt.Println("Failed to acquire lock: Lock is held by another process")
//		}
//	}
//func main() {
//	// 创建 Redis 连接池
//	pool := &redis.Pool{
//		MaxIdle:     5,
//		MaxActive:   10,
//		IdleTimeout: 30 * time.Second,
//		Dial: func() (redis.Conn, error) {
//			return redis.Dial("tcp", "localhost:6379")
//		},
//	}
//
//	// 创建 Redis 锁实例
//	redisLock := NewRedisLock(pool, 10*time.Second)
//
//	// 创建一个等待组
//	var wg sync.WaitGroup
//
//	// 启动 10 个 Goroutine
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(id int) {
//			defer wg.Done()
//
//			// 尝试获取锁
//			lockKey := "mylock"
//			locked, err := redisLock.TryLock(lockKey)
//			if err != nil {
//				fmt.Printf("Goroutine %d: Failed to acquire lock: %s\n", id, err)
//				return
//			}
//
//			if locked {
//				fmt.Printf("Goroutine %d: Lock acquired successfully\n", id)
//
//				// 模拟业务操作
//				time.Sleep(1 * time.Second)
//
//				// 释放锁
//				err := redisLock.Unlock(lockKey)
//				if err != nil {
//					fmt.Printf("Goroutine %d: Failed to release lock: %s\n", id, err)
//					return
//				}
//
//				fmt.Printf("Goroutine %d: Lock released successfully\n", id)
//			} else {
//				fmt.Printf("Goroutine %d: Failed to acquire lock: Lock is held by another process\n", id)
//			}
//		}(i)
//	}
//
//	// 等待所有 Goroutine 完成
//	wg.Wait()
//}
