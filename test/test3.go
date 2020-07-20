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
	ret := model.Command("incrby", "bind-phone-index_msg_limit_16619709811",20)

	fmt.Println(ret.Value())

}
