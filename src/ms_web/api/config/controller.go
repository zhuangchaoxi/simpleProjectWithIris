package config

import (
	"github.com/kataras/iris"
	"ms_web/core/response"
	"ms_web/core/response/status_code"
	"ms_web/core/validator"
)

type Controller struct{}

func (m *Controller) Get(ctx iris.Context) interface{} {
	var lc listCheckOption
	err := validator.CheckParams(ctx, &lc)
	if err != nil {
		msg, data := "参数校验失败，请检查！", ""
		return response.HttpResponse(status_code.FORMAT_ERROR, msg, data, err.Error())
	}
	return GetList(ctx)
}
