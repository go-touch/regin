package tag

// 结构体属性接口
type BaseTag interface {
	SetValue(value string)
	Verify(key string, vMap *map[string]interface{}) *Result
}

// 校验结果
type Result struct {
	Status bool
	Msg    string
}
