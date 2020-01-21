package log_audit

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

func querySearchOrFilter(ctx iris.Context, db *gorm.DB) *gorm.DB {
	r := ctx.Request()
	loginUserName := r.FormValue("login_username")
	if loginUserName != "" {
		db = db.Where("login_username = ?", loginUserName)
	}
	return db
}
