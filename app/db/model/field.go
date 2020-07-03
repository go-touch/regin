package model

// 字段结构体
type Field struct {
	name     string // 字段名称
	trueName string // 真实字段名称
	inType   string // 字段类型
	tag      *Tag   // 字段标签
}

// 获取字段名称
func (f *Field) GetName() string {
	return f.name
}

// 获取真实字段名称
func (f *Field) GetTrueName() string {
	return f.trueName
}

// 获取字段类型
func (f *Field) GetType() string {
	return f.inType
}
