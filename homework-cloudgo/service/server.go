package service

//使用iris框架
import "github.com/kataras/iris"

func NewApp() *iris.Application {
	app := iris.Default()
<<<<<<< HEAD

=======
	
>>>>>>> c4a7558cfc7294299d97797f2b4004549c754c26
	//获取输入的两个参数并打印
	app.Get("/{str1}/{str2}", func(ctx iris.Context) {
		str1 := ctx.Params().Get("str1")
		str2 := ctx.Params().Get("str2")

		ctx.Writef("%s, %s", str1, str2)
	})
	return app
}
