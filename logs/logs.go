package logs

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"time"
)

var IsPrintLog bool

type LoggerService interface {
	Log(string)
}

type PrefixedLogger struct {
	Prefix string
}

func (s *PrefixedLogger) Log(msg string) {
	if s.Prefix == "" {
		PrintlnInfo(msg)
	} else {
		PrintlnInfo(s.Prefix, ":", msg)
	}
}

// 这个包用来统一的日志输出处理
// 目前只做简单两个方法 后续根据具体需要在这里增加日志操作
func Println(v ...interface{}) {
	if IsPrintLog {
		log.Println(v...)
	}
}

func print2(color2 color.Color, format string, v ...interface{}) {
	if IsPrintLog {
		s := time.Now().Format("2006-01-02 15:04:05")
		color2.Light().Println(s, fmt.Sprintf(format, v...))
	}
}

func PrintlnSuccess(format string, v ...interface{}) {
	print2(color.Green, format, v...)
}

func PrintlnInfo(format string, v ...interface{}) {
	print2(color.LightCyan, format, v...)
}

func PrintlnWarning(format string, v ...interface{}) {
	print2(color.Yellow, format, v...)
}

func PrintErr(format string, v ...interface{}) {
	print2(color.FgLightRed, format, v...)
}

func Fatal(format string, v ...interface{}) {
	PrintErr(format, v...)
	os.Exit(1)
}
