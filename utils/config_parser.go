package utils

import (
	"github.com/go-touch/regin/utils/config_parser"
	"strconv"
)

type ConfigFileParser struct {
	Config  interface{}
	Parsers map[string]config_parser.BaseParser
}

// 定义ConfigReader结构体
var ConfigParser *ConfigFileParser

func init() {
	ConfigParser = &ConfigFileParser{
		Config: nil,
		Parsers: map[string]config_parser.BaseParser{
			"json": &config_parser.JsonParser{},
		},
	}
}

// 设置配置
func (cfp *ConfigFileParser) SetConfig(config interface{}) *ConfigFileParser {
	cfp.Config = config
	return cfp
}

// 获取配置
func (cfp *ConfigFileParser) GetConfig(argsGroup []string) interface{} {
	// 根据数据配置类型
	switch t := cfp.Config.(type) {
	case map[string]interface{}:
		return cfp.GetFromMap(argsGroup)
	default:
		_ = t
	}
	return cfp.Config
}

// 使用解析器解析文件至map类型配置数据
func (cfp *ConfigFileParser) ParserToMap(parserName string, filePath string) map[string]interface{} {
	return cfp.Parsers[parserName].ParserToMap(filePath)
}

// 从Map类型数据获取配置
func (cfp *ConfigFileParser) GetFromMap(argsGroup []string) interface{} {
	// 遍历处理
	for _, key := range argsGroup {
		switch t := cfp.Config.(type) {
		case nil:
			return ""
		case string:
			return cfp.Config
		case map[string]interface{}:
			cfp.Config = cfp.Config.(map[string]interface{})[key]
		case []interface{}:
			if intKey, err := strconv.Atoi(key); err == nil {
				cfp.Config = cfp.Config.([]interface{})[intKey]
			}
		default:
			_ = t
			cfp.Config = ""
		}
	}
	return cfp.Config
}
