package logger

import (
	"bufio"
	"fmt"
	"go-websocket/config"
	"go-websocket/tools"
	"go-websocket/tools/fileutil"
	"go-websocket/tools/timer"
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

// 写入日志信息2021-01-08 12:09:22|DEBUG|REQ:App.Auth.Index|
func (fl *Logger) log(saveName string, Level string) {
	logPath := fmt.Sprintf("%s/logger_%s_%s.log", fileutil.GetAbsDirPath(config.GetConf().Logger.LogFolder), saveName, timer.GetDateId())

	//fileutil.GetAbsDirPath(config.GetConf().Logger.LogFolder+"logger_"+saveName+"_"+timer.GetDateId()+".log")
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		tools.EchoError(fmt.Sprintf("open logger file failed err: %v", err))
	}
	defer file.Close()
	//fileName, filePath, line := getRunInfo(1)
	//log := timer.GetNowStr() + "|" + Level + "|" + saveName + "\n" + "|line:" + strconv.Itoa(line) + "|func name:" + fileName + "|file:" + filePath //将要打印的信息
	log := fmt.Sprintf("[%s] | %s | %s \n", timer.GetNowStr(), saveName, Level)
	writer := bufio.NewWriter(file)
	//writer.WriteString(Level + "\n") //将数据先写入缓存
	writer.WriteString(log) //将数据先写入缓存
	writer.Flush()
}

// skip相对当前代码的调用层级 获取运行的行号、文件名称、调用的函数名称
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
