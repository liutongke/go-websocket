package Logger

import (
	"bufio"
	"fmt"
	"go-websocket/utils/Dir"
	"go-websocket/utils/Timer"
	"os"
	"runtime"
)

type Logger struct {
	LogInfoChan chan *LogInfo //管道通知日志
}

type LogInfo struct {
	FileLevel string
	Msg       string
}

var logsClient *Logger

func InitLogger() {
	logsClient = newLogger()
	go logsClient.runLogs()
}

func newLogger() *Logger {
	return &Logger{
		LogInfoChan: make(chan *LogInfo, 10),
	}
}

func (l *Logger) runLogs() {
	for {
		select {
		case logInfo := <-l.LogInfoChan: //记录登录用户
			l.log(logInfo.FileLevel, logInfo.Msg)
		}
	}
}

//写入日志信息2021-01-08 12:09:22|DEBUG|REQ:App.Auth.Index|
func (fl *Logger) log(saveName string, Level string) {
	isDebug := false
	if isDebug { //判断是否是线下开发模式
		return
	}
	file, err := os.OpenFile(Dir.GetAbsolutePath("/runtime/"+saveName+"_"+Timer.DateId()+".log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	//fileName, filePath, line := getRunInfo(1)
	//log := getNowStrTm() + "|" + Level + "|" + saveName + "\n" + "|line:" + strconv.Itoa(line) + "|func name:" + fileName + "|file:" + filePath //将要打印的信息
	writer := bufio.NewWriter(file)
	writer.WriteString(Level + "\n") //将数据先写入缓存
	writer.Flush()
}

//skip相对当前代码的调用层级 获取运行的行号、文件名称、调用的函数名称
func getRunInfo(skip int) (fileName string, file string, line int) {
	funcName, file, line, ok := runtime.Caller(skip)
	if ok {
		fmt.Println("func name: " + runtime.FuncForPC(funcName).Name())
		fmt.Printf("file: %s, line: %d\n", file, line)
		funcName := runtime.FuncForPC(funcName).Name()
		return funcName, file, line
	}
	return
	//return "", "", 0
}

func Info(msg string) {
	loginfo := &LogInfo{
		FileLevel: "info",
		Msg:       msg,
	}
	logsClient.LogInfoChan <- loginfo
}

func Debug(msg string) {
	loginfo := &LogInfo{
		FileLevel: "Debug",
		Msg:       msg,
	}
	logsClient.LogInfoChan <- loginfo
}

func Err(msg string) {
	loginfo := &LogInfo{
		FileLevel: "Err",
		Msg:       msg,
	}
	logsClient.LogInfoChan <- loginfo
}

func Warning(msg string) {
	loginfo := &LogInfo{
		FileLevel: "Warning",
		Msg:       msg,
	}
	logsClient.LogInfoChan <- loginfo
}

func Fatal(msg string) {
	loginfo := &LogInfo{
		FileLevel: "Fatal",
		Msg:       msg,
	}
	logsClient.LogInfoChan <- loginfo
}
