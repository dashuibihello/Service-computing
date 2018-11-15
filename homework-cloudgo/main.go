package main

import (
	"os"

	"github.com/dashuibihello/Service-computing/homework-cloudgo/service"
	"github.com/kataras/iris"
	flag "github.com/spf13/pflag"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	//
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	app := service.NewApp()
	app.Run(iris.Addr(":" + port))
}
