package config_parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type JsonParser struct {
	BaseParser
}

func (jp *JsonParser) Init() BaseParser {
	return jp
}

// 解析文件
func (jp *JsonParser) ParserToMap(filePath string) (map[string]interface{}, error) {
	// 从配置文件中读取json字符串
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		errorMsg := fmt.Sprintf("load config conf file '"+filePath+"'failed: %s", err.Error())
		return nil, errors.New(errorMsg)
	}

	// 存储map
	anyMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &anyMap)
	if err != nil {
		errorMsg := fmt.Sprintf("decode config file '"+filePath+"' failed: %s", err.Error())
		return nil, errors.New(errorMsg)
	}
	return anyMap, nil
}
