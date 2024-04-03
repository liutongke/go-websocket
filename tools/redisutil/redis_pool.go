package redisutil

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-websocket/config"
	"go-websocket/tools"
	"time"
)

var (
	RedisPool *redis.Pool
)

func InitRdLine() {
	RedisPool = initRedisPool()
}

func initRedisPool() *redis.Pool {
	redisConfig := config.GetConf()

	return &redis.Pool{
		MaxIdle:     redisConfig.Redis.MaxIdle,                   //最大闲置连接数量
		MaxActive:   redisConfig.Redis.MaxActive,                 //最大活动连接数
		IdleTimeout: redisConfig.Redis.IdleTimeout * time.Minute, //闲置过期时间 在get函数中会有逻辑 删除过期的连接
		Wait:        redisConfig.Redis.Wait,                      //设置如果活动连接达到上限 再获取时候是等待还是返回错误 如果是false 系统会返回redigo: connection pool exhausted
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConfig.Redis.Addr, redisConfig.Redis.Port))
			if err != nil {
				tools.EchoErrorExit(fmt.Sprintf("Failed to connect to the redis server: %v", err))
			}
			if len(redisConfig.Redis.Password) > 0 { //判断下是否有密码
				if _, authErr := c.Do("AUTH", redisConfig.Redis.Password); authErr != nil {
					c.Close()
					tools.EchoErrorExit(fmt.Sprintf("Login redis authentication failed: %v", authErr))
				}
			}
			if _, selectErr := c.Do("SELECT", redisConfig.Redis.DB); selectErr != nil {
				c.Close()
				tools.EchoErrorExit(fmt.Sprintf("Failed to select redis database: %v", selectErr))
			}
			return c, nil
		},
	}
}

func GetRedisPool() *redis.Pool {
	if RedisPool == nil {
		tools.EchoError("Redis is not initialized")
	}
	return RedisPool
}

// 从连接池中获取一个连接
// conn := pool.Get()
// defer conn.Close()
//
// // 开始管道
// conn.Send("MULTI")
//
// // 将多个 Redis 命令添加到管道中
// conn.Send("SET", "key1", "value1")
// conn.Send("GET", "key1")
// conn.Send("HSET", "hashkey", "field1", "value1")
//
// // 提交管道并执行命令
// results, err := redis.Values(conn.Do("EXEC"))
// if err != nil {
// fmt.Println("管道执行错误：", err)
// return
// }
//
// // 解析结果
// var val1, val2, val3 string
// _, err = redis.Scan(results, &val1, &val2, &val3)
// if err != nil {
// fmt.Println("结果解析错误：", err)
// return
// }
//
// fmt.Println("key1 的值：", val1)
// fmt.Println("hashkey 的 field1 值：", val3)
// 开启管道
func GetPipeClient() *RedisClient {
	return GetRedisClient()
}

// 关闭管道
func (r *RedisClient) ClosePipeClient() {
	r.CloseRedisClient()
}

// 管道添加
func (r RedisClient) PipeAdd(commandName string, args ...interface{}) {
	r.Client.Send(commandName, args...)
}

// 发送管道命令
func (r RedisClient) PipeExec() error {
	return r.Client.Flush()
}

// 接收管道结果
func (r RedisClient) PipeRecv() (reply interface{}, err error) {
	return r.Client.Receive()
}

func GetRedisClient() *RedisClient {
	conn := RedisPool.Get() //获取一个连接
	if conn.Err() != nil {  //连接获取失败
		tools.EchoError(fmt.Sprintf("get Redis Pool Client Error: %v", conn.Err()))
	}
	return &RedisClient{conn}
}

type RedisClient struct {
	Client redis.Conn
}

// 回收这个连接
func (r *RedisClient) CloseRedisClient() {
	r.Client.Close() //回收这个连接
}

// 为redis-go Do操作入口
func (r *RedisClient) Exec(cmd string, args ...interface{}) (interface{}, error) {
	return r.Client.Do(cmd, args...)
}

// bool 类型转换
func (r *RedisClient) Bool(reply interface{}, err error) (bool, error) {
	return redis.Bool(reply, err)
}

// string 类型转换
func (r *RedisClient) String(reply interface{}, err error) (string, error) {
	return redis.String(reply, err)
}

// strings 类型转换
func (r *RedisClient) Strings(reply interface{}, err error) ([]string, error) {
	return redis.Strings(reply, err)
}

// Float64 类型转换
func (r *RedisClient) Float64(reply interface{}, err error) (float64, error) {
	return redis.Float64(reply, err)
}

// int 类型转换
func (r *RedisClient) Int(reply interface{}, err error) (int, error) {
	return redis.Int(reply, err)
}

// int64 类型转换
func (r *RedisClient) Int64(reply interface{}, err error) (int64, error) {
	return redis.Int64(reply, err)
}

// uint64 类型转换
func (r *RedisClient) Uint64(reply interface{}, err error) (uint64, error) {
	return redis.Uint64(reply, err)
}

// Bytes 类型转换
func (r *RedisClient) Bytes(reply interface{}, err error) ([]byte, error) {
	return redis.Bytes(reply, err)
}

// StringMap 类型转换
func (r *RedisClient) StringMap(reply interface{}, err error) (map[string]string, error) {
	return redis.StringMap(reply, err)
}
