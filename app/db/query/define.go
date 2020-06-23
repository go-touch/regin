package query

import (
	"database/sql"
)

type BaseQuery interface {
	Copy() BaseQuery
	SetDb(db *sql.DB) error
	Sql() string
	Query(sql string, args ...interface{}) (result interface{}, err error)
	Begin()    // Begin starts a transaction.
	Commit()   // Commit commits the transaction.
	Rollback() // Rollback aborts the transaction.
}
