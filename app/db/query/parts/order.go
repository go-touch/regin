package parts

import (
	"regexp"
	"strings"
)

type Order struct {
	members []*SubOrder
}

type SubOrder struct {
	expr string // 表达式, for example: id ASC
}

func MakeOrder() *Order {
	return &Order{
		members: make([]*SubOrder, 0),
	}
}

// 设置表达式
func (o *Order) SetExpr(expr string) {
	o.members = append(o.members, &SubOrder{
		expr: expr,
	})
}

// 获取表达式
func (o *Order) GetExpr() string {
	sqlExpr := make([]string, 0, 10)
	for _, subOrder := range o.members {
		subOrder.expr = regexp.MustCompile(`\s+`+"").ReplaceAllString(subOrder.expr, "|")
		exprArray := strings.Split(subOrder.expr, "|")
		if len(exprArray) == 1 {
			exprArray[0] = "`" + exprArray[0] + "`"
		} else if len(exprArray) > 1 {
			exprArray = exprArray[0:2]
			if f := regexp.MustCompile(`\.` + "").FindString(exprArray[0]); f != "" {
				exprArray[0] = strings.Replace(exprArray[0], ".", ".`", 1) + "`"
			} else {
				exprArray[0] = "`" + exprArray[0] + "`"
			}
			exprArray[1] = strings.ToUpper(exprArray[1])
		}
		if len(exprArray) > 0{
			sqlExpr = append(sqlExpr, strings.Join(exprArray, " "))
		}
	}
	if len(sqlExpr) == 0 {
		return ""
	}
	return "ORDER BY " + strings.Join(sqlExpr, ",")
}
