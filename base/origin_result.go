package base

/**************************************** 数据类型 - 结构体Result ****************************************/
// 响应结构体
type Result struct {
	Type   string // 可选值为:String、Json、Html、
	Page   string // 响应页面(Type = Html时必填)
	Status int    // 状态码 200正常状态
	Msg    string // 提示消息
	Data   AnyMap // 业务数据
}

// 定义RespResult
var ResultInvoker *Result

func init() {
	ResultInvoker = &Result{}
}

// 创建Json result
func (r *Result) CreateJson(status int, msg string) *Result {
	return &Result{
		Type:   "Json",
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
		Page:   page,
		Status: status,
		Msg:    msg,
		Data:   AnyMap{},
	}
}

// Business data method.
func (r *Result) SetData(key string, value interface{}) error {
	r.Data[key] = value
	return nil
}

// business data method - 获取Data
func (r *Result) GetData(key string) interface{} {
	return r.Data[key]
}

