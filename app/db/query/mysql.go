package query

import (
	"database/sql"
	"fmt"
	"github.com/go-touch/regin/app/db/query/parts"
	"regexp"
	"strings"
)

const (
	sqlSELECT = "SELECT|{FIELD}|FROM|{TABLE}|{WHERE}|{ORDER}|{LIMIT}"
	sqlINSERT = "INSERT|INTO|{TABLE}|{FIELD}|VALUES|{VALUES}"
	sqlUPDATE = "UPDATE|{TABLE}|SET|{SET}|{WHERE}"
	sqlDELETE = "DELETE|FROM|{TABLE}|{WHERE}"
)

type MysqlQuery struct {
	Combine
	table       *parts.Table
	field       *parts.Field
	where       *parts.Where
	order       *parts.Order
	limit       *parts.Limit
	values      *parts.Values
	set         *parts.Set
	sqlExprType string        // SQL表达式类型
	sqlExpr     string        // SQL表达式
	sqlParam    []interface{} // SQL参数
	result      interface{}
}

// 复制结构体
func (mq *MysqlQuery) Clone() BaseQuery {
	return &MysqlQuery{
		Combine:     mq.Combine,
		table:       parts.MakeTable(""),
		field:       parts.MakeField(),
		where:       parts.MakeWhere(),
		order:       parts.MakeOrder(""),
		limit:       parts.MakeLimit(0, 1000),
		values:      parts.MakeValues(),
		set:         parts.MakeSet(),
		sqlExprType: "",
		sqlExpr:     "",
		sqlParam:    make([]interface{}, 0),
	}
}

// 重置结构体
func (mq *MysqlQuery) Reset() error {
	mq.table = parts.MakeTable("")
	mq.field = parts.MakeField()
	mq.where = parts.MakeWhere()
	mq.order = parts.MakeOrder("")
	mq.limit = parts.MakeLimit(0, 1000)
	mq.values = parts.MakeValues()
	mq.set = parts.MakeSet()
	mq.sqlExprType = ""
	mq.sqlExpr = ""
	mq.sqlParam = make([]interface{}, 0)
	return nil
}

// 设置table
func (mq *MysqlQuery) Table(tableName string) BaseQuery {
	mq.table.SetExpr(tableName)
	return mq
}

// 字段设置
func (mq *MysqlQuery) Field(field interface{}) BaseQuery {
	mq.field.SetExpr(field)
	return mq
}

// 获取字段设置
func (mq *MysqlQuery) GetField() *parts.Field {
	return mq.field
}

// 查询Where
func (mq *MysqlQuery) Where(expr string, value interface{}, linkSymbol ...string) BaseQuery {
	ls := "AND"
	if linkSymbol != nil && linkSymbol[0] != "" {
		ls = linkSymbol[0]
	}
	mq.where.SetExpr(ls, expr, value)
	return mq
}

// 设置Order
func (mq *MysqlQuery) Order(expr string) BaseQuery {
	mq.order.SetExpr(expr)
	return mq
}

// 设置limit
func (mq *MysqlQuery) Limit(limit ...int) BaseQuery {
	mq.limit.SetExpr(limit...)
	return mq
}

// 设置数值
func (mq *MysqlQuery) Values(valueMap map[string]interface{}) BaseQuery {
	mq.values.SetExpr(valueMap)
	return mq
}

// 设置数值
func (mq *MysqlQuery) Set(valueMap map[string]interface{}) BaseQuery {
	mq.set.SetExpr(valueMap)
	return mq
}

// 设置表达式类型
func (mq *MysqlQuery) SetSQLType(t string) error {
	mq.sqlExprType = t
	return nil
}

// 创建SQL表达式
func (mq *MysqlQuery) CreateSQL() BaseQuery {
	switch strings.ToUpper(mq.sqlExprType) {
	case "SELECT":
		mq.sqlExpr = sqlSELECT
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{FIELD}", mq.field.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{TABLE}", mq.table.GetExpr(), 1)
		// 查询条件
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{WHERE}", mq.where.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{ORDER}", mq.order.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{LIMIT}", mq.limit.GetExpr(), 1)
		mq.sqlParam = append(mq.sqlParam, mq.where.GetArgs()...)
	case "INSERT":
		mq.sqlExpr = sqlINSERT
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{TABLE}", mq.table.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{FIELD}", mq.values.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{VALUES}", mq.values.GetExprValues(), 1)
		mq.sqlParam = append(mq.sqlParam, mq.values.GetArgs()...)
	case "UPDATE":
		mq.sqlExpr = sqlUPDATE
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{TABLE}", mq.table.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{SET}", mq.set.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{WHERE}", mq.where.GetExpr(), 1)
		mq.sqlParam = append(mq.sqlParam, mq.set.GetArgs()...)
		mq.sqlParam = append(mq.sqlParam, mq.where.GetArgs()...)
	case "DELETE":
		mq.sqlExpr = sqlDELETE
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{TABLE}", mq.table.GetExpr(), 1)
		mq.sqlExpr = strings.Replace(mq.sqlExpr, "{WHERE}", mq.where.GetExpr(), 1)
		mq.sqlParam = append(mq.sqlParam, mq.where.GetArgs()...)
	}
	mq.sqlExpr = strings.Replace(mq.sqlExpr, "|", " ", -1)
	mq.sqlExpr = regexp.MustCompile("\\s+").ReplaceAllString(mq.sqlExpr, " ")
	return mq
}

// 获取SQL语句
func (mq *MysqlQuery) GetSql() string {
	return mq.sqlExpr + "\n" + fmt.Sprintf("Args: %v", mq.sqlParam)
}

// 获取一条记录
func (mq *MysqlQuery) FetchRow() *sql.Row {
	return mq.QueryRow(mq.sqlExpr, mq.sqlParam...)
}

// 获取多条记录
func (mq *MysqlQuery) FetchAll() (*sql.Rows, error) {
	return mq.QueryAll(mq.sqlExpr, mq.sqlParam...)
}

// 插入|更新|删除记录
func (mq *MysqlQuery) Modify() (sql.Result, error) {
	return mq.Exec(mq.sqlExpr, mq.sqlParam...)
}
