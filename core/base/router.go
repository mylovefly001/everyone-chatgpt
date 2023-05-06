package base

import (
	"github.com/kataras/iris/v12"
	recover "github.com/kataras/iris/v12/middleware/recover"
	"time"
)

func Router(app *iris.Application) {
	app.UseRouter(recover.New())
	app.Get("/", func(ctx iris.Context) {
		_, _ = ctx.WriteString("hello...")
	})
	app.Get("/test", func(ctx iris.Context) {
		//睡10秒，模拟耗时操作
		time.Sleep(10 * time.Second)
		_, _ = ctx.WriteString("hello...")
	})
	app.Get("/elb-status", func(ctx iris.Context) {
		_, _ = ctx.WriteString("ok")
	})
}
