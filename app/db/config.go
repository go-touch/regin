package db

import (
	"errors"
	"strconv"
	"strings"
)

type ConfigStorage struct {
	config interface{}
}

var Config *ConfigStorage

func init() {
	Config = &ConfigStorage{}

	Config.Init(map[string]interface{}{
		"master": map[string]interface{}{
			"driverName":     "mysql",
			"dataSourceName": "vdong:qO39eudVDA@tcp(39.106.157.226:3306)/plus_operation?charset=utf8",
			"maxIdleConn":    "100",
			"maxOpenConn":    "100",
		},
	})
}

// 初始化数据库配置
func (cs *ConfigStorage) Init(config map[string]interface{}) {
	cs.config = config
}

// 获取配置
func (cs *ConfigStorage) GetConfig(arg string) (config map[string]string, err error) {
	config = make(map[string]string)

	// 解析参数
	argGroup := strings.Split(arg, ".")
	if argGroup[0] == "" {
		return config, errors.New("the database's identify is not set")
	}

	// 遍历处理
	configTmp := cs.config
	for _, key := range argGroup {
		switch t := configTmp.(type) {
		case nil:
			configTmp = map[string]string{}
		case map[string]string:
			break
		case map[string]interface{}:
			if value, ok := configTmp.(map[string]interface{})[key]; ok {
				configTmp = value
			} else {
				configTmp = map[string]string{}
			}
		case []interface{}:
			if intKey, err := strconv.Atoi(key); err == nil {
				configTmp = configTmp.([]interface{})[intKey]
			}
		default:
			_ = t
			configTmp = map[string]string{}
			break
		}
	}

	// 类型断言并处理
	if value, ok := configTmp.(map[string]interface{}); ok {
		for k, v := range value {
			config[k] = v.(string)
		}
	} else { // 类型断言并处理
		value := configTmp.(map[string]string)
		if value != nil {
			config = value
		} else {
			config = map[string]string{}
		}
	}
	return config, nil
}
