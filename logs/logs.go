package logs

import (
	"github.com/gookit/color"
	"log"
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

//这个包用来统一的日志输出处理
//目前只做简单两个方法 后续根据具体需要在这里增加日志操作
func Println(v ...interface{}) {
	if IsPrintLog {
		log.Println(v...)
	}
}

func print2(color2 color.Color, v ...interface{}) {
	if IsPrintLog {
		s := time.Now().Format("2006-01-02 15:04:05")
		v = append([]interface{}{"[" + s + "]"}, v...)
		color2.Light().Println(v...)
	}
}

func PrintlnSuccess(v ...interface{}) {
	print2(color.Green, v...)
}

func PrintlnInfo(v ...interface{}) {
	print2(color.LightCyan, v...)
}

func PrintlnWarning(v ...interface{}) {
	print2(color.Yellow, v...)
}

func PrintErr(v ...interface{}) {
	print2(color.FgLightRed, v...)
}

func Fatal(v ...interface{}) {
	log.Fatal(v...)
}
