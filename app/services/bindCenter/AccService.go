package bindCenter

import (
	"fmt"
	"go-websocket/tools/RdLine"
	"go-websocket/tools/Timer"
)

const (
	serviceKey    = "bind:service" // 全部的服务器
	serversKeyTtl = 86400          // key过期时间
	//ServiceTimeout = 60             // 超时时间
	F5ServiceTm = 10 //刷新写入时间
)

func getServiceKey() string {
	//return serviceKey
	return fmt.Sprintf("%s", serviceKey)
}

// 将服务器信息存入redis
func SetService() {
	pipe := RdLine.GetPipeClient()
	defer pipe.ClosePipeClient()
	pipe.Add("ZADD", getServiceKey(), Timer.GetNowUnix(), GetServiceToStr())
	pipe.Add("Expire", getServiceKey(), serversKeyTtl)
	pipe.ExecPipe()
	pipe.RecvPipe()
}

// 清除超时的服务
func DelTimeoutService() {
	RdLine := RdLine.GetRedisClient()
	defer RdLine.CloseRedisClient()
	RdLine.Exec("ZREMRANGEBYSCORE", getServiceKey(), 0, Timer.GetOffsetUnix(-(F5ServiceTm * 2)))
}
