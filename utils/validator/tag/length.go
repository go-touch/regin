package tag

import (
	"strconv"
	"strings"
)

type Length struct {
	whether bool // 是否校验
	min     int  // 最小值
	max     int  // 最大值
}

// 创建 Length Tag
func MakeLength(whether bool) *Length {
	return &Length{
		whether: whether,
	}
}

// 设置值
func (l *Length) SetValue(value string) {
	l.whether = true
	rule := strings.Split(value, "|")

	l.min = l.toInt(rule[0]) // 最小值
	if len(rule) == 2 {      // 最大值
		l.max = l.toInt(rule[1])
	}
}

// 数据校验
func (l *Length) Verify(key string, vMap *map[string]interface{}) *Result {
	result := &Result{
		Status: true,
		Msg:    "success",
	}
	if l.whether == true {
		if value, ok := (*vMap)[key]; ok {
			// 最小值
			if len(value.(string)) < l.min {
				result.Status = false
				result.Msg = "Param '" + key + "' value length must be greater than " + l.toString(l.min) + "."
			}
			// 最小值
			if l.max > l.min && len(value.(string)) > l.max {
				result.Status = false
				result.Msg = "Param '" + key + "' value length must be less than " + l.toString(l.max) + "."
			}
		}
	}
	return result
}

// int转字符串
func (l *Length) toString(str int) string {
	return strconv.Itoa(str)
}

// 字符串转int
func (l *Length) toInt(str string) int {
	if v, err := strconv.Atoi(str); err != nil {
		return 0
	} else {
		return v
	}
}
