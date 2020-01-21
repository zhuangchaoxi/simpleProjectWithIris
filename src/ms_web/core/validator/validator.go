package validator

//doc:  https://godoc.org/gopkg.in/go-playground/validator.v10

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/spf13/cast"
	"gopkg.in/go-playground/validator.v10"
	"ms_web/conf"
	"reflect"
	"strconv"
)

var validate *validator.Validate

func CheckParams(ctx iris.Context, paramsPtr interface{}) error {

	var err error
	validate = validator.New()
	request := ctx.Request()
	pagCfg := conf.GetCfg().Pagination
	rType := reflect.TypeOf(paramsPtr)
	rVal := reflect.ValueOf(paramsPtr)
	size := request.FormValue("size")
	if size != "" {
		size, err := strconv.Atoi(size)
		if err != nil {
			return err
		}
		if size > pagCfg.MaxSize {
			err := fmt.Errorf("the arg `size` must be less than %d", pagCfg.MaxSize)
			return err
		}
	}

	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("paramsPtr must be ptr to struct")
	}

	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		key := t.Tag.Get("json")
		value := request.FormValue(key)
		if t.Type.Kind() == reflect.Int {
			f.Set(reflect.ValueOf(cast.ToInt(value)))
		} else {
			f.Set(reflect.ValueOf(value))
		}
	}

	err = validate.Struct(paramsPtr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}
