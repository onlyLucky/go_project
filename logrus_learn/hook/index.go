package hook

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

/*
logrus在记录levels()返回的日志级别的消息时会触发HOOK
按照Fire方法定义的内容修改logrus.Entry。
*/
type MyHook struct{
	Writer *os.File
}
// 设置一个field
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	hook.Writer.Write([]byte(line))
	entry.Data["app"] = "fengfeng"
	return nil
}

// 哪些等级的日志才会生效
func (hook *MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}


// logrus最令人心动的功能就是其可扩展的HOOK机制了，通过在初始化时为logrus添加hook，logrus可以实现各种扩展功能。
func HookFunc(){
	// 日志的打开格式是追加，所以不能用os.Create
  logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
  logrus.SetReportCaller(true)
  file, _ := os.OpenFile("err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
  hook := &MyHook{Writer: file}
  logrus.AddHook(hook)
  logrus.Errorf("hello")
}