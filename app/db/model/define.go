package model

import "reflect"

type Model interface {
	Identify() string  // 数据库标识(此方法可重构,用于切换数据库,默认master)
	TableName() string // 数据库表名(此方法可重构,用于切换数据表)
}

// 获取结构体名称
func GetName(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		return ""
	}
	return t.Elem().String()
}
