package parts

import (
	"regexp"
	"strings"
)

type Table struct {
	expr string
}

// make &Table.
func MakeTable() *Table {
	return &Table{}
}

// Get SQL Express statement.
func (t *Table) SetExpr(expr string) {
	t.expr = expr
}

// Get SQL Express statement.
func (t *Table) GetExpr() string {
	sqlExpr := make([]string, 0)

	// 表名处理
	t.expr = strings.TrimSpace(t.expr)
	t.expr = regexp.MustCompile(`\s+`+"").ReplaceAllString(t.expr, "|")
	exprGroup := strings.Split(t.expr, "|")
	for _, v := range exprGroup {
		if f := regexp.MustCompile(`\.` + "").FindString(v); f != "" {
			v = strings.Replace(v, ".", ".`", 1) + "`"
		} else if v != "as" {
			v = "`" + v + "`"
		}
		sqlExpr = append(sqlExpr, v)
	}
	if len(sqlExpr) == 0 {
		return ""
	}
	return strings.Join(sqlExpr, " ")
}
