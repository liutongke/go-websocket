package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "/etcd_server_list/192.168.1.106:go-websocket"

	// 定义匹配IP地址的正则表达式
	ipRegex := `(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`

	// 创建正则表达式对象
	re := regexp.MustCompile(ipRegex)

	// 查找匹配的IP地址
	ipMatches := re.FindStringSubmatch(str)

	// 提取第一个匹配的IP地址
	if len(ipMatches) > 1 {
		ip := ipMatches[1]
		fmt.Println("Extracted IP address:", ip)
	} else {
		fmt.Println("No IP address found")
	}
}
