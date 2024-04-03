package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-websocket/tools/jwt"
	"go-websocket/tools/timer"
	"net/http"
)

// Client 是用户连接管理器
type Client struct {
	Hub           *Hub
	Ws            *websocket.Conn // 客户端的用户连接
	Send          chan []byte     // 等待发送的数据
	UserId        int             // 用户Id，用户登录后才有
	OpenId        string          // openid
	FirstTime     uint64          // 首次连接时间
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间，登录后才有
}

// NewClient 初始化客户端
func NewClient(Hub *Hub, userId int, ws *websocket.Conn, firstTime uint64) *Client {
	client := &Client{
		Hub:           Hub,
		Ws:            ws,
		Send:          make(chan []byte, 100),
		UserId:        userId,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		LoginTime:     firstTime,
	}
	return client
}

// StartWs 启动 WebSocket
func (hub *Hub) StartWs(ctx *gin.Context) {
	upgrade := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 跨域请求，允许所有来源
		},
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")}, // 处理 Sec-WebSocket-Protocol 头部
	}

	// 检查并设置请求头部
	upgradeHeader := http.Header{}
	if hdr := ctx.Request.Header.Get("Sec-Websocket-Protocol"); hdr != "" {
		upgradeHeader.Set("Sec-Websocket-Protocol", hdr)
	}
	if hdr := ctx.Request.Header.Get("Set-Cookie"); hdr != "" {
		upgradeHeader.Set("Set-Cookie", hdr)
	}

	// 升级 HTTP 连接为 WebSocket
	conn, err := upgrade.Upgrade(ctx.Writer, ctx.Request, upgradeHeader)
	if err != nil {
		fmt.Printf("建立 WebSocket 连接失败: %v", err)
		return
	}

	// 获取并验证 token
	//token := ctx.GetHeader("nat-x-token")
	token := ctx.Param("token")

	claims, isValid := jwt.Verify(token)
	if !isValid {
		// 无效 token，返回错误消息并关闭连接
		response, _ := json.Marshal(Response{
			Id:   -1,
			Code: http.StatusBadRequest,
			Msg:  "无效的 token",
			Data: nil,
		})
		conn.WriteMessage(websocket.TextMessage, response)
		conn.Close()
		return
	}

	// 创建新的客户端并注册到 Hub
	client := NewClient(hub, claims.Uid, conn, uint64(timer.GetNowUnix()))
	client.Hub.Register <- client

	// 启动客户端读写协程
	go client.writePump()
	go client.readPump()
}
