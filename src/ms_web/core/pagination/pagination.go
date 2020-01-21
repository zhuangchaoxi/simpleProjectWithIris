package pagination

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
	"ms_web/conf"
	"strconv"
)

type MetaTableName interface {
	TableName() string
}

func Pagination(ctx iris.Context, db *gorm.DB, modelObjects interface{}, md MetaTableName) interface{} {
	var data map[string]interface{}
	data = make(map[string]interface{})
	total := 0
	next, previous := true, true
	r := ctx.Request()
	pagCfg := conf.GetCfg().Pagination
	page := r.FormValue("page")
	size := r.FormValue("size")
	db.Table(md.TableName()).Count(&total)
	if page == "" {
		page = strconv.Itoa(pagCfg.Page)
	}
	if size == "" {
		size = strconv.Itoa(pagCfg.Size)
	}
	result := db.Limit(cast.ToInt(size)).Offset((cast.ToInt(page) - 1) * cast.ToInt(size)).Find(modelObjects)
	if total < cast.ToInt(page)*cast.ToInt(size) {
		next = false
	}
	if cast.ToInt(page) == 1 || total < (cast.ToInt(page)-2)*cast.ToInt(size)+1 {
		previous = false
	}
	data["count"] = total
	data["next"] = next
	data["previous"] = previous
	data["results"] = result.Value
	return data
}
