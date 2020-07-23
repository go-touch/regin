package main

import (
	"fmt"
	"github.com/go-touch/regin/app/db"
)

type BaseModel struct {
}

// 数据库标识(此方法可重构,用于切换数据库,默认master)
func (this BaseModel) Identify() string {
	return "master"
}

type AdminUsers2 struct {
	BaseModel
	Uid      int    `field:"uid"`
	Account  string `field:"account"`
	Username string `field:"username"`
}

// 数据库表名(此方法可重构,用于切换数据表)
func (this AdminUsers2) TableName() string {
	return "admin_users"
}

type AdminRoles2 struct {
	BaseModel
	Id   int    `field:"id"`
	Name string `field:"name"`
	Desc string `field:"desc"`
}

// 数据库标识(此方法可重构,用于切换数据库,默认master)
func (this AdminRoles2) TableName() string {
	return "admin_roles"
}

func main() {
	db.Config.Init(map[string]interface{}{
		"master": map[string]string{
			"driverName":     "mysql",
			"dataSourceName": "vdong:qO39eudVDA@tcp(39.106.157.226:3306)/plus_operation?charset=utf8",
			"maxIdleConn":    "100",
			"maxOpenConn":    "1",
		},
	})

	dao := db.Model(&AdminUsers2{})
	fmt.Printf("dao的值为:%v\n", dao.GetQuery())

	// 开启事务
	dao.Begin()
	fmt.Printf("dao的值为:%v\n", dao.GetQuery())

	// 查询
	ret1 := dao.Query("SELECT uid,username,account FROM admin_users limit 1")
	fmt.Printf("ret1的值为:%v\n", ret1)

	// 查询
	ret2 := dao.FetchAll(func(dao *db.Dao) {
		dao.Limit(1)
	})
	fmt.Printf("ret2的值为:%v\n", ret2)
	
	// 查询
	dao2 := db.Model(&AdminRoles2{}).Tx(dao)
	fmt.Printf("dao2的值为:%v\n", dao2.GetQuery())

	ret3 := dao2.FetchAll(func(dao *db.Dao) {
		dao.Limit(1)
	})
	fmt.Printf("ret3的值为:%v\n", ret3)
	fmt.Printf("dao2的值为:%v\n", dao2.GetQuery())

	// 插入数据
	ret4 := dao.Insert(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{
			"username": "张小三",
		})
	})
	fmt.Printf("ret4的值为:%v\n", ret4)
	fmt.Printf("dao的值为:%v\n", dao.GetQuery())

	// 插入数据
	ret5 := dao2.Tx(dao).Insert(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{
			"name": "角色张小三",
		})
	})
	fmt.Printf("ret5的值为:%v\n", ret5)
	fmt.Printf("dao的值为:%v\n", dao.GetQuery())

	dao.Commit()

	fmt.Printf("dao的值为:%v\n", dao.GetQuery())
}
