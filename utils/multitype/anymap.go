package multitype

import (
	"strconv"
	"strings"
)

//  key is string,value is any type.
type AnyMap map[string]interface{}

// Set value.
func (am *AnyMap) Set(key string, value interface{}) {
	(*am)[key] = value
}

// Update value
func (am *AnyMap) Update(args string, dstValue interface{}) {
	var src interface{} = *am
	argsGroup := strings.Split(args, ".")
	// 遍历处理
	for _, key := range argsGroup {
		switch src.(type) {
		case map[string]string:
			if v, ok := src.(map[string]string)[key]; ok {
				src = v
			}
		case map[string]interface{}:
			if c, ok := src.(map[string]interface{})[key]; ok {
				src = c
			}
		case []interface{}:
			intKey, err := strconv.Atoi(key)
			if err != nil {
				continue
			}
			if c := src.([]interface{})[intKey]; c != nil {
				src = c
			}
		default:
			break
		}
	}

}

// Get value.
func (am *AnyMap) Get(key ...string) *AnyValue {
	if key == nil {
		return Eval(*am)
	} else if value, ok := (*am)[key[0]]; ok {
		return Eval(value)
	}
	return Eval(nil)
}
