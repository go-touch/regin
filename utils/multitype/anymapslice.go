package multitype

// Slice type:  key is index,value is map[string]interface{}.
type AnyMapSlice []map[string]interface{}

// Set value.
func (ams *AnyMapSlice) Set(key int, value map[string]interface{}) {
	(*ams)[key] = value
}

// Get value.
func (ams *AnyMapSlice) Get(key ...int) *AnyValue {
	if key == nil {
		return Eval(*ams)
	} else if !ams.KeyExist(key[0]) {
		return Eval(nil)
	}
	return Eval((*ams)[key[0]])
}

// Check key is exist.
func (ams *AnyMapSlice) KeyExist(key int) bool {
	for k := range *ams {
		if k == key {
			return true
		}
	}
	return false
}
