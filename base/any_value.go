package base

import (
	"github.com/go-touch/regin/utils"
)

// 用于转换数据类型
type AnyValue struct {
	value interface{}
}

// 声明一个存储任意类型数据的结构体,然后可以进行类型转换
func Eval(value interface{}) *AnyValue {
	return &AnyValue{value: value}
}

// 返回原值
func (av *AnyValue) ToValue() interface{} {
	return av.value
}

// 转成int类型
func (av *AnyValue) ToInt() int {
	return utils.Convert.ToTargetType(av.value, utils.Int).(int)
}

// 转成byte类型
func (av *AnyValue) ToByte() byte {
	return utils.Convert.ToTargetType(av.value, utils.Byte).(byte)
}

// 转成string类型
func (av *AnyValue) ToString() string {
	return utils.Convert.ToTargetType(av.value, utils.String).(string)
}

// 转成bool类型
func (av *AnyValue) ToBool() bool {
	return utils.Convert.ToTargetType(av.value, utils.Bool).(bool)
}

// 转成[]int类型
func (av *AnyValue) ToIntSlice() []int {
	value := make([]int, 0)
	v := utils.Convert.ToTargetType(av.value, utils.IntSlice)

	if v != nil {
		value = v.([]int)
	}
	return value
}

// 转成[]byte类型
func (av *AnyValue) ToByteSlice() []byte {
	value := make([]byte, 0)
	v := utils.Convert.ToTargetType(av.value, utils.ByteSlice)

	if v != nil {
		value = v.([]byte)
	}
	return value
}

// 转成[]string类型
func (av *AnyValue) ToStringSlice() []string {
	value := make([]string, 0)
	v := utils.Convert.ToTargetType(av.value, utils.StringSlice)

	if v != nil {
		value = v.([]string)
	}
	return value
}

// 转成map[string]string类型
func (av *AnyValue) ToStringMap() map[string]string {
	value := map[string]string{}
	v := utils.Convert.ToTargetType(av.value, utils.StringMap)
	if v != nil {
		value = v.(map[string]string)
	}
	return value
}

// 转成map[string]interface{}类型
func (av *AnyValue) ToAnyMap() map[string]interface{} {
	value := map[string]interface{}{}
	v := utils.Convert.ToTargetType(av.value, utils.AnyMap)
	if v != nil {
		value = v.(map[string]interface{})
	}
	return value
}

// 转成[]map[string]string类型
func (av *AnyValue) ToStringMapSlice() []map[string]string {
	value := make([]map[string]string, 0)
	t := utils.Convert.GetType(av.value)

	if t == "StringMapSlice" {
		value = av.value.([]map[string]string)
	} else if t == "AnyMapSlice" {
		for k, v := range av.value.([]map[string]interface{}) {
			value[k] = utils.Convert.ToTargetType(v, utils.StringMap).(map[string]string)
		}
	}
	return value
}
