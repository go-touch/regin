package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-touch/regin/app/db/query"
	"strconv"
	"errors"
	"strings"
)

type QueryAdapter struct {
	container map[string]query.BaseQuery
	config    interface{}
}

// 定义QueryDispatcher
var Query *QueryAdapter

func init() {
	Query = &QueryAdapter{
		container: map[string]query.BaseQuery{
			"mysql":      &query.MysqlQuery{},
			"sql_server": &query.SqlServerQuery{},
		},
	}
}

// 载入配置
func (qa *QueryAdapter) LoadConfig(config map[string]interface{}) error {
	qa.config = config
	return nil
}

// 读取配置
func (qa *QueryAdapter) GetConfig(arg string) (config map[string]string, err error) {
	config = make(map[string]string)

	// 解析参数
	argGroup := strings.Split(arg, ".")

	if argGroup[0] == "" {
		return config, errors.New("the database's identify is not set")
	}

	// 遍历处理
	configTmp := qa.config

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

// 获取查询器
func (qa *QueryAdapter) GetAdapter(identify string) (adapter query.BaseQuery, err error) {
	// 读取容器
	if queryAdapter, ok := qa.container[identify]; ok {
		return queryAdapter.Copy(), nil
	}

	// 读取配置
	config, err := qa.GetConfig(identify)
	if err != nil {
		return adapter, err
	}

	// 获取驱动name
	driverName, ok := config["driverName"]
	if !ok {
		return adapter, errors.New("the config's driverName is not set")
	}

	// 获取驱动dsn
	dataSourceName, ok := config["dataSourceName"]
	if !ok {
		return adapter, errors.New("the config's dataSourceName is not set")
	}

	// 创建Query
	queryAdapter, ok := qa.container[driverName]
	if !ok {
		return adapter, errors.New("create '" + driverName + "' query failed")
	}

	// 获取db
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return adapter, err
	}

	// 立即检查数据库是否可用并可访问
	if err = db.Ping(); err != nil {
		return adapter, errors.New("Connect '" + driverName + "' failed: the target machine actively refused it")
	} else {
		queryAdapter.SetDb(db)                // 设置db
		qa.container[identify] = queryAdapter // 存储
	}
	return queryAdapter.Copy(), nil
}