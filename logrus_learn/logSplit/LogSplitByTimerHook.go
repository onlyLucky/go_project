package logSplit

import (
	"fmt"
	"os"

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
	os.MkdirAll(fmt.Sprintf(""))
}

func LogSplitByTimerHook() {}