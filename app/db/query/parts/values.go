package parts

import "strings"

type Values struct {
	expr string        // 表达式
	args []interface{} // 参数
}

// 创建values结构体
func MakeValues() *Values {
	return &Values{
		expr: "",
		args: make([]interface{}, 0),
	}
}

// 设置SQL表达式
func (v *Values) SetExpr(valueMap map[string]interface{}) {
	sqlExpr := make([]string, 0)
	for key, value := range valueMap {
		key = "`" + key + "`"
		sqlExpr = append(sqlExpr, strings.TrimSpace(key))
		v.args = append(v.args, value)
	}
	if len(sqlExpr) > 0 {
		v.expr = "(" + strings.Join(sqlExpr, ",") + ")"
	}
}

// 获取SQL表达式
func (v *Values) GetExpr() string {
	return v.expr
}

// 获取SQL表达式
func (v *Values) GetExprValues() string {
	sqlExpr := make([]string, 0)
	for i := 0; i < len(v.args); i++ {
		sqlExpr = append(sqlExpr, "?")
	}
	return "(" + strings.Join(sqlExpr, ",") + ")"
}

// 获取SQL表达式参数
func (v *Values) GetArgs() []interface{} {
	return v.args
}
