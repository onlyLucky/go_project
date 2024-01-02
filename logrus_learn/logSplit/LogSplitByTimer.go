package logSplit

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// 日志分割

// logFormatter 日志自定义格式
type logFormatter struct{}

// Format 格式详情
func (s *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	fmt.Println("Format")
	timestamp := time.Now().Local().Format("2006-01-01 00:00:00")
	var file string
	var len int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
	}
	msg := fmt.Sprintf("[%s] %s [%s:%d] %s\n",strings.ToUpper(entry.Level.String()),timestamp,file,len, entry.Message)
	return []byte(msg), nil
}

type logFileWriter struct {
	file *os.File
	logPath string
	fileDate string // 判断日期切换目录
	appName string
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0,errors.New("file not opened")
	}
	// 判断是否需要切换日期
	fileDate := time.Now().Format("2006-01-02")
	if p.fileDate != fileDate {
		p.file.Close()
		err = os.MkdirAll(fmt.Sprintf("%s/%s",p.logPath,fileDate), os.ModePerm)
		if err != nil {
			return 0, err
		}
		filename := fmt.Sprintf("%s/%s/%s-%s.log",p.logPath,fileDate,p.appName,fileDate)

		p.file, err = os.OpenFile(filename,os.O_WRONLY|os.O_APPEND|os.O_CREATE,0600)
		if err != nil {
			return 0, err
		}
	}
	n, e := p.file.Write(data)
	return n, e
}

// initLog 初始化日志
func initLog(logPath string, appName string){
	fileDate := time.Now().Format("20060102")
	// 创建目录
	err := os.MkdirAll(fmt.Sprintf("%s/%s",logPath, fileDate), os.ModePerm)
	if err != nil {
		log.Error(err)
		return
	}
	filename := fmt.Sprintf("%s/%s/%s-%s.log",logPath, fileDate, appName, fileDate)
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Error(err)
		return
	}
	fileWriter := logFileWriter{file, logPath, fileDate, appName}
	log.SetOutput(os.Stdout)
	writers := []io.Writer{
		&fileWriter,
		os.Stdout,
	}
	// 同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	if err == nil {
		log.SetOutput(fileAndStdoutWriter)
	} else {
		log.Info("failed to log to file")
	}
	log.SetReportCaller(true)
	log.SetFormatter(new(logFormatter))

}

// 按时间分割 自定义write方法
func LogSplitByTimer(){
	initLog("logs","logrusLearn")
	log.Infoln("hello world")
}
