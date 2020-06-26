package db

import (
	"github.com/go-touch/regin/app/db/query"
	"github.com/go-touch/regin/base"
)

/**************************************** 数据库模型接口 ****************************************/
type ModelInterface interface {
	Init()
	Self(ModelInterface) ModelInterface                      // 设置自身指针
	Identify() string                                        // 数据库标识(此方法可重构,用于切换数据库,默认master)
	TableName() string                                       // 数据库表名(此方法可重构,用于切换数据表)
	GetQuery() query.BaseQuery                               // 获取查询构造器(此方法禁止须重构)
	FetchRow()                                               // 查询一条记录(此方法禁止须重构)
	Where()                                                  // 查询条件(此方法禁止须重构)
	Call(methodName string, args ...interface{}) interface{} // 调用自定义方法(此方法禁止重构)
}

type Model struct {
	ModelInterface
	self         ModelInterface
	queryAdapter query.BaseQuery
	Fields       struct{}
	Method       map[string]base.UserFunc
}

// 设置Redis自身
func (m *Model) Init() {
	m.Method = make(map[string]base.UserFunc)
}

// 设置model自身指针
func (b *Model) Self(model ModelInterface) ModelInterface {
	b.self = model
	return model
}

// 获取数据库标识
func (b *Model) Identify() string {
	return "master"
}

// 获取表名
func (b *Model) TableName() string {
	return ""
}

// 获取查询构造器
func (b *Model) GetQuery() query.BaseQuery {
	if b.self == nil {
		panic("model error")
	}

	// 获取
	queryAdapter, err := Query.GetAdapter(b.self.Identify())
	if err != nil {
		panic(err)
	}
	return queryAdapter
}

func (b *Model) Where() {
	return
}

// 获取一条记录
func (b *Model) FetchRow() {
	return
}

// 调用自定义方法
func (m *Model) Call(methodName string, args ...interface{}) interface{} {
	if method, ok := m.Method[methodName]; ok {
		return method(args...)
	}
	return ""
}
