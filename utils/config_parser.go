package utils

import (
	"errors"
	"github.com/go-touch/regin/utils/config_parser"
)

type ConfigFileParser struct {
	Parsers map[string]config_parser.BaseParser
}

// 定义ConfigReader结构体
var ConfigParser *ConfigFileParser

func init() {
	ConfigParser = &ConfigFileParser{
		Parsers: map[string]config_parser.BaseParser{
			"json": new(config_parser.JsonParser).Init(),
			"ini":  new(config_parser.IniParser).Init(),
		},
	}
}

// 使用解析器解析文件至map类型配置数据
func (cfp *ConfigFileParser) ParserToMap(parserName string, filePath string) (map[string]interface{}, error) {
	parser, ok := cfp.Parsers[parserName]
	if !ok {
		return nil, errors.New("can't find the config parser" + parserName + "...")
	}
	return parser.ParserToMap(filePath)
}