package websocket

import "go-websocket/config"

// 用户心跳
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime

	return
}

// 心跳超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+config.GetConf().WebSocket.HeartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}
