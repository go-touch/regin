package utils

import (
	"encoding/json"
	"strconv"
)

type ConvertHandler struct {
}

// 定义类型常量
const (
	Int            = "Int"            // 类型: int
	String         = "String"         // 类型: string
	IntSlice       = "IntSlice"       // 类型: []int
	StringSlice    = "StringSlice "   // 类型: []string
	AnySlice       = "AnySlice "      // 类型: []interface{}
	StringMapSlice = "StringMapSlice" // 类型: []map[string]string
	AnyMapSlice    = "AnyMapSlice"    // 类型: []map[string]interface{}
	StringMap      = "StringMap"      // 类型: map[string]string
	AnyMap         = "AnyMap"         // 类型: map[string]interface{}
)

// 定义ConvertHandler
var Convert *ConvertHandler

func init() {
	Convert = &ConvertHandler{}
}

// 类型断言
func (ch *ConvertHandler) GetType(object interface{}) string {
	switch t := object.(type) {
	case int:
		return Int
	case string:
		return String
	case []int:
		return IntSlice
	case []string:
		return StringSlice
	case []interface{}:
		return AnySlice
	case []map[string]string:
		return StringMapSlice
	case []map[string]interface{}:
		return AnyMapSlice
	case map[string]string:
		return StringMap
	case map[string]interface{}:
		return AnyMap
	default:
		_ = t
	}
	return ""
}

// 将当前对象转换为指定类型
func (ch *ConvertHandler) ToTargetType(object interface{}, targetType string) interface{} {
	switch ch.GetType(object) {
	case Int:
		if targetType == Int { // int 转 int
			return object
		} else if targetType == String { // int 转 string
			return strconv.Itoa(object.(int))
		} else if targetType == IntSlice { // int 转 []int
			return []int{object.(int)}
		} else if targetType == StringSlice { // int 转 []string
			V := ch.ToTargetType(object, String)
			return []string{V.(string)}
		}
	case String:
		if targetType == Int { // string 转 int
			if v, err := strconv.Atoi(object.(string)); err != nil {
				return 0
			} else {
				return v
			}
		} else if targetType == String { // string 转 string
			return object
		} else if targetType == IntSlice { // string 转 []int
			V := ch.ToTargetType(object, Int)
			return []int{V.(int)}
		} else if targetType == StringSlice { // string 转 []string
			return []string{object.(string)}
		}
	case IntSlice:
		if targetType == IntSlice { // []int 转 []int
			return object
		} else if targetType == StringSlice { // []int 转 []string
			var v []string

			for _, value := range object.([]int) {
				v = append(v, ch.ToTargetType(value, String).(string))
			}
			return v
		}
	case StringSlice:
		if targetType == IntSlice { // []string 转 []int
			var v []int

			for _, value := range object.([]int) {
				v = append(v, ch.ToTargetType(value, Int).(int))
			}
			return v
		} else if targetType == StringSlice { // []string 转 []string
			return object
		}
	case AnySlice:
		if targetType == StringSlice { // []interface{} 转 []string
			var v []string

			for _, value := range object.([]interface{}) {
				v = append(v, ch.ToTargetType(value, String).(string))
			}
			return v
		}
	case StringMapSlice:
		return StringMapSlice
	case AnyMapSlice:
		if targetType == StringMapSlice { // []map[string]interface{} 转 []map[string]string
			var v []map[string]string

			for _, value := range object.([]map[string]interface{}) {
				v = append(v, ch.ToTargetType(value, StringMap).(map[string]string))
			}
			return v
		}
	case StringMap:
		if targetType == StringMap { // map[string]string 转 map[string]string
			return object
		} else if targetType == AnyMap { // map[string]string 转 map[string]interface{}
			v := map[string]interface{}{}

			for key, value := range object.(map[string]string) {
				v[key] = value
			}
			return v
		} else if targetType == StringSlice { // map[string]string 转 []string
			var v []string
			for _, value := range object.(map[string]string) {
				v = append(v, value)
			}
			return v
		}
	case AnyMap:
		if targetType == StringMap { // map[string]interface{} 转 map[string]string
			v := map[string]string{}

			for key, value := range object.(map[string]interface{}) {
				v[key] = ch.ToTargetType(value, String).(string)
			}
			return v
		} else if targetType == AnyMap { // map[string]interface{} 转 map[string]interface{}
			return object
		}
	}

	// 其余类型转
	if targetType == Int {
		return 0
	} else if targetType == String {
		return ""
	}
	return nil
}

// 转换 []byte 到 map[string]interface{}
func (ch *ConvertHandler) ByteSliceToMap(byteSliceTData []byte) (result map[string]interface{}, err error) {
	err = json.Unmarshal(byteSliceTData, &result)
	return result, err
}

// 转换 map[string]interface{} 到 json(string)
func (ch *ConvertHandler) MapToJson(mapData map[string]interface{}) (jsonData string, err error) {
	result, err := json.Marshal(mapData)

	if err == nil {
		jsonData = string(result)
	}
	return jsonData, err
}

// 转换 json(string) 到 map[string]interface{}
func (ch *ConvertHandler) JsonToMap(jsonData string) (mapData map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(jsonData), &mapData)
	return mapData, err
}