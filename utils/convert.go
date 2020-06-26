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
	Byte           = "Byte"           // 类型: byte
	String         = "String"         // 类型: string
	Bool           = "Bool"           // 类型: bool
	IntSlice       = "IntSlice"       // 类型: []int
	ByteSlice      = "ByteSlice"      // 类型: []byte
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
	case byte:
		return Byte
	case string:
		return String
	case bool:
		return Bool
	case []int:
		return IntSlice
	case []byte:
		return ByteSlice
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
		} else if targetType == Byte { // int 转 byte
			return byte(object.(int))
		} else if targetType == String { // int 转 string
			return strconv.Itoa(object.(int))
		} else if targetType == Bool { // int 转 bool
			if object.(int) > 0 {
				return true
			}
			return false
		}
	case Byte:
		if targetType == Int { // byte 转 int
			return int(object.(byte))
		} else if targetType == Byte { // byte 转 byte
			return object
		} else if targetType == String { // byte 转 string
			return string([]byte{object.(byte)})
		} else if targetType == Bool { // byte 转 bool
			if object.(byte) > 0 {
				return true
			}
			return false
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
		} else if targetType == Bool { // string 转 bool
			if object.(string) == "true" {
				return true
			}
			return false
		} else if targetType == ByteSlice { // string 转 []Byte
			return []byte(object.(string))
		}
	case Bool:
		if targetType == Int { // bool 转 int
			if object == true {
				return 1
			}
			return 0
		} else if targetType == Byte { // bool 转 byte
			if object == true {
				return byte(1)
			}
			return byte(0)
		} else if targetType == String { // bool 转 string
			if object == true {
				return "true"
			}
			return "false"
		} else if targetType == Bool { // bool 转 bool
			return object
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
	case ByteSlice:
		if targetType == String { // []byte 转 string
			return string(object.([]byte))
		} else if targetType == Bool { // []byte 转 bool
			if object != nil {
				return true
			}
			return false
		} else if targetType == IntSlice { // []byte 转 []int
			var v []int
			for _, value := range object.([]byte) {
				v = append(v, int(value))
			}
			return v
		} else if targetType == StringSlice { // []byte 转 []string
			var v []string
			for _, value := range object.([]byte) {
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
		if targetType == IntSlice { // []interface{} 转 []int
			var v []int
			for _, value := range object.([]int) {
				v = append(v, ch.ToTargetType(value, Int).(int))
			}
			return v
		} else if targetType == StringSlice { // []interface{} 转 []string
			var v []string
			for _, value := range object.([]interface{}) {
				v = append(v, ch.ToTargetType(value, String).(string))
			}
			return v
		} else if targetType == AnySlice { // []interface{} 转 []interface{}
			return object
		}
	case StringMap:
		if targetType == StringSlice { // map[string]string 转 []string
			var v []string
			for _, value := range object.(map[string]string) {
				v = append(v, value)
			}
			return v
		} else if targetType == StringMap { // map[string]string 转 map[string]string
			return object
		} else if targetType == AnyMap { // map[string]string 转 map[string]interface{}
			v := map[string]interface{}{}
			for key, value := range object.(map[string]string) {
				v[key] = value
			}
			return v
		}
	case AnyMap:
		if targetType == StringMap { // map[string]interface{} 转 map[string]string
			value := map[string]string{}
			for k, v := range object.(map[string]interface{}) {
				value[k] = ch.ToTargetType(v, String).(string)
			}
			return value
		} else if targetType == AnyMap { // map[string]interface{} 转 map[string]interface{}
			return object
		}
	}

	// 其余类型转
	if targetType == Int {
		return 0
	} else if targetType == Byte {
		return byte(0)
	} else if targetType == String {
		return ""
	} else if targetType == Bool {
		return false
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
