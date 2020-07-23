package query

import (
	"database/sql"
)

type Combine struct {
	BaseQuery
	db *sql.DB
	tx *sql.Tx
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

// 查询一条记录
func (c *Combine) QueryRow(sql string, args ...interface{}) *sql.Row {
	if c.tx != nil {
		return c.tx.QueryRow(sql, args...)
	}
	return c.db.QueryRow(sql, args...)
}

// 查询多条记录
func (c *Combine) QueryAll(sql string, args ...interface{}) (*sql.Rows, error) {
	if c.tx != nil {
		return c.tx.Query(sql, args...)
	}
	return c.db.Query(sql, args...)
}

// 插入[更新][删除]n条记录
func (c *Combine) Exec(sql string, args ...interface{}) (sql.Result, error) {
	if c.tx != nil {
		return c.tx.Exec(sql, args...)
	}
	return c.db.Exec(sql, args...)
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
