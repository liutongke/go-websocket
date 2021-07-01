package test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go-websocket/app/services/proto"
	"net/http"
	"testing"
)

func TestStart(t *testing.T) {
	b, _ := json.Marshal(proto.Msg{
		MsgId: 1,
		Err:   http.StatusOK,
		Data:  "// Decode 解码消息\nfunc Decode(reader *bufio.Reader) (string, error) {\n\t// 读取消息的长度\n\tlengthByte, _ := reader.Peek(4) // 读取前4个字节的数据\n\tlengthBuff := bytes.NewBuffer(lengthByte)\n\tvar length int32\n\terr := binary.Read(lengthBuff, binary.LittleEndian, &length)\n\tif err != nil {\n\t\treturn \"\", err\n\t}\n\t// Buffered返回缓冲中现有的可读取的字节数。",
	})
	data, _ := proto.Encode(string(b))
	fmt.Println(data)
	byteData := bytes.NewReader(data)
	reader := bufio.NewReader(byteData)
	fmt.Println(proto.Decode(reader))
}
