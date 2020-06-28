package model

import (
	"errors"
	"fmt"
	"reflect"
)

type Table struct {
	Type      reflect.Type  // 反射类型
	Value     reflect.Value // 反射值
	Identify  string        // 数据库标识
	Name      string        // 结构体名称
	TableName string        // 表名
	FieldLen  int           // 字段数
	Field     []*Field      // 字段数组
}



// 获取反射类型
func (t *Table) reflectType(model interface{}) (reflect.Type, error) {
	if t.Type == nil {
		t.Type = reflect.TypeOf(model)
	}
	if t.Type.Kind() != reflect.Ptr || t.Type.Elem().Kind() != reflect.Struct {
		return nil, errors.New("accept an pointer model")
	} else {
		t.Type = t.Type.Elem()
	}
	return t.Type, nil
}

// 工厂方法
func (t *Table) Factory(model interface{}) {

}

func NewTable(model interface{}) (*Table, error) {
	modelType := reflect.TypeOf(model)
	if modelType.Kind() != reflect.Ptr || modelType.Elem().Kind() != reflect.Struct {
		return nil, errors.New("accept an pointer model")
	} else {
		modelType = modelType.Elem()
	}

	//
	table := &Table{
		Type:      modelType,
		Name:      modelType.String(),
		Identify:  "master",
		TableName: "",
		FieldLen:  modelType.NumField(),
		Field:     make([]*Field, 0),
	}

	// 获取表属性
	if value, ok := reflect.New(modelType).Interface().(Model); ok {
		if identify := value.Identify(); identify != "" {
			table.Identify = identify
		}
		if tableName := value.TableName(); tableName != "" {
			table.TableName = tableName
		} else {
			// 拆分
			// modelType.Name()
		}
	}

	// 字段
	table.Field = make([]*Field, table.FieldLen)
	for i := 0; i < table.FieldLen; i++ {
		table.Field[i] = &Field{
			Name: modelType.Field(i).Name,
			Type: modelType.Field(i).Type.String(),
			// Tag: modelType.Field(i).Tag,
		}

		fmt.Println(*table.Field[i])
	}
	return table, nil
}
