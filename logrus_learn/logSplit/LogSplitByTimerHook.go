package logSplit

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type fileDateHook struct {
	file *os.File
	logPath string
	fileDate string // 判断日期切换目录
	appName string
}

func (hook fileDateHook) Levels() []logrus.Level{
	return logrus.AllLevels
}
func (hook fileDateHook) Fire(entry *logrus.Entry) error {
	timer := entry.Time.Format("2006-01-02_03-03")
	line, _ := entry.String()
	if hook.fileDate == timer {
		hook.file.Write([]byte(line))
		return nil
	}
	// 时间不等
	hook.file.Close()
  os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)
  filename := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.appName)

  hook.file, _ = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
  hook.fileDate = timer
  hook.file.Write([]byte(line))
  return nil
}

func initFile(logPath, appName string){
	fileDate := time.Now().Format("2006-01-02_15-04")
  //创建目录
  err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), os.ModePerm)
  if err != nil {
    logrus.Error(err)
    return
  }

  filename := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
  file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
  if err != nil {
    logrus.Error(err)
    return
  }
  fileHook := fileDateHook{file, logPath, fileDate, appName}
  logrus.AddHook(&fileHook)
}

func LogSplitByTimerHook() {
	initFile("logs","logrusLearn")
	for {
    logrus.Errorf("error")
    time.Sleep(20 * time.Second)
    logrus.Warnln("warning")
  }
}