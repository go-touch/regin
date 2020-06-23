package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/go-touch/regin/base"
)

type RedisInterface interface {
	Identify() string                                        // Redis库标识(此方法可重构,用于切换库,默认master)
	Init()                                                   // Model初始化方法(如有自定义函数,可用此函数初始化Method Map)
	Self(model RedisInterface) RedisInterface                // 设置自身指针(此方法禁止重构)
	Pool() *redis.Pool                                       // 获取Redis连接池对象(此方法禁止重构)
	Connect() redis.Conn                                     // 获取一个Redis连接(此方法禁止重构)
	Call(methodName string, args ...interface{}) interface{} // 调用自定义方法(此方法禁止重构)
}

type RedisModel struct {
	RedisInterface
	poll   *redis.Pool
	self   RedisInterface
	Method map[string]base.UserFunc
}

// 设置Redis自身
func (rm *RedisModel) Init() {
	rm.Method = make(map[string]base.UserFunc)
}

// Redis库标识
func (rm *RedisModel) Identify() string {
	return "master"
}

// 设置Redis自身
func (rm *RedisModel) Self(model RedisInterface) RedisInterface {
	rm.self = model
	return model
}

// 调用自定义方法
func (rm *RedisModel) Call(methodName string, args ...interface{}) interface{} {
	if method, ok := rm.Method[methodName]; ok {
		return method(args...)
	}
	return ""
}

// 获取Redis连接池对象
func (rm *RedisModel) Pool() *redis.Pool {
	redisPool, err := Redis.Pool(rm.self.Identify())
	if err != nil {
		panic(err.Error())
	}
	return redisPool
}

// 获取Redis连接
func (rm *RedisModel) Connect() redis.Conn {
	return rm.Pool().Get()
}