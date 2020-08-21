package query

import (
	"database/sql"
	"time"
)

// Define common struct.
type Combine struct {
	BaseQuery
	db      *sql.DB
	tx      *sql.Tx
	Runtime Runtime
}

// Runtime data struct.
type Runtime struct {
	Sql      string        // sql语句
	Args     []interface{} // 参数
	ExecTime string        // 执行时间
	Err      error         // 运行错误
}

// Set Db.
func (c *Combine) SetDb(db *sql.DB) {
	c.db = db
}

// Set Tx.
func (c *Combine) SetTx(tx *sql.Tx) {
	c.tx = tx
}

// Unset Tx.
func (c *Combine) UnsetTx() {
	c.tx = nil
}

// Get Tx.
func (c *Combine) GetTx() *sql.Tx {
	return c.tx
}

// 获取运行期间数据
func (c *Combine) GetDuration() Runtime {
	return c.Runtime
}

// 查询一条记录
func (c *Combine) QueryRow(sql string, args ...interface{}) (row *sql.Row) {
	startTime := time.Now()
	if c.tx != nil {
		row = c.tx.QueryRow(sql, args...)
	} else {
		row = c.db.QueryRow(sql, args...)
	}
	// Runtime data.
	c.Runtime.Sql = sql
	c.Runtime.Args = args
	c.Runtime.ExecTime = time.Since(startTime).String()
	c.Runtime.Err = nil
	return row
}

// 查询多条记录
func (c *Combine) QueryAll(sql string, args ...interface{}) (rows *sql.Rows, err error) {
	startTime := time.Now()
	if c.tx != nil {
		rows, err = c.tx.Query(sql, args...)
	} else {
		rows, err = c.db.Query(sql, args...)
	}
	// Runtime data.
	c.Runtime.Sql = sql
	c.Runtime.Args = args
	c.Runtime.ExecTime = time.Since(startTime).String()
	c.Runtime.Err = err
	return rows, err
}

// 插入[更新][删除]n条记录
func (c *Combine) Exec(sql string, args ...interface{}) (result sql.Result, err error) {
	startTime := time.Now()
	if c.tx != nil {
		result, err = c.tx.Exec(sql, args...)
	} else {
		result, err = c.db.Exec(sql, args...)
	}
	// Runtime data.
	c.Runtime.Sql = sql
	c.Runtime.Args = args
	c.Runtime.ExecTime = time.Since(startTime).String()
	c.Runtime.Err = err
	return result, err
}

// Begin starts a transaction.
func (c *Combine) Begin() {
	if tx, err := c.db.Begin(); err != nil {
		panic(err.Error())
	} else {
		c.tx = tx
	}
}

//  Commit commits the transaction.
func (c *Combine) Commit() {
	err := c.tx.Commit()
	if err != nil {
		panic(err.Error())
	}
	c.tx = nil
}

//  Rollback aborts the transaction.
func (c *Combine) Rollback() {
	err := c.tx.Rollback()
	if err != nil {
		panic(err.Error())
	}
	c.tx = nil
}
