package common

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func CommonFunc() {

	/* 
	输出到日志文件
	*/
	/* file, _ := os.Create("info.log")
	logrus.SetOutput(file) */
	/* 
	同时输出屏幕和文件
	*/
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	//  设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File	
	file, _ := os.Create("info.log")
	writers := []io.Writer{
		file,
		os.Stdout,
	}
	//  同时写文件和屏幕
	fileAndStdoutWriter := io.MultiWriter(writers...)
	logrus.SetOutput(fileAndStdoutWriter)
	/* 
	logrus常用方法
	输出如下：
	time="2024-01-01T21:06:33+08:00" level=info msg=Infoln
	time="2024-01-01T21:06:33+08:00" level=warning msg=Warnln
	time="2024-01-01T21:06:33+08:00" level=error msg=Errorln
	time="2024-01-01T21:06:33+08:00" level=info msg=Println
	debug的没有输出，是因为logrus默认的日志输出等级是 info

	日志等级
	PanicLevel  // 会抛一个异常
	FatalLevel  // 打印日志之后就会退出
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel  // 低级别
	*/
	logrus.Debugln("Debugln")
	logrus.Infoln("Infoln")
	logrus.Warnln("Warnln")
	logrus.Errorln("Errorln")
	logrus.Println("Println")

	// 查看日志等级
	fmt.Println(logrus.GetLevel()) //info

	/* 
	更改日志等级
	*/
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debugln("Debugln")

	/* 
	设置特定字段
	*/
	log1 := logrus.WithField("project", "study")
	log1.Errorln("hello")
	// time="2022-12-17T15:02:28+08:00" level=error msg=hello project=study
	log2 := logrus.WithFields(logrus.Fields{
		"func": "main",
	})
	log2.Warningf("你好")
	// time="2022-12-17T15:02:28+08:00" level=warning msg="你好" func=main
	log3 := log2.WithFields(logrus.Fields{
		"auth": "枫枫",
	})
	// time="2022-12-17T15:02:28+08:00" level=warning msg="你好" auth="枫枫" func=main
	log3.Warnln("你好")

	/* 
	显示样式 Text和Json
	*/

	logrus.SetFormatter(&logrus.JSONFormatter{})
	log4 := logrus.WithField("project", "study")
	log4.Errorln("hello")

	/* 
	自定义颜色
	ANSI 控制码，用于设置文本颜色。\033 是控制码的开始，是八进制数字，[31m 表示将文本设置为红色。

	ANSI 控制码是用于在终端和控制台中控制文本格式和颜色的一种标准。它们通常用于在命令行界面 (CLI) 程序中输出彩色文本或者在文本模式下的图形界面 (GUI) 中输出文本。
	*/
	// 前景色
  fmt.Println("\033[30m 黑色 \033[0m")
  fmt.Println("\033[31m 红色 \033[0m")
  fmt.Println("\033[32m 绿色 \033[0m")
  fmt.Println("\033[33m 黄色 \033[0m")
  fmt.Println("\033[34m 蓝色 \033[0m")
  fmt.Println("\033[35m 紫色 \033[0m")
  fmt.Println("\033[36m 青色 \033[0m")
  fmt.Println("\033[37m 灰色 \033[0m")
  // 背景色
  fmt.Println("\033[40m 黑色 \033[0m")
  fmt.Println("\033[41m 红色 \033[0m")
  fmt.Println("\033[42m 绿色 \033[0m")
  fmt.Println("\033[43m 黄色 \033[0m")
  fmt.Println("\033[44m 蓝色 \033[0m")
  fmt.Println("\033[45m 紫色 \033[0m")
  fmt.Println("\033[46m 青色 \033[0m")
  fmt.Println("\033[47m 灰色 \033[0m")
	// 也可以这样写
	fmt.Printf("\x1b[0;%dm%s\x1b[0m", 31, "你好\n")
	/* 
	logrus也是支持颜色输出的
	配置项：==
	ForceColors：是否强制使用颜色输出。
	DisableColors：是否禁用颜色输出。
	ForceQuote：是否强制引用所有值。
	DisableQuote：是否禁用引用所有值。
	DisableTimestamp：是否禁用时间戳记录。
	FullTimestamp：是否在连接到 TTY 时输出完整的时间戳。
	TimestampFormat：用于输出完整时间戳的时间戳格式。

	*/
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		DisableColors:false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	})

	logrus.WithField("project", "learn").Infoln("message info")

	/* 
	自定义格式
	需要实现Formatter(entry *logrus.Entry) ([]byte, error) 接口
	*/
	// logrus.SetOutput(os.Stdout) //设置输出类型
	logrus.SetReportCaller(true) //开启返回函数名和行号
	logrus.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	logrus.SetLevel(logrus.DebugLevel) //设置最低的Level
	logrus.Errorln("你好")
	logrus.Infoln("你好")
	logrus.Warnln("你好")
	logrus.Println("你好")
	/* 
	显示行号
	logrus.SetReportCaller(true)
	*/
}

// 颜色
const (
	red = 31
	yellow = 33
	blue = 36
	gray = 37
)

type LogFormatter struct {}
// Format 实现Formatter(entry *logrus.Entry) ([]byte, error)接口
func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte,error) {
	// 根据不同的level去展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}
	var b *bytes.Buffer
	if entry.Buffer != nil{
		b = entry.Buffer
	}else{
		b = &bytes.Buffer{}
	}
	// 自定义日期格式
	timestamp := entry.Time.Format("2006-01-02 01:02:05")
	if entry.HasCaller() {
		// 自定义日期格式
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d",path.Base(entry.Caller.File),entry.Caller.Line)
		// 自定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp,levelColor,entry.Level,fileVal,funcVal,entry.Message)
	}else{
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s\n", timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
