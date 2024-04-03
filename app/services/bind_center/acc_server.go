package bind_center

import (
	"encoding/json"
	"fmt"
	"go-websocket/config"
	"go-websocket/tools/redisutil"
	"go-websocket/tools/utils"
)

const keys = "uBindServ"

type BindUserInfo struct {
	RpcAddr string //最后一次登录的服务器的rpc链接地址
	UserId  int    // 用户Id，用户登录以后才有
}

// GetBindInfo 获取用户的绑定信息
func GetBindInfo(userId int) BindUserInfo {
	RdLine := redisutil.GetRedisClient()
	defer RdLine.CloseRedisClient()
	bindInfo, err := RdLine.Bytes(RdLine.Exec("hGet", keys, userId))
	var bindUserInfo BindUserInfo
	if err != nil {
		return bindUserInfo
	}
	json.Unmarshal(bindInfo, &bindUserInfo)
	return bindUserInfo
}

// BindUidAndService 将用户与应用服务器地址绑定
func BindUidAndService(userId int) bool {
	b, err := json.Marshal(BindUserInfo{
		RpcAddr: fmt.Sprintf("%s:%s", utils.GetLocalIp(), config.GetConf().Grpc.RpcPort),
		UserId:  userId,
	})
	if err == nil {
		RdLine := redisutil.GetRedisClient()
		defer RdLine.CloseRedisClient()
		RdLine.Exec("hSet", keys, userId, string(b))
		RdLine.Exec("EXPIRE", keys, 84600)
		return true
	}
	return false
}

// DelBindUidAndService 退出登录的时候删掉防止占用内容太大
func DelBindUidAndService(userId int) {
	RdLine := redisutil.GetRedisClient()
	defer RdLine.CloseRedisClient()
	RdLine.Exec("hDel", keys, userId)
}
