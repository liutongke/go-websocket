package socket

import (
	"bufio"
	"fmt"
	"go-websocket/config"
	"go-websocket/tools/Timer"
	"net"
	"strconv"
)

// 开始tcp连接
func StartTcp() {
	go func() {
		listen, err := net.Listen("tcp", "0.0.0.0:"+config.GetConf().Server.TcpPort)
		if err != nil {
			fmt.Println(fmt.Sprintf("listen failed, err:%v", err))
			return
		}
		for {
			conn, err := listen.Accept() // 建立连接
			if err != nil {
				fmt.Println(fmt.Sprintf("accept failed, err:%v", err))
				continue
			}
			client := NewTcpClient(1, "1", conn, uint64(Timer.GetNowUnix()))
			go client.writePump() //发送客户端信息
			go client.readPump()  //读取客户端信息
		}
	}()
}

// 写入消息
func (t *TcpClient) writePump() {
	for {
		select {
		case message, ok := <-t.Send:
			fmt.Println("写入处理", ok)
			t.SendMsg(message) // 发送数据
		}
	}
}

func (t *TcpClient) SendMsg(data []byte) {
	t.Conn.Write(data)
}

// 读取消息
func (t *TcpClient) readPump() {
	for {
		reader := bufio.NewReader(t.Conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		t.Send <- []byte("服务器收到了你的消息了" + strconv.FormatInt(Timer.GetNowUnix(), 10))
	}
}

// 用户连接管理
type TcpClient struct {
	Conn          net.Conn    // 客户端的连接
	Send          chan []byte // 等待发送的数据
	UserId        int         // 用户Id，用户登录以后才有
	OpenId        string      //openid
	FirstTime     uint64      // 首次连接事件
	HeartbeatTime uint64      // 用户上次心跳时间
	LoginTime     uint64      // 登录时间 登录以后才有
}

// 初始化
func NewTcpClient(userId int, openId string, conn net.Conn, firstTime uint64) (client *TcpClient) {
	client = &TcpClient{
		Conn:          conn,
		Send:          make(chan []byte, 100),
		UserId:        userId,
		OpenId:        openId,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		LoginTime:     firstTime,
	}
	return client
}

func MsgHandle(msgType int, msg []byte) {}
