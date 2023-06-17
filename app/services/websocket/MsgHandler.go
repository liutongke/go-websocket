package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/spf13/cast"
	"go-websocket/tools/Logger"
	"go-websocket/tools/Timer"
	"go-websocket/tools/Tools"
	"net/http"
	"reflect"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

type DisposeFunc func(c *Client, msg string) Response

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
)

func GetFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// 注册
func Register(key string, value DisposeFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value
	if Tools.IsDebug() {
		fmt.Println("[GIN-debug] ws  " + key + "-->" + GetFunctionName(value))
	}
	return
}

func GetHandlers(key string) (value DisposeFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()

	value, ok = handlers[key]

	return
}

// 请求的结构体
type Request struct {
	Id   int         `json:"id"`   //消息id
	Ver  string      `json:"ver"`  //版本号
	Path string      `json:"path"` // 请求命令字
	Data interface{} `json:"data"` // 数据 json
}

// 返回的结构体
type Response struct {
	Id   int         `json:"id"`   //消息id
	Err  int         `json:"err"`  // 返回的错误码
	Msg  string      `json:"msg"`  // 返回的信息
	Data interface{} `json:"data"` // 返回数据json
}

// 统一处理下信息
func MsgHandle(c *Client, msgType int, msg []byte) {
	data := Request{}
	err := json.Unmarshal(msg, &data)
	if err != nil {
		fmt.Println(fmt.Sprintf("处理数据 json Unmarshal%v", err))
		d, _ := json.Marshal(Response{
			Id:   0,
			Err:  http.StatusBadRequest,
			Msg:  "数据不合法",
			Data: nil,
		})
		c.SendMsg(d)

		return
	}

	if len(data.Ver) <= 0 {
		d, _ := json.Marshal(Response{
			Id:   0,
			Err:  http.StatusBadRequest,
			Msg:  "ver empty",
			Data: nil,
		})
		c.SendMsg(d)

		return
	}

	var d []byte
	if f, ok := GetHandlers(data.Path); ok {
		v := f(c, string(msg))
		v.Id = data.Id //消息的唯一的id
		d, _ = json.Marshal(v)
		if Tools.IsDebug() {
			logInfo := "[GIN-ws] " + Timer.GetNowStr() + "-->userId: " + cast.ToString(c.UserId) + "-->Request data:" + string(msg) + "\n\t-->Response body: " + string(d)
			fmt.Println(logInfo)
			Logger.Info(logInfo)
		}
	} else {
		d, _ = json.Marshal(Response{
			Id:   0,
			Err:  http.StatusBadRequest,
			Msg:  "路由错误",
			Data: nil,
		})
	}
	c.SendMsg(d)
}

// 向对应的客户端发送信息
func (c *Client) SendMsg(msg []byte) {
	if c == nil {
		return
	}
	c.Send <- msg
}

const (
	writeWait      = 10 * time.Second    // 允许向对等方写入消息的时间.
	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
	maxMessageSize = 512                 // Maximum message size allowed from peer.
)

var (
	newline  = []byte{'\n'}
	space    = []byte{' '}
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// 读取客户端消息
func (c *Client) readPump() {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println("readPump err : " + string(debug.Stack()))
		}
	}()

	defer func() {
		c.Hub.Unregister <- c
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		msgType, message, err := c.Ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(fmt.Sprintf("读取客户端数据错误或者客户端主动关闭连接 error: %v", err))
			}
			break
		}
		MsgHandle(c, msgType, message) //路由处理客户的消息
	}
}

// 发送消息给客户端
func (c *Client) writePump() {
	defer func() {
		// recover() 可以将捕获到的panic信息打印
		if err := recover(); err != nil {
			fmt.Println("writePump err :" + string(debug.Stack()))
		}
	}()

	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.Ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
