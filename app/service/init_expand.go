package service

import (
	"github.com/go-touch/regin/app/cache"
	"github.com/go-touch/regin/app/core"
	"github.com/go-touch/regin/app/db"
)

// 初始化扩展包
func (a *Application) initExpand() {
	cache.Redis.Init(a.GetConfig("redis").ToAnyMap())  // 初始化Redis
	db.Config.Init(a.GetConfig("database").ToAnyMap()) // 初始化Database
	db.Config.InitLogWriter(func(anyMap map[string]interface{}) { // 数据库运行日志
		_ = a.logger.Local(core.Info, anyMap) // SQL执行日志
	})
}