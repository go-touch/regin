package validator

import "github.com/go-touch/regin/utils/validator/tag"

type Field struct {
	Name string
	Type string
	Tags map[string]tag.BaseTag
}

// 校验
func (f *Field) Verify(key string, vMap *map[string]interface{}) []*tag.Result {
	resultGroup := make([]*tag.Result, 0)
	for _, tagElement := range f.Tags {
		resultGroup = append(resultGroup, tagElement.Verify(key, vMap))
	}
	return resultGroup
}