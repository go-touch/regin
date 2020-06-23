package config_parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type JsonParser struct {

}

// 解析文件
func (jp *JsonParser) ParserToMap(filePath string) map[string]interface{} {
	// 从配置文件中读取json字符串
	buf, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Panicln("load config conf failed: ", err)
	}

	// 存储map
	mapConfig := make(map[string]interface{})
	err = json.Unmarshal(buf, &mapConfig)

	if err != nil {
		log.Panicln("decode config file failed:", string(buf), err)
	}
	return mapConfig
}