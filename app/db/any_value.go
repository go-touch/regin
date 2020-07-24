package db

import (
	"database/sql"
	"fmt"
	"strconv"
)

// 定义类型常量
const (
	Int            = "Int"            // 类型: int
	Int64          = "Int64"          // 类型: int64
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
	SqlResult      = "SqlResult"      // 类型: sql.result
	Error          = "Error"          // 类型: error
)

type AnyValue struct {
	value interface{}
}

// 声明一个存储任意类型数据的结构体,然后可以进行类型转换
func Eval(value interface{}) *AnyValue {
	return &AnyValue{value: value}
}

// 返回预定义类型
func (av *AnyValue) getType() string {
	switch t := av.value.(type) {
	case int:
		return Int
	case int64:
		return Int64
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
	case sql.Result:
		return SqlResult
	case error:
		return Error
	default:
		_ = t
	}
	return ""
}

// 打印原值信息
func (av *AnyValue) PrintValue() {
	fmt.Printf("[orgin value]%v\n", av.value)
	return
}

// 获取原值
func (av *AnyValue) ToValue() interface{} {
	return av.value
}

// 转换成 - 操作数据库返回错误接口
func (av *AnyValue) ToError() error {
	switch av.getType() {
	case SqlResult:
		if _, err := av.value.(sql.Result).RowsAffected(); err != nil {
			return err
		}
		if _, err := av.value.(sql.Result).LastInsertId(); err != nil {
			return err
		}
	case Error:
		return av.value.(error)
	}
	return nil
}

// 转换成 int
func (av *AnyValue) ToInt() int {
	switch av.getType() {
	case String:
		if v, err := strconv.Atoi(av.value.(string)); err != nil {
			return 0
		} else {
			return v
		}
	case Int:
		return av.value.(int)
	case Int64:
		return int(av.value.(int64))
	}
	return 0
}

// 转换成 - 插入数据库记录主键id
func (av *AnyValue) ToLastInsertId() int {
	switch av.getType() {
	case SqlResult:
		if v, err := av.value.(sql.Result).LastInsertId(); err == nil {
			return int(v)
		}
	}
	return 0
}

// 转换成 - 修改数据库记录受影响行数
func (av *AnyValue) ToAffectedRows() int {
	switch av.getType() {
	case SqlResult:
		if v, err := av.value.(sql.Result).RowsAffected(); err == nil {
			return int(v)
		}
	}
	return 0
}

// 转换成 string
func (av *AnyValue) ToString() string {
	switch av.getType() {
	case String:
		return av.value.(string)
	case Int:
		strconv.Itoa(av.value.(int))
	}
	return ""
}

// 转换成 map[string]string
func (av *AnyValue) ToStringMap() map[string]string {
	rValue := map[string]string{}
	switch av.getType() {
	case StringMap:
		rValue = av.value.(map[string]string)
	case AnyMap:
		for key, value := range av.value.(map[string]interface{}) {
			if dv, ok := value.(string); ok {
				rValue[key] = dv
			} else if dv, ok := value.([]byte); ok {
				rValue[key] = string(dv)
			} else if dv, ok := value.(int64); ok {
				rValue[key] = strconv.Itoa(int(dv))
			} else if value == nil {
				rValue[key] = ""
			}
		}
	}
	return rValue
}

// 转换成 map[string]interface{}
func (av *AnyValue) ToAnyMap() map[string]interface{} {
	rValue := map[string]interface{}{}
	switch av.getType() {
	case AnyMap:
		for key, value := range av.value.(map[string]interface{}) {
			if dv, ok := value.([]byte); ok {
				rValue[key] = string(dv)
			} else if dv, ok := value.(int64); ok {
				rValue[key] = strconv.Itoa(int(dv))
			} else if value == nil {
				rValue[key] = ""
			}
		}
	}
	return rValue
}

// 转换成 []map[string]interface
func (av *AnyValue) ToAnyMapSlice() []map[string]interface{} {
	rValue := make([]map[string]interface{}, 0)
	switch av.getType() {
	case AnyMapSlice:
		for _, value := range av.value.([]map[string]interface{}) {
			subValue := map[string]interface{}{}
			for k, v := range value {
				if dv, ok := v.([]byte); ok {
					subValue[k] = string(dv)
				} else if dv, ok := v.(int64); ok {
					subValue[k] = int(dv)
				} else if v == nil {
					subValue[k] = ""
				}
			}
			rValue = append(rValue, subValue)
		}
	}
	return rValue
}

// 转换成 []map[string]string
func (av *AnyValue) ToStringMapSlice() []map[string]string {
	rValue := make([]map[string]string, 0)
	switch av.getType() {
	case AnyMapSlice:
		for _, value := range av.value.([]map[string]interface{}) {
			subValue := map[string]string{}
			for k, v := range value {
				if dv, ok := v.([]byte); ok {
					subValue[k] = string(dv)
				} else if dv, ok := v.(int64); ok {
					subValue[k] = strconv.Itoa(int(dv))
				} else if v == nil {
					subValue[k] = ""
				}
			}
			rValue = append(rValue, subValue)
		}
	}
	return rValue
}
