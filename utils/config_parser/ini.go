package config_parser

import (
	"errors"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

type IniParser struct {
	//BaseParser
	goIni   *ini.File
	options struct {
		LoadOptions ini.LoadOptions
		BlockMode   bool
	}
}

// 初始化
func (ip *IniParser) Init() BaseParser {
	ip.options.LoadOptions = ini.LoadOptions{
		IgnoreInlineComment: true,
	}
	ip.options.BlockMode = false
	return ip
}

// 解析配置到map[string]interface{}
func (ip *IniParser) ParserToMap(filePath string) (map[string]interface{}, error) {
	// 使用gopkg.in包,创建iniFile
	iniFile, err := ini.LoadSources(ip.options.LoadOptions, filePath)
	if err != nil {
		errorMsg := fmt.Sprintf("load config conf file '"+filePath+"'failed: %s", err.Error())
		return nil, errors.New(errorMsg)
	} else {
		iniFile.BlockMode = ip.options.BlockMode
	}

	// 读取分区列表
	list := map[string]map[string]string{}
	SectionArray := iniFile.SectionStrings()
	SectionArray = SectionArray[1:]
	for _, value := range SectionArray {
		keysArray := iniFile.Section(value).KeyStrings()
		keyMap := map[string]string{}
		for _, v := range keysArray {
			keyMap[v] = iniFile.Section(value).Key(v).String()
		}
		list[value] = keyMap
	}
	return ip.stdMap(list), nil
}

// 获取处理后的map
func (ip *IniParser) stdMap(list map[string]map[string]string) map[string]interface{} {
	rValue := map[string]interface{}{}

	for key, value := range list {
		anyMap := map[string]interface{}{}
		for k, v := range value {
			splitK := strings.Split(k, ".")
			if len(splitK) == 1 {
				anyMap[k] = v
			} else if len(splitK) == 2 {
				if _, ok := anyMap[splitK[0]]; !ok {
					anyMap[splitK[0]] = map[string]string{splitK[1]: v}
				} else {
					anyMap[splitK[0]].(map[string]string)[splitK[1]] = v
				}
			}
		}
		rValue[key] = anyMap
	}
	return rValue
}
