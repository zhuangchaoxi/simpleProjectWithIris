package route

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"ms_web/api/config"
	"ms_web/api/log_audit"
)

func Register(app *iris.Application) {
	mvc.New(app.Party("/config")).Handle(new(config.Controller))
	mvc.New(app.Party("/log_audit")).Handle(new(log_audit.Controller))
}
