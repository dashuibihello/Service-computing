package service

import "github.com/kataras/iris"

func NewApp() *iris.Application {
	app := iris.Default()

	app.Get("/{str1}/{str2}", func(ctx iris.Context) {
		str1 := ctx.Params().Get("str1")
		str2 := ctx.Params().Get("str2")

		ctx.Writef("%s, %s", str1, str2)
	})
	return app
}
