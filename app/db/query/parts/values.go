package parts

import (
	"sort"
	"strings"
)

type Values struct {
	member []*SubValues
}

type SubValues struct {
	expr     string
	argsExpr string
	args     []interface{}
}

// Make SubValue.
func MakeSubValues(anyMap map[string]interface{}) *SubValues {
	expr := ""
	argsExpr := ""
	args := make([]interface{}, 0)

	// SQL表达式
	sqlExpr := make([]string, 0)
	sqlValueExpr := make([]string, 0)
	sortBy := make([]string, 0)
	for key := range anyMap {
		key = strings.TrimSpace(key)
		sortBy = append(sortBy, key)
	}
	// Sort key.
	sort.Strings(sortBy)
	for _, key := range sortBy {
		sqlExpr = append(sqlExpr, "`"+key+"`")
		sqlValueExpr = append(sqlValueExpr, "?")
		args = append(args, anyMap[key])
	}
	if len(sqlExpr) > 0 {
		expr = "(" + strings.Join(sqlExpr, ",") + ")"
		argsExpr = "(" + strings.Join(sqlValueExpr, ",") + ")"
	}
	return &SubValues{
		expr:     expr,
		argsExpr: argsExpr,
		args:     args,
	}
}

// Set SQL expr.
func (sv *SubValues) GetExpr() string {
	return sv.expr
}

// Get SQL VALUES expr.
func (sv *SubValues) GetArgsExpr() string {
	return sv.argsExpr
}

// Get SQL VALUES args.
func (sv *SubValues) GetArgs() []interface{} {
	return sv.args
}

// Make Values.
func MakeValues() *Values {
	return &Values{
		member: make([]*SubValues, 0),
	}
}

// Set SQL expr.
func (v *Values) SetExpr(anyMap map[string]interface{}) {
	v.member = append(v.member, MakeSubValues(anyMap))
}

// Get SQL expr.
func (v *Values) GetExpr() string {
	if len(v.member) > 0 {
		return v.member[0].GetExpr()
	}
	return ""
}

// Get SQL args expr.
func (v *Values) GetArgsExpr() string {
	argsExpr := make([]string, 0)
	for _, subValues := range v.member {
		argsExpr = append(argsExpr, subValues.GetArgsExpr())
	}
	return strings.Join(argsExpr, ",")
}

// 获取SQL表达式参数
func (v *Values) GetArgs() []interface{} {
	args := make([]interface{}, 0)
	for _, subValues := range v.member {
		args = append(args, subValues.GetArgs()...)
	}
	return args
}
