package config

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
)

type DBAppConfig struct {
	DbURL  string `config:"dburl"`
	DBName string `config:"dbname"`
}

func ParseConfig(c string) (error, *DBAppConfig) {
	config.AddDriver(json.Driver)

	comicConfig := &DBAppConfig{}

	err := config.LoadFiles(c)
	if err != nil {
		return err, nil
	}
	err = config.BindStruct("", comicConfig)

	fmt.Printf("config data: \n %#v\n", comicConfig)
	return err, comicConfig
}
