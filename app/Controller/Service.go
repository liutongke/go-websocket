package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"go-websocket/app/services/grpc"
	"go-websocket/app/services/websocket"
	"go-websocket/config"
	"go-websocket/tools/DbLine"
	"go-websocket/tools/RdLine"
	"go-websocket/tools/Tools"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

// 发送全服信息
func SendGlobalServMsg(c *gin.Context) {
	data := make(map[string]interface{}, 10)
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["uuid"] = uuid.NewString()
	websocket.SendMsgALl(data)
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendGlobalServMsg success",
		"data": data,
	})
	return
}

// 给单个用户发送消息
func SendUserMsg(c *gin.Context) {
	data := make(map[string]interface{})
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["uuid"] = uuid.NewString()
	userId := cast.ToInt(c.Param("userId"))
	websocket.SendUserMsg(userId, data)
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendUserMsg success",
		"data": data,
	})
	return
}

// 发送分组消息
func SendGroupMsg(c *gin.Context) {
	data := make(map[string]interface{})
	data["tm"] = time.Now().Format("2006-01-02 15:04:05")
	data["uuid"] = uuid.NewString()
	groupId := cast.ToInt(c.Param("groupId"))
	websocket.SendMsgToGroup(groupId, data)
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendGroupMsg success",
		"data": data,
	})
	return
}

// 查询系统状态
func Status(c *gin.Context) {
	grpc.SendUserMsg(fmt.Sprintf("%s:8972", Tools.GetLocalIp()), 1231, []byte("1111111111111111111"))
	data := make(map[string]interface{})
	NumGoroutine := runtime.NumGoroutine()
	NumCPU := runtime.NumCPU()
	data["NumGoroutine"] = NumGoroutine // goroutine数量
	data["NumCPU"] = NumCPU
	sqlDB, _ := DbLine.GetMysqlClient().DB()
	mysqlStatus := sqlDB.Stats() // 获取连接池统计信息

	pool := RdLine.GetRedisPool()
	redisStatus := pool.Stats()

	data["NumGoroutine"] = NumGoroutine // goroutine数量
	data["NumCPU"] = NumCPU
	data["configToml"] = config.GetConf()
	data["Redis status"] = map[string]int{
		"ActiveCount 表示连接池中的活动连接数，包括正在使用的连接和空闲的连接":           redisStatus.ActiveCount,
		"IdleCount 表示连接池中的空闲连接数":                             redisStatus.IdleCount,
		"WaitCount 表示等待获取连接的总数，即连接池中所有连接被占用，需要等待获取连接的请求数量":   int(redisStatus.WaitCount),
		"WaitDuration 表示等待获取连接的总时间，即所有等待获取连接的请求在连接池中等待的时间总和": int(redisStatus.WaitDuration),
	}
	data["MySQL status"] = map[string]string{
		"连接池允许的最大打开连接数:":   strconv.FormatInt(mysqlStatus.MaxLifetimeClosed, 10),
		"当前打开的连接数:":        strconv.Itoa(mysqlStatus.OpenConnections),
		"当前正在使用的连接数":       strconv.Itoa(mysqlStatus.InUse),
		"当前空闲的连接数:":        strconv.Itoa(mysqlStatus.Idle),
		"等待连接的请求数:":        strconv.FormatInt(mysqlStatus.WaitCount, 10),
		"等待连接的总时长:":        strconv.FormatInt(int64(mysqlStatus.WaitDuration), 10),
		"连接池关闭的最大空闲连接数:":   strconv.FormatInt(mysqlStatus.MaxIdleClosed, 10),
		"连接池关闭的最大连接生命周期数:": strconv.FormatInt(mysqlStatus.MaxLifetimeClosed, 10)}
	c.JSON(http.StatusOK, gin.H{
		"err":  http.StatusOK,
		"msg":  "SendGroupMsg success",
		"data": data,
	})
	return
}
