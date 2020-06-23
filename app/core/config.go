package core

import (
	"github.com/go-touch/regin/base"
	"github.com/go-touch/regin/utils"
	"strings"
	"strconv"
)

// Define Config.
type Config struct {
	FileFormat string
	config     map[string]interface{}
}

// Init config
func (c *Config) Init(configPath string) {
	c.FileFormat = "json"
	c.config = make(map[string]interface{})

	// Scan dir list & Load config
	configFiles := utils.File.ScanDir(configPath)
	for _, file := range configFiles {
		fileSplit := strings.Split(file.Name(), ".")
		filePath := utils.File.JoinPath(configPath, file.Name())
		c.config[fileSplit[0]] = utils.ConfigParser.ParserToMap(c.FileFormat, filePath)
	}
}

// Get config.
func (c *Config) GetConfig(args ...string) (data *base.ConfigValue) {
	data = &base.ConfigValue{}

	if args == nil {
		data.Value = c.config
		return data
	}

	// Parser args
	argsGroup := strings.Split(args[0], ".")
	if argsGroup[0] == "" {
		data.Value = c.config
		return data
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
	data.Value = config
	return data
}
