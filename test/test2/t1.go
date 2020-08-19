package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]interface{}{
		"a": "admin",
		"b": []string{"a", "v"},
	}

	test, err := json.Marshal(&m)

	fmt.Println(test, "----", err)

	fmt.Println(string(test))
}
