package pack

import "encoding/json"

//消息体长度
type MsgData struct {
	MsgId int
	Err   int
	Data  interface{}
}

type Msg struct {
	dataLen int32  //消息头长度
	data    []byte //消息体内容
}

func NewMsgData(data MsgData) *Msg {
	dataByte, _ := json.Marshal(data)
	return &Msg{
		dataLen: int32(len(dataByte)),
		data:    dataByte,
	}
}

//获取消息头长度
func (m *Msg) GetDataLen() int32 {
	return m.dataLen
}

//获取消息体内容
func (m *Msg) GetMsgData() MsgData {
	var MsgBodyData MsgData
	err := json.Unmarshal(m.data, &MsgBodyData)
	if err != nil {
		return MsgBodyData
	}
	return MsgBodyData
}
