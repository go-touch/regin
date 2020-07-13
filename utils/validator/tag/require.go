package tag

import "strings"

type Require struct {
	value        string
	whether      bool
	defaultValue string
}

// 创建 Require
func MakeRequire(whether bool) *Require {
	return &Require{
		whether:      whether,
		defaultValue: "",
	}
}

// 设置值
func (r *Require) SetValue(value string) {
	rule := strings.Split(value, "|")
	if len(rule) == 0 {
		return
	}
	// 是否校验
	if rule[0] == "true" {
		r.whether = true
	} else {
		r.whether = false
	}
	// 默认值
	if len(rule) == 2 {
		r.defaultValue = rule[1]
	}
}

// 数据校验
func (r *Require) Verify(key string, vMap *map[string]interface{}) *Result {
	result := &Result{
		Status: true,
		Msg:    "success",
	}
	if r.whether == true {
		if _, ok := (*vMap)[key]; !ok {
			result.Status = false
			result.Msg = "Param '" + key + "' is undefined."
		}
	} else {
		(*vMap)[key] = r.defaultValue
	}
	return result
}
