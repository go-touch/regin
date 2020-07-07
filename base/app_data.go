package base

type Data struct {
	Code int
	Msg  string
	Data interface{}
}

// 创建Data
func MakeData() *Data {
	return &Data{Code: 0, Msg: "", Data: ""}
}

// 设置code
func (d *Data) SetCode(code int) {
	d.Code = code
}

// 设置code
func (d *Data) SetMsg(msg string) {
	d.Msg = msg
}

// 设置code
func (d *Data) SetData(data interface{}) {
	d.Data = data
}

// 读取data
func (d *Data) GetData() *AnyValue {
	return Eval(d.Data)
}