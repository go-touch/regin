package main

import (
	"fmt"
	"github.com/go-touch/regin/app/db"
)

type AdminUsers struct {
	Uid      int    `field:"uid"`
	Account  string `field:"account"`
	Username string `field:"username"`
}

type AdminRoles struct {
	Id   int    `field:"id"`
	Name string `field:"name"`
	Desc string `field:"desc"`
}

// 数据库标识(此方法可重构,用于切换数据库,默认master)
func (this AdminUsers) Identify() string {
	return "master"
}

// 数据库表名(此方法可重构,用于切换数据表)
func (this AdminUsers) TableName() string {
	return "admin_users"
}

func main() {
	db.Config.Init(map[string]interface{}{
		"master": map[string]string{
			"driverName":     "mysql",
			"dataSourceName": "vdong:qO39eudVDA@tcp(39.106.157.226:3306)/plus_operation?charset=utf8",
			"maxIdleConn":    "10",
			"maxOpenConn":    "100",
		},
	})

	// 测试
	model := db.Model(&AdminUsers{})
	model.Begin() // 开始事务
	ret1 := model.Update(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{
			"account": "测试事务1",
		})
		dao.Where("username", "admin1")
		//dao.Sql()
	})
	fmt.Printf("ret1的数据:%v\n", ret1.ToAffectedRows())

	_ =  model.Update(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{
			"account": "测试事务1222",
		})
		dao.Where("username", "admin1")
		//dao.Sql()
	})


	// 查询
	ret2 := db.Model(&AdminRoles{}).FetchAll(func(dao *db.Dao) {

	})
	fmt.Printf("ret2的数据:%v\n", ret2.ToStringMapSlice())

/*	// 查询
	ret4 := db.Model(&AdminUsers{}).FetchRow(func(dao *db.Dao) {
		dao.Where("username", "admin2")
	})
	fmt.Printf("ret3的数据:%v\n", ret4.ToStringMap())
*/
	// 更新
	ret3 := db.Model(&AdminRoles{}).Update(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{
			"name": "测试223",
		})
		dao.Where("id", "115")
	})
	fmt.Printf("ret3的数据:%v\n", ret3.ToAffectedRows())

	ret4 := db.Model(&AdminUsers{}).FetchAll(func(dao *db.Dao) {
		dao.Where("username", "admin2")
	})
	fmt.Printf("ret4的数据:%v\n", ret4.ToStringMapSlice())

	// model.Rollback()

}
