package model

type Identify interface {
	Identify() string // 数据库标识(此方法可重构,用于切换数据库,默认master)
}

type TableName interface {
	TableName() string // 数据库表名(此方法可重构,用于切换数据表)
}

// 表字段类型
type FieldTypes []string
var FieldType *FieldTypes

func init() {
	FieldType = &FieldTypes{"int", "int8", "int16", "int32", "int64", "string"}
}

// 是否设置
func (ft *FieldTypes) InArray(t string) bool {
	for _, v := range *ft {
		if t == v {
			return true
		}
	}
	return false
}
