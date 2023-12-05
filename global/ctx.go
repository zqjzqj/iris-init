package global

import (
	"big_data_new/logs"
	"context"
	context2 "github.com/kataras/iris/v12/context"
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
	c, cc := context.WithTimeout(context.Background(), time.Second*1)
	defer cc()
	//等待时间根据自身应用来决定
	logs.PrintlnInfo("等待1s退出进程。。。")
	time.Sleep(1 * time.Second)
	_ = app.Shutdown(c)
}
