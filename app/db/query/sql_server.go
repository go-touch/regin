package query

import (
	"database/sql"
	"github.com/go-touch/regin/app/db/query/parts"
)

type SqlServerQuery struct {
	Combine
	table       *parts.Table
	field       *parts.Field
	where       *parts.Where
	order       *parts.Order
	limit       *parts.Limit
	values      *parts.Values
	sqlExprType string // SQL表达式类型
	sqlExpr     string
	sqlParam    []interface{}
	result      interface{}
}

// 获取SQL语句
func (ssq *SqlServerQuery) Clone() BaseQuery {
	return &SqlServerQuery{
		Combine:  ssq.Combine,
		table:    parts.MakeTable(""),
		field:    parts.MakeField(),
		where:    parts.MakeWhere(),
		order:    parts.MakeOrder(""),
		limit:    parts.MakeLimit(0, 1000),
		values:   parts.MakeValues(),
		sqlExpr:  "",
		sqlParam: make([]interface{}, 0),
	}
}

// 重置结构体
func (ssq *SqlServerQuery) Reset() error {
	ssq.table = parts.MakeTable("")
	ssq.field = parts.MakeField()
	ssq.where = parts.MakeWhere()
	ssq.order = parts.MakeOrder("")
	ssq.limit = parts.MakeLimit(0, 1000)
	ssq.values = parts.MakeValues()
	ssq.sqlExpr = ""
	ssq.sqlParam = make([]interface{}, 0)
	return nil
}

// 设置table
func (ssq *SqlServerQuery) Table(tableName string) BaseQuery {
	ssq.table.SetExpr(tableName)
	return ssq
}

// 字段设置
func (ssq *SqlServerQuery) Field(field interface{}) BaseQuery {
	return ssq
}

// 获取字段设置
func (ssq *SqlServerQuery) GetField() *parts.Field {
	return ssq.field
}

// 查询条件
func (ssq *SqlServerQuery) Where(expr string, value interface{}, linkSymbol ...string) BaseQuery {
	ssq.where.SetExpr("AND", expr, value)
	return ssq
}

// 设置limit
func (ssq *SqlServerQuery) Order(expr string) BaseQuery {
	ssq.order.SetExpr(expr)
	return ssq
}

// 设置limit
func (ssq *SqlServerQuery) Limit(limit ...int) BaseQuery {
	ssq.limit.SetExpr(limit...)
	return ssq
}

// 设置数值
func (ssq *SqlServerQuery) Values(valueMap map[string]interface{}) BaseQuery {
	ssq.values.SetExpr(valueMap)
	return ssq
}

// 设置数值
func (ssq *SqlServerQuery) Set(valueMap map[string]interface{}) BaseQuery {
	ssq.values.SetExpr(valueMap)
	return ssq
}

// 设置Db
func (ssq *SqlServerQuery) SetDb(db *sql.DB) {
	ssq.db = db
}

// 获取SQL语句
func (ssq *SqlServerQuery) GetSql() string {
	return ""
}

func (ssq *SqlServerQuery) ExecSql() {

}

// 更新一条记录
func (ssq *SqlServerQuery) Update() *SqlServerQuery {
	return ssq
}

// 删除一条记录
func (ssq *SqlServerQuery) Delete() *SqlServerQuery {
	return ssq
}

// 获取一条记录
func (ssq *SqlServerQuery) FetchRow() *sql.Row {
	return nil
}

// 获取多条记录
func (ssq *SqlServerQuery) FetchAll() (*sql.Rows, error) {
	return ssq.QueryAll(ssq.sqlExpr, ssq.sqlParam...)
}

// 设置表达式类型
func (ssq *SqlServerQuery) SetSQLType(t string) error {
	ssq.sqlExprType = t
	return nil
}

// 创建SQL表达式
func (ssq *SqlServerQuery) CreateSQL() BaseQuery {
	return ssq
}

// 插入一条记录
func (ssq *SqlServerQuery) Modify() (sql.Result, error) {
	return ssq.Exec(ssq.sqlExpr, ssq.sqlParam)
}
