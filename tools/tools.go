package tools

import (
	"encoding/json"
	"io/ioutil"
	"utils/logger"
)

func File2JSON(path string, v interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

type IInitConfig interface {
	InitConfig()
}

func LoadConfig(file string, v interface{}) {
	if err := File2JSON(file, v); err != nil {
		logger.F("load config file:%v failed: %v", file, err)
	}
	if c, ok := v.(IInitConfig); ok {
		c.InitConfig()
	}
}
