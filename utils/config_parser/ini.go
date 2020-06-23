package config_parser

type IniParser struct {
	content map[string]interface{}
}

// 定义JsonParser
var Ini *IniParser

func init() {
	Ini = &IniParser{
		content: make(map[string]interface{}),
	}
}

// 解析文件
func (ip *IniParser) ParserFile(filePath string) map[string]interface{} {
	// 读取文件
	/*config, err := fconf.NewFileConf(filePath)

	if err != nil {
		config = &fconf.Config{}
	}*/
	return map[string]interface{}{}
}


