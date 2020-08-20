package main

import (
	"github.com/go-touch/regin/app/db/query"
)

func main() {
	mysql := new(query.MysqlQuery).Clone()
	//mysql.Field("`username`,password`,`address")
	//mysql.Field([]string{"a.menu_name as name", "style", "is_menu"})
	mysql.Table("admin_user as au")
	mysql.Where("id   =      ?", 1)
	mysql.Where("username   =   ?   ","admin")


	mysql.FetchRow()

}
