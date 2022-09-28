package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
)

//{protocolVersion=10, serverVersion='5.7.13', threadId=4055, scramble=[49, 97, 80, 3, 35, 118, 45, 15, 5, 118, 9, 11, 124, 93, 93, 5, 31, 47, 111, 109, 0, 0, 0, 0, 0], serverCapabilities=65535, serverLanguage=33, serverStatus=2}

func Handshake() []byte {
	var byteLen = 1 + len("5.7") + 1 + 4 + 9 + 2 + 1 + 2 + 13 + 13
	idx := 0
	data := make([]byte, byteLen)

	data[idx] = uint8(10)

	idx++
	copy(data[idx:], "5.7")

	idx += len("5.7")

	data[idx] = uint8(0)
	idx++

	binary.LittleEndian.PutUint32(data[idx:], uint32(4500))
	idx += 4
	binary.LittleEndian.PutUint64(data[idx:], uint64(12345678))
	idx += 8
	data[idx] = uint8(0)
	idx += 1
	binary.LittleEndian.PutUint16(data[idx:], uint16(33))
	idx += 2
	data[idx] = uint8(33)
	idx += 1
	binary.LittleEndian.PutUint16(data[idx:], uint16(10))
	idx += 2
	copy(data[idx:], "1234567890123")
	idx += 13
	copy(data[idx:], "1234567890120")
	return data
	//fmt.Printf("->%d", data, len(data), byteLen, string(data[0:1]))
	//fmt.Printf("%d%d", data[0:1], data[1:4])
	//fmt.Println(uint8(data[0:1]), string(data[1:4]))
	//fmt.Println(string(data))
	//data[]
	//var byteLen = 1 + len("5.7") + 1 + 4 + 9 + 2 + 1 + 2 + 13 + 13
	//fmt.Println("包长度：", byteLen)

	//var buf = make([]byte, byteLen)
	//
	//var index = 0
	//buf[index] = uint8(10)
	//fmt.Println(index)
	////verLen := len("5.7") + 1
	//fmt.Println(string(buf))
	//copy(buf[index+1:], "5.7null")
	//fmt.Println(string(buf))
	//var buf []byte = make([]byte, byteLen)
	//buf := new(bytes.Buffer)
	//binary.Write(buf, binary.LittleEndian, uint8(10))
	//binary.Write(buf, binary.LittleEndian, "5.7")
	//binary.Write(buf, binary.LittleEndian, uint8(00))
	//binary.Write(buf, binary.LittleEndian, uint32(4055))
	//
	//binary.Read()
	//binary.Write(buf, binary.LittleEndian, uint8(10))
	//binary.Write(buf, binary.LittleEndian, uint8(0))
	//binary.Write(buf, binary.LittleEndian, uint8(0))
	//binary.LittleEndian.PutUint16(HandshakeBytes, uint8(10))
	//binary.LittleEndian.PutUint8(HandshakeBytes, uint32(testInt))
	//fmt.Println("int32 to bytes:", HandshakeBytes)
	//
	//convInt := binary.LittleEndian.Uint32(HandshakeBytes)
	//fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func register() {
	//user := ""
	//pass := ""
	//hostname, _ := os.Hostname()
	//fmt.Println(hostname)
	//data := make([]byte, 4+1+4+1+len(hostname)+1+len(user)+1+len(pass)+2+4+4)
	//pos := 4
	//data[pos] = 21 //register slave  command
	//pos++
	//binary.LittleEndian.PutUint32(data[pos:], uint32(1111))
	//pos += 4
	//
	//data[pos] = uint8(len(hostname))
	//pos++
	//n := copy(data[pos:], hostname)
	//pos += n
	//
	//data[pos] = uint8(len(user))
	//pos++
	//n = copy(data[pos:], user)
	//pos += n
	//
	//data[pos] = uint8(len(pass))
	//pos++
	//n = copy(data[pos:], pass)
	//pos += n
	//
	//binary.LittleEndian.PutUint16(data[pos:], uint16(3306))
	//pos += 2
	//
	//binary.LittleEndian.PutUint32(data[pos:], 0)
	//pos += 4
	//
	////master id = 0
	//binary.LittleEndian.PutUint32(data[pos:], 0)
	//
	//s.io.writePacket(data)
	//
	////ok
	//res, _ := s.io.readPacket()
	//if res[0] == OK_HEADER {
	//	fmt.Println("register success.")
	//	s.registerSucc = true
	//} else {
	//	s.io.HandleError(data)
	//}
}

func main() {
	//zhuanhuan()
	sendData := Handshake()
	//fmt.Printf("%d", sendData)
	conn, err := net.Dial("tcp", "192.168.0.107:3306")
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}
	_, err = conn.Write(sendData)
	if err != nil {
		fmt.Println("conn write err:", err)
	}

	//fmt.Println("nnnnnnnnn->", write)

	for {
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			//fmt.Println("recv failed, err:", err)
			return
		}
		byteData := buf[:n]
		// 将 byte 装换为 16进制的字符串
		hexStringData := hex.EncodeToString(byteData)
		// byte 转 16进制 的结果
		println("----")
		println(string(hexStringData))
		println(string(byteData))

		// 将 16进制的字符串 转换 byte
		hex_data, _ := hex.DecodeString(hexStringData)
		// 将 byte 转换 为字符串 输出结果
		println(string(hex_data))
		//HandleError(t)
		//fmt.Println(t)
		//fmt.Printf("%d---%s", t, string(t[1:]))
		//fmt.Println(string(t))

		// 将 byte 装换为 16进制的字符串
		//hex_string_data := hex.EncodeToString(t)
		// byte 转 16进制 的结果
		//println(hex_string_data)

	}

	//
	////读入输入的信息
	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	data, err := reader.ReadString('\n')
	//	if err != nil {
	//		fmt.Printf("read from console failed, err:%v\n", err)
	//		break
	//	}
	//
	//	data = strings.TrimSpace(data)
	//	//传输数据到服务端
	//	_, err = conn.Write([]byte(data))
	//	if err != nil {
	//		fmt.Printf("write failed, err:%v\n", err)
	//		break
	//	}
	//}
}

func HandleError(data []byte) {
	pos := 1
	code := binary.LittleEndian.Uint16(data[pos:])
	pos += 2
	pos++
	state := string(data[pos : pos+5])
	pos += 5
	msg := string(data[pos:])
	fmt.Printf("code:%d, state:%s, msg:%s\n", code, state, msg)
}
func zhuanhuan() {
	var v int64 = 0 //默认10进制
	//s2 := strconv.FormatInt(v, 2) //10 转2进制
	//fmt.Printf("%v\n", s2)
	//
	//s8 := strconv.FormatInt(v, 8)
	//fmt.Printf("%v\n", s8)
	//
	//s10 := strconv.FormatInt(v, 10)
	//fmt.Printf("%v\n", s10)

	s16 := strconv.FormatInt(v, 16) //10 yo 16
	fmt.Printf("%v\n", s16)

	//var sv = "11"
	//fmt.Println(strconv.ParseInt(sv, 16, 32)) // 16 to 10
	//fmt.Println(strconv.ParseInt(sv, 10, 32)) // 10 to 10
	//fmt.Println(strconv.ParseInt(sv, 8, 32))  // 8 to 10
	//fmt.Println(strconv.ParseInt(sv, 2, 32))  // 2 to 10

}
