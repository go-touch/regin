package multitype

// Slice type: key is index,value is any type.
type AnySlice []interface{}

// Set value.
func (as *AnySlice) Set(key int, value interface{}) {
	(*as)[key] = value
}

// Get value.
func (as *AnySlice) Get(key ...int) *AnyValue {
	if key == nil {
		return Eval(*as)
	} else if !as.KeyExist(key[0]) {
		return Eval(nil)
	}
	return Eval((*as)[key[0]])
}

// Check key is exist.
func (as *AnySlice) KeyExist(key int) bool {
	for k := range *as {
		if k == key {
			return true
		}
	}
	return false
}
