package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// 4a0000000a352e372e3236000c0000002075546b780e273500fff7c00200ff8115000000000000000000004971082517414a186634754b006d7973716c5f6e61746976655f70617373776f726400
func Handshakes() {
	hexStringData := "0a352e372e3236000c0000002075546b780e273500fff7c00200ff8115000000000000000000004971082517414a186634754b006d7973716c5f6e61746976655f70617373776f726400"
	packet, _ := hex.DecodeString(hexStringData)

	protocolPacket := []byte{00, 00}
	copy(protocolPacket, packet[0:1])
	protocolVersion = binary.LittleEndian.Uint16(protocolPacket)
	fmt.Printf("protocolVersion:%d\n", protocolVersion)

	var dbVer []byte
	var idx int
	for k, item := range packet[1:] {
		dbVer = append(dbVer, item)
		if item == 0 {
			idx = k
			goto next
		}
	}
next:
	fmt.Printf("serverVersion:%s\n", string(dbVer))
	idx = idx + 2

	fmt.Printf("threadId:%d\n", binary.LittleEndian.Uint32(packet[idx:idx+4]))

	fmt.Printf("salt:%s\n", string(packet[idx+4:idx+4+8]))
	fmt.Println("salt1 len:", len(packet[idx+4:idx+4+8]))
	fmt.Printf("serverCapabilities:%d\n", binary.LittleEndian.Uint16(packet[idx+4+8+1:idx+4+8+1+2]))

	languagePacket := []byte{00, 00}
	copy(languagePacket, packet[idx+4+8+1+2:idx+4+8+1+2+1])
	fmt.Printf("server Language:%d\n", binary.LittleEndian.Uint16(append(languagePacket, 00)))

	fmt.Printf("server Status:%d\n", binary.LittleEndian.Uint16(packet[idx+4+8+1+2+1:idx+4+8+1+2+1+2]))

	fmt.Printf("Extended Server Capabilities:%d\n", binary.LittleEndian.Uint16(packet[idx+4+8+1+2+1+2:idx+4+8+1+2+1+2+2]))

	pluginLengthPacket := []byte{00, 00}
	copy(pluginLengthPacket, packet[idx+4+8+1+2+1+2+2:idx+4+8+1+2+1+2+2+1])
	fmt.Printf("plugin Length:%d\n", binary.LittleEndian.Uint16(pluginLengthPacket))

	fmt.Printf("Unused:%s\n", string(packet[idx+4+8+1+2+1+2+2+1:idx+4+8+1+2+1+2+2+1+10]))

	var salt2 []byte
	for _, saltIem := range packet[idx+4+8+1+2+1+2+2+1+10:] {
		if saltIem == 0 {
			goto salt2jump
		}
		salt2 = append(salt2, saltIem)
	}

salt2jump:
	fmt.Println("salt2 len:", len(salt2))
	fmt.Printf("salt2:%s\n", string(salt2))
	fmt.Printf("Authentication Plugin:%s\n", string(packet[idx+4+8+1+2+1+2+2+1+10+len(salt2):]))
	//fmt.Println("Iq\b%\u0017AJ\u0018f4uK")
	//fmt.Println(" uTkx\u000E'5")
	//fmt.Println("0x81ff")

}

var (
	protocolVersion    uint16 //版本协议
	serverVersion      string //版本号
	threadId           uint32 //执行的线程号
	public             string //用于后期加密的salt1
	serverCapabilities uint16 //通信的协议
	serverCharsetIndex uint16 //编码格式
	serverStatus       uint16 //服务端的状态
	restOfScrambleBuff string //这个其实就是seed2
)

func main() {
	//fmt.Printf("serverCapabilities:%d\n", binary.LittleEndian.Uint16(p))
	Handshakes()
	//header := []byte{74, 74, 0, 0}
	//length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	//fmt.Println(uint32(header[1]))
	//fmt.Println("++++++")
	// 转换的用的 byte数据
	//byte_data := []byte(`测试数据`)
	//fmt.Println(byte_data)
	//// 将 byte 装换为 16进制的字符串
	//hex_string_data := hex.EncodeToString(byte_data)
	//// byte 转 16进制 的结果
	//println(hex_string_data)

	//var a byte
	//a = 126 //00011110
	//fmt.Printf("%d二进制为:%v\n", a, biu.ToBinaryString(a))
	//
	//fmt.Println("------")
	/* ====== 分割线 ====== */
	// 将 16进制的字符串 转换 byte

	//header := []byte{74, 00, 00, 0}
	//
	//length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	//fmt.Println(length)
	//hex_data, _ := hex.DecodeString(hexStringData)
	//fmt.Println(hex_data)
	//t := hex_data[:3]
	//t = append(t, 00)
	//fmt.Println("++++++++", binary.LittleEndian.Uint32(t), "++++++")
	//fmt.Println("------")
	//fmt.Println(biu.BytesToBinaryString(hex_data[:3]))
	//fmt.Println("------")
	//var b int32
	//biu.ReadBinaryString(biu.BytesToBinaryString(hex_data[:3]), &b)
	//fmt.Println(b) //259
	//len := hex_data[:3]
	//fmt.Printf("%08b\n", len>>1)
	//fmt.Println(binary.LittleEndian.Uint32(hex_data[:4]))
	//fmt.Printf("%d", hex_data[:3])
	//fmt.Printf("%d", []byte("74"))
	//fmt.Println(binary.LittleEndian.Uint32(hex_data[3:4]))

	// 将 byte 转换 为字符串 输出结果
	//println(string(hex_data))
}
