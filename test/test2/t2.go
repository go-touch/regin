package main

import (
	"fmt"
	"strings"
)

func test(any map[string]interface{}) map[string]interface{} {
	tmpMap := map[string]interface{}{}
	for key, value := range any {
		i := strings.IndexByte(key, '[')
		j := strings.IndexByte(key, ']')

		if i >= 1 && j >= 1 {
			trimKey := strings.Trim(key, "]")
			k := trimKey[0:i]
			subK := trimKey[i+1:]
			if _, ok := tmpMap[k]; !ok {
				tmpMap[k] = map[string]interface{}{
					subK: value,
				}
			} else {
				tmpMap[k].(map[string]interface{})[subK] = value
			}
			delete(any, key)
		}
	}
	for key, value := range tmpMap {
		any[key] = value
	}
	return any
}

func main() {
	data := map[string]interface{}{
		"username":    "admin",
		"group[age]":  10,
		"group[name]": "张三",
	}

	fmt.Println(data)
	fmt.Println(test(data))
}
