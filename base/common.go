package base

import "github.com/go-touch/regin/utils"

/**************************************** 自定义函数类型 UserFunc ****************************************/
type UserFunc func(args ...interface{}) interface{}

// 转换为 string
func (uf UserFunc) ToString(args ...interface{}) string {
	value := utils.Convert.ToTargetType(uf(args...), utils.String)
	return string(value.(string))
}

// 转换为 StringMap
func (uf UserFunc) ToStringMap(args ...interface{}) map[string]string {
	value := utils.Convert.ToTargetType(uf(args...), utils.StringMap)
	return map[string]string(value.(map[string]string))
}