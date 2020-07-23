package main

import (
	"fmt"
	"github.com/go-touch/regin/app/db"
)

type Base struct {
}

type AdminUsers struct {
	Base
	Uid      int    `field:"uid"`
	Account  string `field:"account"`
	Username string `field:"username"`
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
			"maxIdleConn":    "100",
			"maxOpenConn":    "100",
		},
	})

	// 查询单条记录
	/*model := db.Model(&AdminUsers{})
	sql := model.FetchRow(func(dao *db.Dao) {
		dao.Where("account", 15801690885)
		//dao.Where("account", 111)
		dao.Field("uid,account,username,avatar,")
		dao.Limit(0, 10)
		dao.Order("uid desc")
		//dao.Sql()
	}).ToStringMap()*/

	/*// 插入一条数据
	model := db.Model(&AdminUsers{})
	sql := model.Insert(func(dao *db.Dao) {
		dao.Where("account", "test1")
		dao.Values(map[string]interface{}{
			"avatar": "123",
			"username": "admin123",
		})
		//dao.Sql()
	})

	fmt.Println(sql.ToAffectedRows())*/

	// 更新一条数据 -- 错误示例
	/*model := db.Model(&AdminUsers{})
	result := model.Update(func(dao *db.Dao) {
		dao.Where("account", "15116980818")
		dao.Set(map[string]interface{}{
			"mobile": "123",
			"remark": []string{"1"},
		})
		//dao.Sql()
	})

	if result.ToError() != nil{
		panic(result.ToError())
	}
	fmt.Println(result.ToError())*/

	// 更新一条数据 -- 错误示例
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

	// 查询
	ret2 := model.FetchAll(func(dao *db.Dao) {
		dao.Where("username", "admin1")
	})
	fmt.Printf("ret2的数据:%v\n", ret2.ToStringMapSlice())

	ret3 := db.Model(&AdminUsers{}).Update(func(dao *db.Dao) {
		dao.Values(map[string]interface{}{
			"account": "测试事务2",
		})
		dao.Where("username", "admin2")
	})
	fmt.Printf("ret3的数据:%v\n", ret3.ToAffectedRows())
	model.Rollback()

	/*model := db.Model(&AdminUsers{})
	model.Test()
	model.Begin()
	model.Test()

	db.Model(&AdminUsers{}).Test()
	db.Model(&AdminUsers{}).Test()
	db.Model(&AdminUsers{}).Test()*/




}
