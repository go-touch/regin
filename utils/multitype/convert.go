package multitype

import (
	"strconv"
)

// 定义类型常量
const (
	IntT            = "Int"            // 类型: int
	ByteT           = "Byte"           // 类型: byte
	RuneT           = "Rune"           // 类型: rune
	StringT         = "String"         // 类型: string
	BoolT           = "Bool"           // 类型: bool
	IntSliceT       = "IntSlice"       // 类型: []int
	ByteSliceT      = "ByteSlice"      // 类型: []byte
	StringSliceT    = "StringSlice "   // 类型: []string
	AnySliceT       = "AnySlice "      // 类型: []interface{}
	StringMapSliceT = "StringMapSlice" // 类型: []map[string]string
	AnyMapSliceT    = "AnyMapSlice"    // 类型: []map[string]interface{}
	StringMapT      = "StringMap"      // 类型: map[string]string
	AnyMapT         = "AnyMap"         // 类型: map[string]interface{}
	ErrorT          = "Error"          // 类型: error
	NilT            = "Nil"            // 类型: nil
)

// 任意类型数据载体,实现类型转换
type AnyValue struct {
	value interface{}
}

// 传入任意值,获取一个载体结构体
func Eval(value interface{}) *AnyValue {
	return &AnyValue{value: value}
}

// 返回错误信息
func (av *AnyValue) ToError() error {
	switch GetType(av.value) {
	case ErrorT:
		return av.value.(error)
	}
	return nil
}

// 返回原值
func (av *AnyValue) ToValue() interface{} {
	return av.value
}

// 转成int类型
func (av *AnyValue) ToInt() int {
	if dstValue := ToType(av.value, IntT); dstValue != nil {
		return dstValue.(int)
	}
	return 0
}

// 转成byte类型
func (av *AnyValue) ToByte() byte {
	if dstValue := ToType(av.value, ByteT); dstValue != nil {
		return dstValue.(byte)
	}
	return byte(0)
}

// 转成string类型
func (av *AnyValue) ToString() string {
	if dstValue := ToType(av.value, StringT); dstValue != nil {
		return dstValue.(string)
	}
	return ""
}

// 转成bool类型
func (av *AnyValue) ToBool() bool {
	if dstValue := ToType(av.value, BoolT); dstValue != nil {
		return dstValue.(bool)
	}
	return false
}

// 转成[]int类型
func (av *AnyValue) ToIntSlice() []int {
	value := make([]int, 0)
	v := ToType(av.value, IntSliceT)
	if v != nil {
		value = v.([]int)
	}
	return value
}

// 转成[]byte类型
func (av *AnyValue) ToByteSlice() []byte {
	value := make([]byte, 0)
	v := ToType(av.value, ByteSliceT)
	if v != nil {
		value = v.([]byte)
	}
	return value
}

// 转成[]string类型
func (av *AnyValue) ToStringSlice() []string {
	value := make([]string, 0)
	v := ToType(av.value, StringSliceT)
	if v != nil {
		value = v.([]string)
	}
	return value
}

// 转成map[string]string类型
func (av *AnyValue) ToStringMap() map[string]string {
	value := map[string]string{}
	v := ToType(av.value, StringMapT)
	if v != nil {
		value = v.(map[string]string)
	}
	return value
}

// 转成map[string]interface{}类型
func (av *AnyValue) ToAnyMap() map[string]interface{} {
	value := map[string]interface{}{}
	v := ToType(av.value, AnyMapT)
	if v != nil {
		value = v.(map[string]interface{})
	}
	return value
}

// 转成[]map[string]string类型
func (av *AnyValue) ToStringMapSlice() []map[string]string {
	value := make([]map[string]string, 0)
	switch GetType(av.value) {
	case StringMapSliceT:
		return av.value.([]map[string]string)
	case AnyMapSliceT:
		for k, v := range av.value.([]map[string]interface{}) {
			value[k] = ToType(v, StringMapT).(map[string]string)
		}
	}
	return value
}

// 判断值类型
func GetType(value interface{}) string {
	switch value.(type) {
	case int:
		return IntT
	case byte:
		return ByteT
	case rune:
		return RuneT
	case string:
		return StringT
	case bool:
		return BoolT
	case []int:
		return IntSliceT
	case []byte:
		return ByteSliceT
	case []string:
		return StringSliceT
	case []interface{}, AnySlice:
		return AnySliceT
	case []map[string]string:
		return StringMapSliceT
	case []map[string]interface{}:
		return AnyMapSliceT
	case map[string]string:
		return StringMapT
	case map[string]interface{}, AnyMap:
		return AnyMapT
	case error:
		return ErrorT
	case nil:
		return NilT
	}
	return ""
}

// 转换为基本类型
func ToType(src interface{}, dstType string) interface{} {
	switch GetType(src) {
	case IntT:
		switch dstType { // int 转 int
		case IntT:
			return src
		case ByteT: // int 转 byte
			return byte(src.(int))
		case StringT: // int 转 string
			return strconv.Itoa(src.(int))
		case BoolT: // int 转 bool
			if src.(int) > 0 {
				return true
			}
			return false
		}
	case ByteT:
		switch dstType { // byte 转 int
		case IntT:
			return int(src.(byte))
		case ByteT: // byte 转 byte
			return src
		case StringT: // byte 转 string
			return string([]byte{src.(byte)})
		case BoolT: // byte 转 bool
			if src.(byte) > 0 {
				return true
			}
			return false
		}
	case StringT:
		switch dstType {
		case IntT: // string 转 int
			if v, err := strconv.Atoi(src.(string)); err != nil {
				return 0
			} else {
				return v
			}
		case ByteT: // string 转 byte
			return src
		case StringT: // string 转 string
			return src
		case BoolT: // string 转 bool
			if src.(string) == "true" {
				return true
			}
			return false
		case ByteSliceT: // string 转 []Byte
			return []byte(src.(string))
		}
	case BoolT:
		switch dstType {
		case IntT:
			if src == true {
				return 1
			}
			return 0
		case ByteT:
			if src == true {
				return byte(1)
			}
			return byte(0)
		case StringT:
			if src == true {
				return "true"
			}
			return "false"
		case BoolT:
			return src
		}
	case IntSliceT:
		switch dstType { // []int 转 []int
		case IntSliceT:
			return src
		case StringSliceT: // []int 转 []string
			var v []string
			for _, value := range src.([]int) {
				v = append(v, ToType(value, StringT).(string))
			}
			return v
		}
	case ByteSliceT:
		switch dstType {
		case StringT: // []byte 转 string
			return string(src.([]byte))
		case IntSliceT: // []byte 转 []int
			var v []int
			for _, value := range src.([]byte) {
				v = append(v, int(value))
			}
			return v
		case StringSliceT:
			var v []string
			for _, value := range src.([]byte) {
				v = append(v, ToType(value, StringT).(string))
			}
			return v
		}
	case StringSliceT:
		switch dstType {
		case IntSliceT: // []string 转 []int
			var v []int
			for _, value := range src.([]int) {
				v = append(v, ToType(value, IntT).(int))
			}
			return v
		case StringSliceT: // []string 转 []string
			return src
		}
	case AnySliceT:
		switch dstType {
		case IntSliceT: // []interface{} 转 []int
			var v []int

			// 类型断言
			var srcDeclare []interface{}
			if declare, ok := src.(AnySlice); ok {
				srcDeclare = declare
			} else {
				srcDeclare = src.([]interface{})
			}
			for _, value := range srcDeclare {
				if dstValue := ToType(value, IntT); dstValue != nil {
					v = append(v, dstValue.(int))
				} else {
					v = append(v, 0)
				}
			}
			return v
		case StringSliceT: // []interface{} 转 []string
			var v []string
			// 类型断言
			var srcDeclare []interface{}
			if declare, ok := src.(AnySlice); ok {
				srcDeclare = declare
			} else {
				srcDeclare = src.([]interface{})
			}
			for _, value := range srcDeclare {
				if dstValue := ToType(value, StringT); dstValue != nil {
					v = append(v, dstValue.(string))
				} else {
					v = append(v, "")
				}
			}
			return v
		case AnySliceT:
			var v []interface{}
			if declare, ok := src.(AnySlice); ok {
				v = declare
			} else {
				v = src.([]interface{})
			}
			return v
		}
	case StringMapT:
		switch dstType {
		case StringSliceT: // map[string]string 转 []string
			var v []string
			for _, value := range src.(map[string]string) {
				v = append(v, value)
			}
			return v
		case AnySliceT: // map[string]string 转 []interface{}
			var v []interface{}
			for _, value := range src.(map[string]string) {
				v = append(v, value)
			}
			return v
		case StringMapT: // // map[string]string 转 map[string]string
			return src
		case AnyMapT: // map[string]string 转 map[string]interface{}
			v := map[string]interface{}{}
			for key, value := range src.(map[string]string) {
				v[key] = value
			}
			return v
		}
	case AnyMapT:
		switch dstType {
		case StringMapT: // map[string]interface{} 转 map[string]string
			v := map[string]string{}

			// 类型断言
			var srcDeclare map[string]interface{}
			if declare, ok := src.(AnyMap); ok {
				srcDeclare = declare
			} else {
				srcDeclare = src.(map[string]interface{})
			}
			for key, value := range srcDeclare {
				if dstValue := ToType(value, StringT); dstValue != nil {
					v[key] = dstValue.(string)
				} else {
					v[key] = ""
				}
			}
			return v
		case AnyMapT: // map[string]interface{} 转 map[string]interface{}
			var v map[string]interface{}
			if declare, ok := src.(AnyMap); ok {
				v = declare
			} else {
				v = src.(map[string]interface{})
			}
			return v
		}
	}
	return nil
}
