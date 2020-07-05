package main

import (
	"fmt"
	"github.com/go-touch/regin/app/cache"
)

type RedisModel struct {
}

func (rm *RedisModel) Identify() string {
	return "plus_center.master"
}

func main() {
	cache.Redis.Init(map[string]interface{}{
		"plus_center": map[string]interface{}{
			"master": map[string]interface{}{
				"host":        "127.0.0.1:6379",
				"MaxIdle":     "16",
				"MaxActive":   "32",
				"IdleTimeout": "120",
			},
		},
	})

	model := cache.RedisModel(new(RedisModel))
	ret := model.Command("set", "name", "admin`1232")

	fmt.Println(model)
	fmt.Println(ret)
	fmt.Printf("返回值类型:%T\n",ret.ToValue())

}
