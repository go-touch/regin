package base

type DataType interface {
	Set(key string, value interface{})
	Get(key string) *AnyValue
}

// 预定义常见数据类型
type AnyMap map[string]interface{}        // [MapType] key is string,value is 任意类型
type StringMap map[string]string          // [MapType] key is string,value is string type
type IntMap map[string]int                // [MapType] key is string,value is int type
type StringSliceMap map[string][]string   // [MapType] key is string,value is string Slice type
type GeneralMap map[string]AppAction      // [MapType] key is string,value is AppAction type
type AnySlice []interface{}               // [SliceType] key is index,value为任意类型
type StringMapSlice []map[string]string   // [SliceType] key is index,value为(key为string,value为string)的map
type AnyMapSlice []map[string]interface{} // [SliceType] key is index,value为(key为string,value为任意类型)的map

// 设置
func (am *AnyMap) Set(key string, value interface{}) {
	(*am)[key] = value
}

// 读取
func (am *AnyMap) Get(key ...string) *AnyValue {
	if key == nil || key[0] == "" {
		return Eval(*am)
	} else if value, ok := (*am)[key[0]]; ok {
		return Eval(value)
	}
	return Eval("")
}