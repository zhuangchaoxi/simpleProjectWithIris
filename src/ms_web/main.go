package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"ms_web/core/middleware/auth"
	"ms_web/core/ms_log"
	"ms_web/route"
)

func main() {
	ms_log.InitLog()
	app := iris.New()
	app.Use(auth.Authentication)
	app.Use(recover.New())
	app.Use(logger.New())
	route.Register(app)
	err := app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.TOML("ms_web/conf/iris.tml")))
	if err != nil {
		panic(err)
	}
}
