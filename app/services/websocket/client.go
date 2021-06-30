package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-websocket/utils/Jwt"
	"net/http"
	"time"
)

//单个用户管理
// 用户连接
type Client struct {
	Hub           *Hub
	Ws            *websocket.Conn // 客户端的用户连接
	Send          chan []byte     // 等待发送的数据
	UserId        int             // 用户Id，用户登录以后才有
	OpenId        string          //openid
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

// 初始化
func NewClient(Hub *Hub, userId int, openId string, ws *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Hub:           Hub,
		Ws:            ws,
		Send:          make(chan []byte, 100),
		UserId:        userId,
		OpenId:        openId,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		LoginTime:     firstTime,
	}
	return client
}

//启动websocket
func (hub *Hub) StartWs(ctx *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},                                                               //跨域true忽略
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")}, // 处理 Sec-WebSocket-Protocol Header
	}
	upgradeHeader := http.Header{}
	if hdr := ctx.Request.Header.Get("Sec-Websocket-Protocol"); hdr != "" {
		upgradeHeader.Set("Sec-Websocket-Protocol", hdr)
	}
	if hdr := ctx.Request.Header.Get("Set-Cookie"); hdr != "" {
		upgradeHeader.Set("Set-Cookie", hdr)
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, upgradeHeader)
	if err != nil {
		fmt.Println("建立websocket连接失败:", err)
		return
	}
	token := ctx.Param("token")
	e, Claims := Jwt.Verify(token)
	if !e {
		b, _ := json.Marshal(Response{
			Id:   -1,
			Err:  http.StatusBadRequest,
			Msg:  "无效的token",
			Data: nil,
		})
		conn.WriteMessage(websocket.TextMessage, b)
		conn.Close()
		return
	}

	client := NewClient(hub, Claims.UserId, Claims.Openid, conn, uint64(time.Now().Unix()))
	client.Hub.Register <- client

	go client.writePump() //发送客户端信息
	go client.readPump()  //读取客户端信息
}
