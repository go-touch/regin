package base

import "io"

/**************************************** 数据类型 - 结构体Result ****************************************/
// 用户输出函数
type UserOutput func(writer io.Writer) bool

// 响应结构体
type Result struct {
	Type       string            // 可选值为:String、Json、Html、Stream
	header     map[string]string // 头信息
	UserOutput UserOutput        // 用户可控制输出函数
	Page       string            // 响应页面(Type = Html时必填)
	Status     int               // 状态码 200正常状态
	Msg        string            // 提示消息
	Data       AnyMap            // 业务数据
}

// 定义RespResult
var ResultInvoker *Result

func init() {
	ResultInvoker = &Result{}
}

// 设置 header
func (r *Result) SetHeader(key string, value string) {
	r.header[key] = value
}

// 获取 header
func (r *Result) GetHeader() map[string]string {
	return r.header
}

// Business data method.
func (r *Result) SetData(key string, value interface{}) {
	r.Data[key] = value
}

// business data method - 获取Data
func (r *Result) GetData(key string) interface{} {
	return r.Data[key]
}

// 创建Json result
func (r *Result) CreateJson(status int, msg string) *Result {
	return &Result{
		Type:   "Json",
		header: map[string]string{},
		Page:   "",
		Status: status,
		Msg:    msg,
		Data:   AnyMap{"code": 0, "msg": "", "data": ""},
	}
}

// 创建Json result
func (r *Result) CreateHtml(page string, status int, msg string) *Result {
	return &Result{
		Type:   "Html",
		header: map[string]string{},
		Page:   page,
		Status: status,
		Msg:    msg,
		Data:   AnyMap{},
	}
}

// 创建 Stream Result
func StreamResult() *Result {
	return &Result{
		Type:   "Stream",
		header: map[string]string{},
		Page:   "",
		Status: 200,
		Msg:    "",
		Data:   AnyMap{"code": 0, "msg": "", "data": ""},
	}
}