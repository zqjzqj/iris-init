package global

import (
	"context"
	context2 "github.com/kataras/iris/v12/context"
	"iris-init/logs"
	"time"
)

var globalContext = NewContext(context.Background())

type Context struct {
	Ctx context.Context
	Cc  context.CancelFunc
}

func NewContext(parent context.Context) *Context {
	c, cc := context.WithCancel(parent)
	return &Context{
		Ctx: c,
		Cc:  cc,
	}
}

func GetGlobalCtx() context.Context {
	return globalContext.Ctx
}

func CancelGlobalCtx() {
	cc := globalContext.Cc
	cc()
}

func HandleAppEndFunc(app context2.Application) {
	CancelGlobalCtx()
	c, cc := context.WithTimeout(context.Background(), time.Second*3)
	defer cc()
	//等待时间根据自身应用来决定
	logs.PrintlnInfo("等待3s退出进程。。。")
	time.Sleep(3 * time.Second)
	_ = app.Shutdown(c)
}
