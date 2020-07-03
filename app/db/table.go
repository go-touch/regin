package db

import (
	"errors"
	"github.com/go-touch/regin/app/db/model"
)

// Table容器
type TableContainer map[string]*model.Table

var tableContainer *TableContainer

func init() {
	tableContainer = &TableContainer{}
}

// 设置table
func (tc *TableContainer) Set(key string, table *model.Table) error {
	(*tc)[key] = table
	return nil
}

// 获取table
func (tc *TableContainer) Get(key string) (*model.Table, error) {
	if table, ok := (*tc)[key]; ok {
		return table, nil
	}
	return nil, errors.New("this model '" + key + "' is not registered")
}

// Get a table of Model
func GetTable(userModel interface{}) *model.Table {
	var table *model.Table
	if key, ok := userModel.(string); ok {
		if value, err := tableContainer.Get(key); err == nil {
			table = value
		}
	} else if newTable, err := new(model.Table).Init(userModel); err != nil {
		panic(err.Error())
	} else if name, err := newTable.GetName(); err != nil {
		panic(err.Error())
	} else if storage, err := tableContainer.Get(name); err == nil {
		table = storage
	} else {
		table = newTable.Factory()
		if err := tableContainer.Set(name, table); err != nil {
			panic(err.Error())
		}
	}
	return table
}
