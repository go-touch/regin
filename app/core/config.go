package core

import (
	"github.com/go-touch/regin/base"
	"github.com/go-touch/regin/utils"
	"strconv"
	"strings"
)

// Define Config.
type Config struct {
	FileFormat string
	config     map[string]interface{}
}

// Init config
func (c *Config) Init(configPath string) {
	c.config = make(map[string]interface{})

	// Scan dir list & Load config
	configFiles := utils.File.ScanDir(configPath)
	for _, file := range configFiles {
		fileSplit := strings.Split(file.Name(), ".")
		filePath := utils.File.JoinPath(configPath, file.Name())

		// 配置类型过滤
		if fileSplit[1] == c.FileFormat {
			if config, err := utils.ConfigParser.ParserToMap(c.FileFormat, filePath); err != nil {
				panic(err.Error())
			} else {
				c.config[fileSplit[0]] = config
			}
		}
	}
}

// Get config.
func (c *Config) GetConfig(args ...string) *base.AnyValue {
	// If args is nil
	if args == nil {
		return base.Eval(c.config)
	}

	// Parser args
	argsGroup := strings.Split(args[0], ".")
	if argsGroup[0] == "" {
		return base.Eval(c.config)
	}

	// Walk config.
	var config interface{} = c.config
	for _, key := range argsGroup {
		switch config.(type) {
		case map[string]string:
			if c, ok := config.(map[string]string)[key]; ok {
				config = c
			}
		case map[string]interface{}:
			if c, ok := config.(map[string]interface{})[key]; ok {
				config = c
			}
		case []interface{}:
			intKey, err := strconv.Atoi(key)
			if err != nil {
				continue
			}
			if c := config.([]interface{})[intKey]; c != nil {
				config = c
			}
		default:
			break
		}
	}
	return base.Eval(config)
}
