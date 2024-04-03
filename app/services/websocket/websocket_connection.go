package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
	"time"
)

const (
	writeWait      = 10 * time.Second    // 定义了允许向对等方写入消息的时间。在这个时间窗口内，应用程序可以向对等方发送消息。超过这个时间后，写入操作将被中断。
	pongWait       = 60 * time.Second    // 定义了读取下一个Pong消息的允许时间。如果在这个时间内没有接收到Pong消息，连接将被关闭。
	pingPeriod     = (pongWait * 9) / 10 // 定义了发送Ping消息给对等方的周期。Ping消息用于检测连接的活跃状态，以确保连接保持活跃。该值必须小于pongWait，通常设置为pongWait的一部分。
	maxMessageSize = 512                 // 定义了从对等方接收的消息的最大大小限制。如果接收到的消息超过这个限制，可能会导致连接被中断或消息被截断
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
