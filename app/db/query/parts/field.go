package parts

import (
	"regexp"
	"strings"
)

type Field struct {
	expr      []string
	nameArray []string
	sqlExpr   string
}

// 构建Field结构体
func MakeField() *Field {
	return &Field{
		nameArray: make([]string, 0),
		sqlExpr:   "",
	}
}

// 设置表达式
func (f *Field) SetExpr(expr interface{}) {
	exprArray := make([]string, 0)

	// 传入字段处理
	if exprValue, ok := expr.(string); ok {
		exprArray = strings.Split(f.Filter(exprValue), ",")
	} else if exprValue, ok := expr.([]string); ok {
		exprArray = exprValue
	}
	// 过滤字符串
	for k, v := range exprArray {
		exprArray[k] = f.Filter(v)
	}
	// 设置SQL表达式
	sqlExpr := make([]string, 0)
	for _, value := range exprArray {
		valueGroup := strings.Split(value, " ")
		f.nameArray = append(f.nameArray, valueGroup[len(valueGroup)-1])
		for k, v := range valueGroup {
			if f := regexp.MustCompile(`\.` + "").FindString(v); f != "" {
				valueGroup[k] = strings.Replace(v, ".", ".`", 1) + "`"
			} else if regexp.MustCompile(`(SUM|sum|COUNT|count|as|AS)+`+"").FindString(v) == "" {
				valueGroup[k] = "`" + v + "`"
			}
		}
		sqlExpr = append(sqlExpr, strings.Join(valueGroup, " "))
	}
	if len(sqlExpr) > 0 {
		f.sqlExpr = strings.Join(sqlExpr, ",")
	}
}

// 获取表达式
func (f *Field) GetExpr() string {
	return f.sqlExpr
}

// 过滤字符串
func (f *Field) Filter(str string) string {
	str = strings.TrimSpace(str)
	str = strings.Trim(str, ",")
	str = regexp.MustCompile(`\s+`+"").ReplaceAllString(str, " ")
	str = regexp.MustCompile("`"+"").ReplaceAllString(str, "")
	return str
}

// 获取字段列表
func (f *Field) GetNameArray() []string {
	return f.nameArray
}
