package log_audit

import (
	"github.com/kataras/iris"
	"ms_web/core/db_pool"
	"ms_web/core/pagination"
	"ms_web/core/response"
	"ms_web/core/response/status_code"
	"ms_web/models"
)

func GetList(ctx iris.Context) interface{} {
	var (
		la       models.LogAudit
		logAudit []models.LogAudit
	)
	db := db_pool.GetDb0()
	db = querySearchOrFilter(ctx, db)
	data := pagination.Pagination(ctx, db, &logAudit, &la)
	return response.HttpResponse(status_code.SUCCESS, "", data, "")
}
