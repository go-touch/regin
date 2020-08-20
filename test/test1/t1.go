package main

import (
	"fmt"
	"github.com/go-touch/regin/app/db"
)

type Base struct {
}

type PlusMenu2 struct {
	Base
	Id       int    `field:"id"`
	Pid      int    `field:"pid"`
	MenuName string `field:"menu_name"`
}

// 数据库标识(此方法可重构,用于切换数据库,默认master)
func (this PlusMenu2) Identify() string {
	return "master"
}

// 数据库表名(此方法可重构,用于切换数据表)
func (this PlusMenu2) TableName() string {
	return "plus_menu2"
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
	// 更新一条数据 -- 错误示例
	model := db.Model(&PlusMenu2{})
	ret := model.Update(func(dao *db.Dao) {
		dao.Where("id", "94")
		dao.Values(map[string]interface{}{
			"pid": "1",
		})
		//dao.Sql()
	})
	fmt.Println(ret.ToError())
	fmt.Println(ret.ToString())
	fmt.Println(ret.ToAffectedRows())
}
