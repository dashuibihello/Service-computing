package main

import (
	"os"

	"github.com/dashuibihello/Service-computing/homework-cloudgo/service"
	"github.com/kataras/iris"
	flag "github.com/spf13/pflag"
)

//设置默认端口为8080
const (
	PORT string = "8080"
)

func main() {
	//如果没有监听到端口，则设为默认端口
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	//允许用户可以通过 -p设置端口
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}
	
	//新建服务
	app := service.NewApp()
	
	//服务启动
	app.Run(iris.Addr(":" + port))
}
