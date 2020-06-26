package service

import (
	"github.com/go-touch/regin/app/cache"
	"github.com/go-touch/regin/app/db"
)

// 初始化扩展包
func (a *Application) initExpand() {
	cache.Redis.Init(a.GetConfig("redis").ToAnyMap()) // 初始化Redis
	db.Query.Init(a.GetConfig("database").ToAnyMap()) // 初始化Database
}
