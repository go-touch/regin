package utils

import "fmt"

type ErrorFormat struct {
	error
	errorMsg string
}

// 获取一个实现了error接口的ErrorFormat结构体
func MakeError(msg string, errorMsg ...interface{}) error {
	return &ErrorFormat{
		errorMsg: fmt.Sprintf(msg, errorMsg...),
	}
}

// 重载Error方法
func (ef *ErrorFormat) Error() string {
	return ef.errorMsg
}
