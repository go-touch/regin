package validator

import (
	"errors"
)

// Form Container
type FormContainer map[string]*FormHandle
var fc *FormContainer

func init() {
	fc = &FormContainer{}
}

// 获取 Form
func (fc *FormContainer) Get(key string) (*FormHandle, error) {
	if form, ok := (*fc)[key]; ok {
		return form, nil
	}
	return nil, errors.New("this model '" + key + "' is not registered")
}

// 设置 Form
func (fc *FormContainer) Set(key string, form *FormHandle) {
	(*fc)[key] = form
}

// 注册 Form
func RegisterForm(userModel interface{}, alias ...string) {
	key := ""
	formHandle := NewFormHandle(userModel)
	if alias != nil && alias[0] != "" {
		key = alias[0]
	} else {
		key = formHandle.Name
	}
	if _, err := fc.Get(key); err == nil {
		return
	} else {
		formHandle.Init()
		fc.Set(key, formHandle)
	}
}

// 创建 Form
func Form(userModel interface{}) *FormHandle {
	if key, ok := userModel.(string); ok {
		if storage, err := fc.Get(key); err == nil {
			return storage
		} else {
			panic(err)
		}
	}
	formHandle := NewFormHandle(userModel)
	if storage, err := fc.Get(formHandle.Name); err == nil {
		return storage
	} else {
		formHandle.Init()
		fc.Set(formHandle.Name, formHandle)
	}
	return formHandle
}
