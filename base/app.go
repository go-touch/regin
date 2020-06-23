package base

import "github.com/go-touch/regin/utils"

// 获取json
func JsonResult() *Result {
	return ResultInvoker.CreateJson(200, "")
}

/**************************************** 结构体ConfigValue ****************************************/
// Define Data
type ConfigValue struct {
	Key   string
	Value interface{}
}

// Convert to int.
func (cv *ConfigValue) ToInt() int {
	return int(utils.Convert.ToTargetType(cv.Value, utils.Int).(int))
}

// Convert to string.
func (cv *ConfigValue) ToString() string {
	return string(utils.Convert.ToTargetType(cv.Value, utils.String).(string))
}

// Convert to stringMap.
func (cv *ConfigValue) ToStringMap() map[string]string {
	return map[string]string(utils.Convert.ToTargetType(cv.Value, utils.StringMap).(map[string]string))
}

// Convert to StringSlice.
func (cv *ConfigValue) ToStringSlice() []string {
	return []string(utils.Convert.ToTargetType(cv.Value, utils.StringSlice).([]string))
}

// Convert to AnyMap.
func (cv *ConfigValue) ToAnyMap() map[string]interface{} {
	return map[string]interface{}(utils.Convert.ToTargetType(cv.Value, utils.AnyMap).(map[string]interface{}))
}
