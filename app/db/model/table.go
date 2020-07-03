package model

import (
	"errors"
	"reflect"
	"strings"
)

type Table struct {
	Model      interface{}  // 数据模型结构体指针
	Type       reflect.Type // 反射类型
	Name       string       // 结构体名称
	Identify   string       // 数据库标识
	TableName  string       // 表名
	FieldNum   int          // 字段数
	FieldArray []*Field     // 字段数组
}

// 获取反射类型
func (t *Table) Init(model interface{}) (*Table, error) {
	t.Type = reflect.TypeOf(model)
	if t.Type.Kind() != reflect.Ptr || t.Type.Elem().Kind() != reflect.Struct {
		return nil, errors.New("accept an pointer model")
	} else {
		t.Model = model
		t.Type = t.Type.Elem()
		t.Name = t.Type.String()
		t.Identify = "master"
		t.TableName = t.Capitalize(t.Type.Name())
		t.FieldNum = t.Type.NumField()
		t.FieldArray = make([]*Field, 0)
	}
	return t, nil
}

// 工厂方法
func (t *Table) Factory() *Table {
	// 表处理
	reflectModel := reflect.New(t.Type).Interface()
	if value, ok := reflectModel.(Identify); ok {
		if identify := value.Identify(); identify != "" {
			t.Identify = identify
		}
	}
	if value, ok := reflectModel.(TableName); ok {
		if tableName := value.TableName(); tableName != "" {
			t.TableName = tableName
		}
	}

	// 字段处理
	t.FieldArray = make([]*Field, 0)
	for i := 0; i < t.FieldNum; i++ {
		inType := t.Type.Field(i).Type.String()
		if FieldType.InArray(inType) {
			field := &Field{
				name:     t.Type.Field(i).Name,
				trueName: "",
				inType:   inType,
				tag:      TagAttribute.Clone(),
			}
			// 结构体标签
			for k, _ := range *field.tag {
				if value, ok := t.Type.Field(i).Tag.Lookup(k); ok {
					field.tag.Set(k, value)
				}
			}
			if trueName := field.tag.GetField(); trueName != "" {
				field.trueName = trueName
			} else {
				field.trueName = t.Capitalize(field.name)
			}
			t.FieldArray = append(t.FieldArray, field)
		}
	}
	return t
}

// 首字母大写转换成 _
func (t *Table) Capitalize(str string) string {
	newStr := strings.ToLower(string(str[0]))
	value := []rune(str)
	for i := 1; i < len(value); i++ {
		if value[i] >= 65 && value[i] <= 96 {
			newStr += "_" + strings.ToLower(string(value[i]))
		} else {
			newStr += string(value[i])
		}
	}
	return newStr
}

// 获取结构体名称
func (t *Table) GetName() (string, error) {
	if t.Type == nil {
		return "", errors.New("the table struct is not init")
	}
	return t.Type.String(), nil
}

// 获取字段组
func (t *Table) GetField() []*Field {
	return t.FieldArray
}

// 获取表字段列表
func (t *Table) GetTableFields() []string {
	tableFields := make([]string, 0)
	for _, v := range t.FieldArray {
		tableFields = append(tableFields, v.trueName)
	}
	return tableFields
}
