package websocket

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"go-websocket/tools/logger"
	"go-websocket/tools/timer"
	"go-websocket/tools/utils"
	"net/http"
	"reflect"
	"runtime"
	"sync"
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
	if utils.IsDebug() {
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
	Code int         `json:"code"` // 返回的错误码
	Msg  string      `json:"msg"`  // 返回的信息
	Data interface{} `json:"data"` // 返回数据json
}

// 统一处理下信息
func MsgHandle(c *Client, msgType int, msg []byte) {
	data := Request{}
	err := json.Unmarshal(msg, &data)
	if err != nil {
		fmt.Printf("处理数据 json Unmarshal%v", err)
		d, _ := json.Marshal(Response{
			Id:   0,
			Code: http.StatusBadRequest,
			Msg:  "数据不合法",
			Data: nil,
		})
		c.SendMsg(d)

		return
	}

	if len(data.Ver) <= 0 {
		d, _ := json.Marshal(Response{
			Id:   0,
			Code: http.StatusBadRequest,
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
		if utils.IsDebug() {
			logInfo := "[GIN-ws] " + timer.GetNowStr() + "-->userId: " + cast.ToString(c.UserId) + "-->Request data:" + string(msg) + "\n\t-->Response body: " + string(d)
			fmt.Println(logInfo)
			logger.Info(logInfo)
		}
	} else {
		d, _ = json.Marshal(Response{
			Id:   0,
			Code: http.StatusBadRequest,
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
