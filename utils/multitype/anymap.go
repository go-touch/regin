package multitype

//  key is string,value is any type.
type AnyMap map[string]interface{}

// Set value.
func (am *AnyMap) Set(key string, value interface{}) {
	(*am)[key] = value
}

// Get value.
func (am *AnyMap) Get(key ...string) *AnyValue {
	if key == nil {
		return Eval(*am)
	} else if value, ok := (*am)[key[0]]; ok {
		return Eval(value)
	}
	return Eval(nil)
}