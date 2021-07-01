package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Msg struct {
	MsgId int
	Err   int
	Data  interface{}
}

//获取信息头的长度
func GetHeadLen() int {
	return 4
}

// Encode 将消息编码len+{"id":123,"path":"/Chat/C2C","ver":"1.0.0","data":{"toUid":11,"toMsg":"你好啊"}}
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	// 写入消息头
	if err := binary.Write(pkg, binary.LittleEndian, length); err != nil {
		return nil, err
	}
	//写入消息id
	var msgId = int32(2147483646)
	if err := binary.Write(pkg, binary.LittleEndian, msgId); err != nil {
		return nil, err
	}
	// 写入消息实体
	if err := binary.Write(pkg, binary.LittleEndian, []byte(message)); err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

// Decode 解码消息
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(8) // 读取前4个字节的数据
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	if err := binary.Read(lengthBuff, binary.LittleEndian, &length); err != nil {
		return "", err
	}
	var msgId int32
	if err := binary.Read(lengthBuff, binary.LittleEndian, &msgId); err != nil {
		return "", err
	}
	fmt.Println("msgid:", msgId)
	// Buffered返回缓冲中现有的可读取的字节数。
	if int32(reader.Buffered()) < length+8 {
		return "", errors.New("Buffered返回缓冲中现有的可读取的字节数")
	}

	// 读取真正的消息数据
	pack := make([]byte, int(8+length))
	_, err := reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[8:]), nil
}
