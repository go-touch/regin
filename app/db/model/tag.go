package model

type Tag map[string]string
var TagAttribute *Tag

func init() {
	TagAttribute = &Tag{
		"field": "",
	}
}

// 是否存在
func (t *Tag) Clone() *Tag {
	return &Tag{
		"field": (*t)["field"],
	}
}

// 设置数据
func (t *Tag) Set(key string, value string) {
	(*t)[key] = value
}

// 是否存在
func (t *Tag) InArray(aType string) bool {
	for _, v := range *t {
		if aType == v {
			return true
		}
	}
	return false
}

// 数据表真实字段值
func (t *Tag) GetField() string {
	return (*t)["field"]
}
