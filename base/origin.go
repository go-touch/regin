package base

import "github.com/go-touch/regin/utils"

/**************************************** 数据类型 ****************************************/
type Any interface{}                      // [AnyType] can be any type
type AnyMap map[string]interface{}        // [MapType] key is string,value is any type
type StringMap map[string]string          // [MapType] key is string,value is string type
type SliceMap map[string][]string         // [MapType] key is string,value is Slice type
type GeneralMap map[string]AppAction      // [MapType] key is string,value is AppAction type
type StringMapSlice []map[string]string   // [SliceType] key is index,value为(key为string,value为string)的map
type AnyMapSlice []map[string]interface{} // [SliceType] key is index,value为(key为string,value为任意类型)的map

// Convert value To string value.
func (am *AnyMap) String(key string) string {
	value, ok := (*am)[key]
	if !ok {
		return ""
	}
	return string(utils.Convert.ToTargetType(value, utils.String).(string))
}

// Convert value To string map value.
func (am *AnyMap) StringMap(key string) StringMap {
	value, ok := (*am)[key]
	if !ok {
		return StringMap{}
	}
	return StringMap(utils.Convert.ToTargetType(value, utils.StringMap).(map[string]string))
}

// Convert value To string map value.
func (am *AnyMap) StringMapSlice(key string) StringMapSlice {
	value, ok := (*am)[key]
	if !ok {
		return StringMapSlice{}
	}
	return StringMapSlice(utils.Convert.ToTargetType(value, utils.StringMapSlice).([]map[string]string))
}

/**************************************** 接口 ****************************************/
type WebServer interface {
	Work(*Request) *Result
	Addr() string
	SSLCertPath() string
	ErrorCatch()
	GetError() error
}
