package conf

import (
	"fmt"
	"github.com/gogf/gf/third/github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

var (
	cfg  *IrisConfig
	once sync.Once
)

type IrisConfig struct {
	DataBase   map[string]database
	Pagination pagination
}

type database struct {
	Server   string
	Port     string
	Username string
	Password string
	DbName   string
	ConnMax  int `toml:"connection_max"`
}

type pagination struct {
	Page    int
	Size    int
	MaxSize int `toml:"max_size"`
}

func GetCfg() *IrisConfig {
	once.Do(func() {
		filePath, err := filepath.Abs("ms_web/conf/iris.tml")
		if err != nil {
			panic(err)
		}
		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			panic(err)
		}
	})
	return cfg
}
