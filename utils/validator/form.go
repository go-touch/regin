package validator

import (
	"errors"
	"github.com/go-touch/regin/utils/validator/tag"
	"reflect"
	"strings"
)

// Form 字段类型限定
type FieldType []string

// 是否存在
func (ft *FieldType) InArray(t string) bool {
	for _, v := range *ft {
		if t == v {
			return true
		}
	}
	return false
}

// 获取一个字段类型实例 FieldType
func NewFieldType() *FieldType {
	return &FieldType{"int", "int8", "int16", "int32", "int64", "string"}
}

// Form载体
type FormHandle struct {
	Model    interface{}       // 数据模型结构体指针
	Type     reflect.Type      // 反射类型
	Name     string            // 结构体名称
	FieldNum int               // 字段数
	FieldMap map[string]*Field // 字段map
}

// 获取一个 Form实例
func NewFormHandle(model interface{}) *FormHandle {
	formHandle := &FormHandle{
		Model: model,
		Type:  reflect.TypeOf(model),
	}
	if formHandle.Type.Kind() != reflect.Ptr || formHandle.Type.Elem().Kind() != reflect.Struct {
		panic(errors.New("accept an pointer model"))
	} else {
		formHandle.Type = formHandle.Type.Elem()
		formHandle.Name = formHandle.Type.String()
		formHandle.FieldNum = formHandle.Type.NumField()
		formHandle.FieldMap = map[string]*Field{}
	}
	return formHandle
}

// 初始化
func (mh *FormHandle) Init() {
	for i := 0; i < mh.FieldNum; i++ {
		fieldType := mh.Type.Field(i).Type.String()
		if NewFieldType().InArray(fieldType) {
			field := &Field{
				Name: mh.Type.Field(i).Name,
				Type: fieldType,
				Tags: map[string]tag.BaseTag{
					"require": tag.MakeRequire(false),
					"length":  tag.MakeLength(false),
				},
			}
			// 遍历处理
			fieldKey := mh.Capitalize(field.Name)
			if value, ok := mh.Type.Field(i).Tag.Lookup("key"); ok {
				fieldKey = value
			}
			// 设置标签
			for key, tagElement := range field.Tags {
				if value, ok := mh.Type.Field(i).Tag.Lookup(key); ok {
					tagElement.SetValue(value)
				}
			}
			mh.FieldMap[fieldKey] = field
		}
	}
}

// 数据校验
func (mh *FormHandle) Verify(vMap *map[string]interface{}) []*tag.Result {
	resultGroup := make([]*tag.Result, 0)
	for key, fieldElement := range mh.FieldMap {
		result := fieldElement.Verify(key, vMap)
		for _, v := range result {
			if v.Status == false {
				resultGroup = append(resultGroup, v)
			}
		}
	}
	return resultGroup
}

// 首字母大写转换成 _
func (mh *FormHandle) Capitalize(str string) string {
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
