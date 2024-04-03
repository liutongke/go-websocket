package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-websocket/tools/fileutil"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 可以包含在随机字符串中的字符集
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))] // 从字符集中随机选择字符
	}
	return string(result)
}

// GenerateRandomNumber 生成指定范围内的随机整数
func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

// BytesToMegabytes 字节转MB
func BytesToMegabytes(bytes float64) float64 {
	megabytes := bytes / (1024 * 1024)
	return megabytes
}

// BytesToGigabytes 字节转GB
func BytesToGigabytes(bytes float64) float64 {
	gigabytes := bytes / (1024 * 1024 * 1024)
	return gigabytes
}

// GetRequestIP 获取客户端ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

// GetLocalIp 获取本机IP地址
func GetLocalIp() string {
	if os.Getenv("DOCKER_IN") == "1" { //容器内
		return os.Getenv("MY_IP")
	}
	ip, err := GetOutBoundIP()
	if err != nil {
		return ""
	}
	return ip
}

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	//fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// GetServIp 获取服务器地址
func GetServIp() string {
	ip, err := externalIP()
	if err != nil {
		return ""
	}
	return ip.String()
}

func externalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, err
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

// GetImgExt 获取文件后缀
func GetImgExt(headerByte []byte) (ext string) {
	xStr := fmt.Sprintf("%x", headerByte)
	switch {
	case xStr == "89504e470d0a1a0a":
		ext = ".png"
	case xStr == "0000010001002020":
		ext = ".ico"
	case xStr == "0000020001002020":
		ext = ".cur"
	case xStr[:12] == "474946383961" || xStr[:12] == "474946383761":
		ext = ".gif"
	case xStr[:10] == "0000020000" || xStr[:10] == "0000100000":
		ext = ".tga"
	case xStr[:8] == "464f524d":
		ext = ".iff"
	case xStr[:8] == "52494646":
		ext = ".ani"
	case xStr[:4] == "4d4d" || xStr[:4] == "4949":
		ext = ".tiff"
	case xStr[:4] == "424d":
		ext = ".bmp"
	case xStr[:4] == "ffd8":
		ext = ".jpg"
	case xStr[:2] == "0a":
		ext = ".pcx"
	default:
		ext = ""
	}
	return ext
}

// GetUrlToByte 通过url获取字节流
func GetUrlToByte(url string) ([]byte, error) {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// ArrayKeys 获取map所有的值
func ArrayKeys(elements map[string]int) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key, _ := range elements {
		keys[i] = key
		i++
	}
	return keys
}

// CreateDirectoryIfNotExist 创建文件夹
func CreateDirectoryIfNotExist(dir string) (string, error) {
	var savePath string

	if filepath.IsAbs(dir) {
		savePath = dir
	} else {
		// 如果 dir 是相对路径，则将其与当前工作目录连接
		dirPath, err := os.Getwd()
		if err != nil {
			return "", err
		}
		savePath = filepath.Join(dirPath, dir)
	}

	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		if err := os.Mkdir(dir, 0777); err != nil {
			return "", fmt.Errorf("failed to create directory: %v", err)
		}
		return savePath, nil
	}

	return "", nil
}

// IsDebug true开发模式，false生产模式
func IsDebug() bool {
	//return runtime.GOOS == "windows"
	return os.Getenv("DEBUG") == "1"
}

// IsRsapi true 树莓派 false 非树莓派
func IsRsapi() bool {
	return GetCPUModel() == "aarch64"
}
func GetCPUModel() string {
	return os.Getenv("CPU_MODEL")
}

// Ttl 过期时间
func Ttl(t int) time.Duration {
	return time.Second * time.Duration(t)
}

// DD 打印
func DD(msg ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(file, line, "=====>")
	for k, v := range msg {
		fmt.Println("key:", k, "--->", "vls:", v)
	}
}

// ReadConfTb 读取配表
func ReadConfTb(name string, ver string) []byte {
	filePath := fileutil.GetAbsDirPath("./bs/v1/tb/" + name) // 打开json文件
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer jsonFile.Close() // 要记得关闭
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil
	}
	return byteValue
}
