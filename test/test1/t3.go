package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//var username string
	// 数据库连接对象
	db, e := sql.Open("mysql", "vdong:qO39eudVDA@tcp(39.106.157.226:3306)/plus_operation?charset=utf8")
	if e != nil {
		panic(e)
	}
	db.SetMaxOpenConns(2)

	// 查询1
	rows1, _ := db.Query("SELECT username FROM admin_users" + "")
	fmt.Printf("row1的值:%v\n", rows1)

	tx, _ := db.Begin()
	fmt.Printf("开启事务,tx的值:%v\n", tx)

	// 查询2
	rows2, _ := tx.Query("SELECT username FROM admin_users" + "")
	fmt.Printf("rows2的值:%v\n", rows2)
	_ = rows2.Close()

	fmt.Printf("tx的值:%v\n", tx)


	// 查询3
	rows3, _ := tx.Query("SELECT username FROM admin_users" + "")
	fmt.Printf("rows3的值:%v\n", rows3)
	_ = rows3.Close()

	_ = tx.Commit()

	fmt.Printf("tx的值:%v\n", tx)

	/*fmt.Println(rows)
	_ = rows.Close()

	tx, _ := db.Begin()

	tx.Exec("insert into admin_users (username) values ('测试一下')")

	// 查询2
	rows2, err2 := db.Query("SELECT * FROM admin_nodes")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(rows2)
	_ = rows2.Close()

	// 查询3
	rows3, err3 := db.Query("SELECT * FROM admin_roles")
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(rows3)
	_ = rows3.Close()

	_ = tx.Commit()*/
}
