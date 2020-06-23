package query

import (
	"database/sql"
)

type SqlServerQuery struct {
	BaseQuery
	db *sql.DB
}

// 获取SQL语句
func (ssq *SqlServerQuery) Copy() BaseQuery {
	return &SqlServerQuery{
		db: ssq.db,
	}
}

// 设置Db
func (ssq *SqlServerQuery) SetDb(db *sql.DB) error {
	ssq.db = db
	return nil
}

// 获取SQL语句
func (ssq *SqlServerQuery) Sql() string {
	return ""
}

func (ssq *SqlServerQuery) ExecSql() {

}

// 插入一条记录
func (ssq *SqlServerQuery) Insert() *SqlServerQuery {
	return ssq
}

// 更新一条记录
func (ssq *SqlServerQuery) Update() *SqlServerQuery {
	return ssq
}

// 删除一条记录
func (ssq *SqlServerQuery) Delete() *SqlServerQuery {
	return ssq
}

// 查询条件
func (ssq *SqlServerQuery) Where() *SqlServerQuery {
	return ssq
}

// 获取一条记录
func (ssq *SqlServerQuery) FetchRow() {

}

// 获取多条记录
func (ssq *SqlServerQuery) FetchAll() {

}
