package Tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	Dir "go-websocket/tools/Dir"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"
)

// GetRequestIP 获取客户端ip
func GetRequestIP(c *gin.Context) string {
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	return reqIP
}

// 获取本机IP地址
func GetLocalIp() string {
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

// 获取服务器地址
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

// 生成count个[start,end)结束的不重复的随机数
func GenerateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

//生成count个[start,end)结束的不重复的随机数字符串类型
//func GenerateRandomStr(start int, end int, count int) []string {
//范围检查
//if end < start || (end-start) < count {
//	return nil
//}
//
////存放结果的slice
//nums := make([]int, 0)
////随机数生成器，加入时间戳保证每次生成的随机数不一样
//r := rand.New(rand.NewSource(time.Now().UnixNano()))
//for len(nums) < count {
//	//生成随机数
//	num := r.Intn((end - start)) + start
//
//	//查重
//	exist := false
//	for _, v := range nums {
//		if v == num {
//			exist = true
//			break
//		}
//	}
//
//	if !exist {
//		nums = append(nums, num)
//	}
//}
//var strList []string
//for _, v := range nums {
//	str, _ := IntToStr(v)
//	strList = append(strList, str)
//}
//return strList
//}

// 获取文件后缀
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

// 通过url获取字节流
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
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// 获取map所有的值
func ArrayKeys(elements map[string]int) []interface{} {
	i, keys := 0, make([]interface{}, len(elements))
	for key, _ := range elements {
		keys[i] = key
		i++
	}
	return keys
}

/*
*
@保存图片到指定文件夹
@file 要保存的文件字节流
@Dir 文件将要保存的文件夹
return 文件名称
return 文件的sid，唯一值
*/
func SaveImg(file []byte, dir string) (string, string) {
	dirPath, _ := os.Getwd()
	savePath := dirPath + "/" + dir
	_, err := os.Stat(savePath)
	if err != nil { //文件夹不存在，创建
		os.Mkdir(dir, 777)
	}
	var newFileName string
	suffixName := GetImgExt(file) //获取图片后缀
	uuidName := uuid.NewString()
	newFileName = savePath + "/" + uuidName + suffixName
	ioutil.WriteFile(newFileName, file, 0644)
	return uuidName + suffixName, uuid.NewString()
}

// 初始化创建配置文件
func InitMkdir(dir string) string {
	dirPath, _ := os.Getwd()
	savePath := dirPath + "/" + dir
	_, err := os.Stat(savePath)
	if err != nil { //文件夹不存在，创建
		os.Mkdir(dir, 777)
		return savePath
	}
	return ""
}

// true开发模式，false生产模式
func IsDebug() bool {
	sysType := runtime.GOOS
	if sysType == "linux" {
		return false // LINUX系统
	}
	if sysType == "windows" {
		return true // windows系统
	}
	return false
}

// 过期时间
func Ttl(t int) time.Duration {
	return time.Second * time.Duration(t)
}

// 打印
func DD(msg ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Println(file, line, "=====>")
	for k, v := range msg {
		fmt.Println("key:", k, "--->", "vls:", v)
	}
}

// 读取配表
func ReadConfTb(name string, ver string) []byte {
	filePath := Dir.GetAbsDirPath("./bs/v1/tb/" + name) // 打开json文件
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer jsonFile.Close() // 要记得关闭
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil
	}
	return byteValue
}
