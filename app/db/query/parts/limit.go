package parts

import (
	"strconv"
	"strings"
)

type Limit struct {
	expr []int
}

func MakeLimit(limit ...int) *Limit {
	return &Limit{
		expr: limit,
	}
}

func (l *Limit) SetExpr(limit ...int) {
	l.expr = limit
}

func (l *Limit) GetExpr() string {
	sqlExpr := []string{"LIMIT"}
	if len(l.expr) == 0 {
		return ""
	} else if len(l.expr) == 1 {
		sqlExpr = append(sqlExpr, strconv.Itoa(l.expr[0]))
	} else {
		sqlExpr = append(sqlExpr, strconv.Itoa(l.expr[0])+","+strconv.Itoa(l.expr[1]))
	}
	return strings.Join(sqlExpr, " ")
}
