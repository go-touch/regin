package parts

import (
	"strings"
)

type Set struct {
	expr string        // 表达式
	args []interface{} // 参数
}

// 创建Set结构体
func MakeSet() *Set {
	return &Set{
		expr: "",
		args: make([]interface{}, 0),
	}
}

// 设置SQL表达式
func (s *Set) SetExpr(valueMap map[string]interface{}) {
	sqlExpr := make([]string, 0)
	for key, value := range valueMap {
		key = "`" + strings.TrimSpace(key) + "`"
		unit := []string{key, "=", "?"}
		sqlExpr = append(sqlExpr, strings.Join(unit, " "))
		s.args = append(s.args, value)
	}
	if len(sqlExpr) > 0 {
		s.expr = strings.Join(sqlExpr, ",")
	}
}

// 获取SQL表达式
func (s *Set) GetExpr() string {
	return s.expr
}

// 获取SQL表达式参数
func (s *Set) GetArgs() []interface{} {
	return s.args
}