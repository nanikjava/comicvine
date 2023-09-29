package configuration

import (
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/json"
)

type ComicAppConfig struct {
	Token string `config:"token"`
}

func ParseConfig(c string) (error, *ComicAppConfig) {
	config.AddDriver(json.Driver)

	comicConfig := &ComicAppConfig{}

	err := config.LoadFiles(c)
	if err != nil {
		return err, nil
	}
	err = config.BindStruct("", comicConfig)

	fmt.Printf("config data: \n %#v\n", comicConfig)
	return err, comicConfig
}
