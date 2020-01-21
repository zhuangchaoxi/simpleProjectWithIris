package config

import (
	"github.com/kataras/iris"
	"ms_web/core/response"
	"ms_web/core/response/status_code"
)

func GetList(ctx iris.Context) interface{} {
	return response.HttpResponse(status_code.SUCCESS, "", "test config...", "")
}
