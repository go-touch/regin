package parts

import (
	"regexp"
	"strings"
)

type Where struct {
	members []*SubWhere   //
	args    []interface{} // 参数
}

type SubWhere struct {
	linkSymbol string      // 连接符, for example: AND
	expr       string      // 表达式, for example: `username` = ?
	value      interface{} // 对应值, for example: admin
}

// 生成Where构件
func MakeWhere() *Where {
	return &Where{
		members: make([]*SubWhere, 0),
		args:    make([]interface{}, 0),
	}
}

// 设置表达式单元
func (w *Where) SetExpr(linkSymbol string, expr string, value interface{}) {
	if len(w.members) == 0 {
		linkSymbol = ""
	}
	w.members = append(w.members, &SubWhere{
		linkSymbol: strings.ToUpper(linkSymbol),
		expr:       strings.TrimSpace(expr),
		value:      value,
	})
}

// 获取sql表达式
func (w *Where) GetExpr() string {
	sqlExpr := []string{"WHERE"}

	// 遍历处理
	for _, subWhere := range w.members {
		subWhere.expr = regexp.MustCompile(`\s+`+"").ReplaceAllString(subWhere.expr, "|")
		exprArray := strings.Split(subWhere.expr, "|")
		// 长度判断
		if len(exprArray) > 0 {
			exprArray[0] = "`" + exprArray[0] + "`"
			subWhere.expr = strings.Join(exprArray, " ")
			if regexp.MustCompile(`\s(=|!=|like|not like|>|>=|<|<=)\s(\?)`+"$").FindString(subWhere.expr) != "" {
				sqlExpr = append(sqlExpr, subWhere.linkSymbol, subWhere.expr)
				w.args = append(w.args, subWhere.value)
			} else if regexp.MustCompile(`\s(in)\s\(\S+\)`+"$").FindString(subWhere.expr) != "" {
				sqlExpr = append(sqlExpr, subWhere.linkSymbol, subWhere.expr)
				if v, ok := subWhere.value.([]interface{}); ok {
					w.args = append(w.args, v...)
				}
			}
		}
	}
	if len(sqlExpr) == 1 {
		return ""
	}
	return strings.Join(sqlExpr, " ")
}

// 获取值
func (w *Where) GetArgs() []interface{} {
	return w.args
}
