package db

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-touch/regin/app/db/query"
)

// 查询构造器容器
type QueryContainer map[string]query.BaseQuery

var Query *QueryContainer

func init() {
	Query = &QueryContainer{}
}

// 设置queryBuilder
func (qc *QueryContainer) Set(key string, queryBuilder query.BaseQuery) error {
	(*qc)[key] = queryBuilder
	return nil
}

// 获取queryBuilder
func (qc *QueryContainer) Get(key string) (query.BaseQuery, error) {
	if queryBuilder, ok := (*qc)[key]; ok {
		return queryBuilder, nil
	}
	return nil, errors.New("can't find this Query Builder '" + key + "'")
}

// 获取查询构造器
func GetQueryBuilder(identify string) query.BaseQuery {
	if queryBuilder, err := Query.Get(identify); err == nil {
		return queryBuilder.Clone()
	}

	// 读取数据库配置并校验
	configParam := map[string]string{"driverName": "", "dataSourceName": "", "maxIdleConn": "", "maxOpenConn": ""}
	if config, err := Config.GetConfig(identify); err != nil {
		panic(err.Error())
	} else {
		for key, _ := range configParam {
			if v, ok := config[key]; !ok {
				panic("The database config's " + key + " is not set.")
			} else {
				configParam[key] = v
			}
		}
	}

	// 创建db对象、创建查询构造器
	if db, err := sql.Open(configParam["driverName"], configParam["dataSourceName"]); err != nil {
		panic(err.Error())
	} else if err = db.Ping(); err != nil {
		panic("Connect '" + configParam["driverName"] + "' failed: the target machine actively refused it.")
	} else if queryBuilder := query.GetQueryBuilder(configParam["driverName"]); queryBuilder == nil {
		panic("Create '" + configParam["driverName"] + "' Query Builder failed.")
	} else {
		queryBuilder.SetDb(db)
		_ = Query.Set(identify, queryBuilder)
		return queryBuilder.Clone()
	}
}
