package db_pool

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ms_web/conf"
	"strings"
	"time"
)

var db0 *gorm.DB
var db1 *gorm.DB

func GetDb0() *gorm.DB {
	return db0
}

func GetDb1() *gorm.DB {
	return db1
}

func db0Init() {
	DbCfg := conf.GetCfg().DataBase["db0"]
	path := strings.Join([]string{DbCfg.Username, ":", DbCfg.Password, "@(", DbCfg.Server, ":", DbCfg.Port, ")/", "db_sentry", "?charset=utf8&parseTime=true"}, "")
	var err error
	db0, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db0.SingularTable(true)
	db0.DB().SetConnMaxLifetime(1 * time.Second)
	db0.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db0.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db0.SingularTable(true)        //表生成结尾不带s
	db0.LogMode(true)              // 启用Logger，显示详细日志
}

func db1Init() {
	DbCfg := conf.GetCfg().DataBase["db1"]
	path := strings.Join([]string{DbCfg.Username, ":", DbCfg.Password, "@(", DbCfg.Server, ":", DbCfg.Port, ")/", "db_sentry", "?charset=utf8&parseTime=true"}, "")
	var err error
	db1, err = gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	db1.SingularTable(true)
	db1.DB().SetConnMaxLifetime(1 * time.Second)
	db1.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db1.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db1.SingularTable(true)        //表生成结尾不带s
	db1.LogMode(true)              // 启用Logger，显示详细日志
}

func init() {
	db0Init()
	db1Init()
}
