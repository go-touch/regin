package cache

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

type RedisCache struct {
	Abstract
	container map[string]*redis.Pool
}

// Define RedisDispatch.
var Redis *RedisCache

func init() {
	Redis = &RedisCache{
		container: make(map[string]*redis.Pool),
	}
}

// Get redis pool object.
func (rc *RedisCache) Pool(identify string) *redis.Pool {
	if redisPool, ok := rc.container[identify]; ok {
		return redisPool
	}
	// Get config
	config, err := rc.GetConfig(identify)
	if err != nil {
		panic(err)
	}

	// 连接参数
	paramString := map[string]string{"host": "", "password": ""}
	for key := range paramString {
		if value, ok := config[key]; ok {
			paramString[key] = value
		}
	}

	// Redis配置参数
	paramInt := map[string]int{"db": 0, "MaxIdle": 0, "MaxActive": 0, "IdleTimeout": 0}
	for key := range paramInt {
		if value, ok := config[key]; ok {
			if v, err := strconv.Atoi(value); err == nil {
				paramInt[key] = v
			}
		}
	}

	// 服务地址
	if paramString["host"] == "" {
		panic(errors.New("the config's host can't be empty"))
	}

	// 创建连接池
	rc.container[identify] = &redis.Pool{
		MaxIdle:     paramInt["MaxIdle"],
		MaxActive:   paramInt["MaxActive"],
		IdleTimeout: time.Duration(paramInt["IdleTimeout"]),
		Dial: func() (redis.Conn, error) {
			dialOption := []redis.DialOption{
				redis.DialReadTimeout(time.Duration(paramInt["IdleTimeout"]) * time.Millisecond),
				redis.DialWriteTimeout(time.Duration(paramInt["IdleTimeout"]) * time.Millisecond),
				redis.DialConnectTimeout(time.Duration(paramInt["IdleTimeout"]) * time.Millisecond),
				redis.DialDatabase(paramInt["db"]),
			}
			if paramString["password"] != "" {
				dialOption = append(dialOption, redis.DialPassword(paramString["password"]))
			}
			return redis.Dial("tcp", paramString["host"], dialOption...)
		},
	}
	return rc.container[identify]
}
