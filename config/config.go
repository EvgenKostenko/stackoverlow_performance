package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	APPName string `default:"app name"`

	DB struct {
		Host string
		Name string
	}

	API struct {
		Host string `default:"127.0.0.1:8080"`
	}
}{}

func init() {
	configor.Load(&Config, "./config/config.yml")
}
