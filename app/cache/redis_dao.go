package cache

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"github.com/go-touch/regin/base"
)

type Identify interface {
	Identify() string // Redis库标识(此方法可重构,用于切换库,默认master)
}

type RedisDao struct {
	pool *redis.Pool
}

// 获取RedisDao对象
func RedisModel(model interface{}) *RedisDao {
	var redisPool *redis.Pool
	if ptr, ok := model.(Identify); !ok {
		panic(errors.New("need an pointer model"))
	} else {
		redisPool = Redis.Pool(ptr.Identify())
	}
	return &RedisDao{pool: redisPool}
}

// 获取连接池对象
func (rd *RedisDao) Pool() *redis.Pool {
	return rd.pool
}

// 执行redis命令
func (rd *RedisDao) Command(name string, args ...interface{}) *base.AnyValue {
	defer func() { _ = rd.pool.Close() }()
	result, err := rd.Pool().Get().Do(name, args...)
	if err != nil {
		return base.Eval(err)
	}
	return base.Eval(result)
}
