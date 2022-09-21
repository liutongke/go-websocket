package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
)

func read(reader *bufio.Reader) []byte {
	line, err := reader.ReadBytes('\n')
	if err != nil {
		errors.New("读取行错误")
	}
	return bytes.TrimRight(line, "\r\n")
}

var Redis1Conn net.Conn

// TCP server端
func main() {
	Redis1Conn, _ = net.Dial("tcp", "192.168.0.105:6379") //监听Redis服务器

	listen, err := net.Listen("tcp", "0.0.0.0:12345") //监听本地端口
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func process(conn net.Conn) {
	defer func() {
		fmt.Println("连接断开了\n")
		conn.Close() // 关闭连接
	}()

	for {

		cmd := handlerCliCmd(conn) //处理命令
		//fmt.Println(cmd)
		//conn.Write([]byte("+ok\r\n"))

		//str := strings.Join(cmd, "\r\n")
		//fmt.Println("------>",str)
		//str := "*3\r\n" +
		//	"$3\r\n" +
		//	"set\r\n" +
		//	"$4\r\n" +
		//	"name\r\n" +
		//	"$5\r\n" +
		//	"pdudo\r\n"
		fmt.Println("to Reids Server cmd:--->", cmd)
		if cmd[2] == "COMMAND" {
			conn.Write([]byte("+OK\r\n"))
		} else {
			var toRedis []byte
			for _, cmds := range cmd {
				//fmt.Println(cmds)
				t := cmds + "\r\n"
				toRedis = append(toRedis, []byte(t)...)
			}
			//fmt.Println(str)
			n1, err := Redis1Conn.Write(toRedis)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("n1", n1)
			buf := make([]byte, 1024)
			n2, err2 := Redis1Conn.Read(buf[:])

			if err2 != nil {
				fmt.Println(err)
			}
			fmt.Println("n2", n2)

			r := buf[:n2]
			fmt.Println("Redis Server back--->", string(r))

			conn.Write(r)
		}
	}
}

// 处理redis-cli发送的程序
func handlerCliCmd(conn net.Conn) []string {
	var cmd []string
	reader := bufio.NewReader(conn)

	cmdHead := read(reader)

	cmd = append(cmd, string(cmdHead))

	if string(cmdHead[:1]) == "*" {

		cmdArrByteLen := cmdHead[1:] //发送的命令数组长度

		cmdArrLen := BytesToInt(cmdArrByteLen)

		for i := 0; i < cmdArrLen; i++ {
			cmdLen := read(reader)
			cmd = append(cmd, string(cmdLen))
			if string(cmdLen[:1]) == "$" {
				cmdData := read(reader)
				cmd = append(cmd, string(cmdData))
			}
		}
	}
	return cmd
}

//COMMAND redis-cli连接的命令
//*3
//$3
//set
//$4
//name
//$5
//pdudo

func BytesToInt(b []byte) int {
	s := string(b)
	num, _ := strconv.Atoi(s)
	return num
}
