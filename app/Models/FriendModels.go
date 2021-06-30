package Models

import (
	"go-websocket/tool/RdLine"
	"go-websocket/tool/Timer"
	"encoding/json"
	"fmt"
)

//好友申请的信息格式
type FriendRequestInfo struct {
	ForUserId int    `gorm:"column:forUserId" json:"forUserId"` //请求人的userid
	ForUId    int    `gorm:"column:forUid" json:"forUid"`       //请求人的uid
	ForMsg    string `gorm:"column:forMsg" json:"forMsg"`       //请求人的uid
	RequestTm int    `gorm:"column:requestTm" json:"requestTm"` //请求的时间
	Stat      int    `gorm:"column:stat" json:"stat"`           //0未处理1同意2拒绝
	Tm        string `gorm:"column:tm" json:"tm"`               //信息发送时间
	RedMsg    bool   `gorm:"column:redMsg" json:"redMsg"`       //0未查看1已查看
}

//好友申请信息key
func getFriendRequestInfoKey(uid int) string {
	return fmt.Sprintf("uFriend:%d", uid)
}

//获取好友申请列表
func GetFriendRequestList(uid int) []FriendRequestInfo {
	keys := getFriendRequestInfoKey(uid)
	RdLine := RdLine.GetRedisClient()
	defer RdLine.CloseRedisClient()
	vals, err := RdLine.StringMap(RdLine.Exec("HGETALL", keys))
	if err != nil {
		return nil
	}
	var list []FriendRequestInfo
	for _, s2 := range vals {
		var info FriendRequestInfo
		err = json.Unmarshal([]byte(s2), &info)
		if err == nil {
			list = append(list, info)
		}
	}
	return list
}

//发送好友请求
func SendFriendRequest(friendRequestInfo FriendRequestInfo, toUid int) {
	keys := getFriendRequestInfoKey(toUid)
	pipe := RdLine.GetPipeClient()
	defer pipe.ClosePipeClient()
	b, err := json.Marshal(friendRequestInfo)
	if err == nil {
		pipe.Add("hset", keys, Timer.NowUnix(), string(b))
		pipe.Add("Expire", keys, ttl*7)
		pipe.ExecPipe()
		pipe.RecvPipe()
	}
	return
}

//处理好友申请
func HandlerFriendRequest() {

}
