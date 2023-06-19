package tools

import (
	"fmt"
	"github.com/yinheli/qqwry"
	"go-websocket/config"
	"go-websocket/tools/Dir"
	"net"
)

// IpToAddr https://pkg.go.dev/github.com/yinheli/qqwry@v0.0.0-20160229183603-f50680010f4a#section-readme
// 使用ip地址获得用户城市 qqwry 不是线程安全的 qqwry 没有使用缓存
func IpToAddr(ipStr string) (*qqwry.QQwry, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, fmt.Errorf("%s is not a valid IP address", ipStr)
	}

	q := qqwry.NewQQwry(Dir.GetAbsDirPath(config.GetConf().CommonConf.ChunzhenIP))
	q.Find(ipStr)
	return q, nil
}

//addr, _ := tools.IpToAddr("114.83.73.202")
//log.Println("ip地址:", addr.Ip, addr.City, addr.Country)
