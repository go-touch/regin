package cache

import (
	"github.com/garyburd/redigo/redis"
	"strings"
	"strconv"
	"errors"
	"time"
)

type RedisDispatch struct {
	config      interface{}
	container   map[string]*redis.Pool
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
}

// Define RedisDispatch.
var Redis *RedisDispatch

func init() {
	Redis = &RedisDispatch{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		container:   make(map[string]*redis.Pool),
	}
}

// Init config.
func (ed *RedisDispatch) Init(config map[string]interface{}) error {
	ed.config = config
	return nil
}

// 读取配置
func (ed *RedisDispatch) GetConfig(arg string) (config map[string]string, err error) {
	config = make(map[string]string)

	// 解析参数
	argGroup := strings.Split(arg, ".")

	if argGroup[0] == "" {
		return config, errors.New("the database's identify is not set")
	}

	// 遍历处理
	configTmp := ed.config

	for _, key := range argGroup {
		switch t := configTmp.(type) {
		case nil:
			configTmp = map[string]string{}
		case map[string]string:
			break
		case map[string]interface{}:
			if value, ok := configTmp.(map[string]interface{})[key]; ok {
				configTmp = value
			} else {
				configTmp = map[string]string{}
			}
		case []interface{}:
			if intKey, err := strconv.Atoi(key); err == nil {
				configTmp = configTmp.([]interface{})[intKey]
			}
		default:
			_ = t
			configTmp = map[string]string{}
			break
		}
	}

	// 类型断言并处理
	if value, ok := configTmp.(map[string]interface{}); ok {
		for k, v := range value {
			config[k] = v.(string)
		}
	} else { // 类型断言并处理
		value := configTmp.(map[string]string)

		if value != nil {
			config = value
		} else {
			config = map[string]string{}
		}
	}
	return config, nil
}

// Get redis pool object.
func (ed *RedisDispatch) Pool(identify string) (redisPool *redis.Pool, err error) {
	// Get from container, if exists then return.
	if redisPool, ok := ed.container[identify]; ok {
		return redisPool, err
	}

	// Get config
	config, err := ed.GetConfig(identify)
	if err != nil {
		return redisPool, err
	}

	// 配置参数校验
	connParam := map[string]string{"host": "", "password": "", "db": ""}
	for key := range connParam {
		if v, ok := config[key]; !ok {
			return redisPool, errors.New("the config's " + key + " is not set")
		} else {
			connParam[key] = v
		}
	}

	// 服务地址
	if connParam["host"] == "" {
		return redisPool, errors.New("the config's host cannot be null")
	}

	// 数据库库标
	dbIndex := 0
	if connParam["db"] != "" {
		if v, err := strconv.Atoi(connParam["db"]); err == nil {
			dbIndex = v
		}
	}

	// 创建连接池
	redisPool = &redis.Pool{
		MaxIdle:     ed.MaxIdle,
		MaxActive:   ed.MaxActive,
		IdleTimeout: time.Duration(ed.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			dialOption := []redis.DialOption{
				redis.DialReadTimeout(time.Duration(1000) * time.Millisecond),
				redis.DialWriteTimeout(time.Duration(1000) * time.Millisecond),
				redis.DialConnectTimeout(time.Duration(1000) * time.Millisecond),
				redis.DialDatabase(dbIndex),
			}
			if connParam["password"] != "" {
				dialOption = append(dialOption, redis.DialPassword(connParam["password"]))
			}
			return redis.Dial("tcp", connParam["host"], dialOption...)
		},
	}

	// 存储
	ed.container[identify] = redisPool
	return redisPool, err
}
