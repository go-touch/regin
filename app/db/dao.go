package db

import (
	"fmt"
	"github.com/go-touch/regin/app/db/model"
	"github.com/go-touch/regin/app/db/query"
)

// 用户自定义函数
type UserFunc func(*Dao)

// 数据对象Dao
type Dao struct {
	table *model.Table
	query query.BaseQuery
	isSQL bool
}

// 注册model
func RegisterModel(userModel interface{}, alias ...string) {
	key := ""
	table, err := new(model.Table).Init(userModel)
	if err != nil {
		panic(err.Error())
	}
	if alias != nil && alias[0] != "" {
		key = alias[0]
	} else if name, err := table.GetName(); err != nil {
		panic(err.Error())
	} else {
		key = name
	}
	if err := tableContainer.Set(key, table.Factory()); err != nil {
		panic(err.Error())
	}
}

// 获取Dao
func Model(userModel interface{}) *Dao {
	tableInstance := GetTable(userModel)
	queryInstance := GetQueryBuilder(tableInstance.Identify)
	queryInstance.Table(tableInstance.TableName)
	return &Dao{
		table: tableInstance,
		query: queryInstance,
		isSQL: false,
	}
}

func Test(){
	fmt.Println(tableContainer)
}

// 执行SQL
func (d *Dao) Query(sql string, args ...interface{}) (result interface{}, err error) {
	return d.query.Query(sql, args...)
}

// 设置表名
func (d *Dao) Table(tableName string) *Dao {
	d.query.Table(tableName)
	return d
}

// 设置字段
func (d *Dao) Field(field interface{}) *Dao {
	d.query.Field(field)
	return d
}

// Where条件
func (d *Dao) Where(field string, value interface{}, linkSymbol ...string) *Dao {
	d.query.Where(field+" = ?", value, linkSymbol...)
	return d
}

// Where组条件
func (d *Dao) WhereMap(fieldMap map[string]interface{}, linkSymbol ...string) *Dao {
	for field, value := range fieldMap {
		d.query.Where(field+" = ?", value, linkSymbol...)
	}
	return d
}

// 设置字段
func (d *Dao) Order(expr ...string) *Dao {
	for _, v := range expr {
		d.query.Order(v)
	}
	return d
}

// 设置字段
func (d *Dao) OrderSlice(expr []string) *Dao {
	for _, v := range expr {
		d.query.Order(v)
	}
	return d
}

// 分页查询
func (d *Dao) Limit(limit ...int) *Dao {
	d.query.Limit(limit...)
	return d
}

// 插入时绑定数据
func (d *Dao) Values(valueMap map[string]interface{}) *Dao {
	d.query.Values(valueMap)
	return d
}

// 更新时绑定数据
func (d *Dao) Set(valueMap map[string]interface{}) *Dao {
	d.query.Set(valueMap)
	return d
}

// 是否打印SQL
func (d *Dao) Sql() *Dao {
	d.isSQL = true
	return d
}

// 查询一条记录
func (d *Dao) FetchRow(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	_ = d.query.Limit(1)
	_ = d.query.SetSQLType("SELECT")
	_ = d.process(userFunc...)
	if d.isSQL {
		return Eval(d.query.GetSql())
	}
	return d.parserRow()
}

// 查询多条记录
func (d *Dao) FetchAll(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	_ = d.query.SetSQLType("SELECT")
	_ = d.process(userFunc...)
	if d.isSQL {
		return Eval(d.query.GetSql())
	}
	return d.parserRows()
}

// 插入记录
func (d *Dao) Insert(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	_ = d.query.SetSQLType("INSERT")
	_ = d.process(userFunc...)
	if d.isSQL {
		return Eval(d.query.GetSql())
	}
	// 执行结果
	result, err := d.query.Modify()
	if err != nil {
		return Eval(err)
	}
	return Eval(result)
}

// 更新记录
func (d *Dao) Update(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	_ = d.query.SetSQLType("UPDATE")
	_ = d.process(userFunc...)
	if d.isSQL {
		return Eval(d.query.GetSql())
	}
	// 执行结果
	result, err := d.query.Modify()
	if err != nil {
		return Eval(err)
	}
	return Eval(result)
}

// 删除记录
func (d *Dao) DELETE(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	_ = d.query.SetSQLType("DELETE")
	_ = d.process(userFunc...)
	if d.isSQL {
		return Eval(d.query.GetSql())
	}
	// 执行结果
	result, err := d.query.Modify()
	if err != nil {
		return Eval(err)
	}
	return Eval(result)
}

// Common SELECT part.
func (d *Dao) process(userFunc ...UserFunc) error {
	if userFunc != nil { // 执行过程
		userFunc[0](d)
	}
	if d.query.GetField().GetExpr() == "" { // 字段处理
		d.query.Field(d.table.GetTableFields())
	}
	_ = d.query.CreateSQL() // SQL处理
	return nil
}

// 解析单行记录
func (d *Dao) parserRow() *AnyValue {
	column := d.query.GetField().GetNameArray()
	sqlRow := d.query.FetchRow()
	// 接收参数
	args := make([]interface{}, len(column))
	for k := range args {
		args[k] = &args[k]
	}
	err := sqlRow.Scan(args...)
	if err != nil {
		return Eval(err)
	}
	// 接收查询结果
	row := make(map[string]interface{})
	for i := 0; i < len(column); i++ {
		row[column[i]] = args[i]
	}
	return Eval(row)
}

// 解析多行记录
func (d *Dao) parserRows() *AnyValue {
	// 返回字段的数组
	rows, err := d.query.FetchAll()
	if err != nil {
		return Eval(err)
	}
	defer func() {
		_ = rows.Close() // 关闭连接
	}()
	// 获取字段
	columns, err2 := rows.Columns()
	if err2 != nil {
		return Eval(err2)
	}
	// 迭代后者的 Next() 方法，然后使用 Scan() 方法给对应类型变量赋值，以便取出结果，最后再把结果集关闭（释放连接）
	list := make([]map[string]interface{}, 0)
	length := len(columns) // 字段数组长度
	for rows.Next() {
		// 接收参数
		args := make([]interface{}, length)
		for k := range args {
			args[k] = &args[k]
		}
		// 数据行接收
		if err := rows.Scan(args...); err != nil {
			Eval(err)
		}
		// 数据赋值
		row := make(map[string]interface{})
		for i := 0; i < length; i++ {
			row[columns[i]] = args[i]
		}
		list = append(list, row)
	}
	return Eval(list)
}

// 调用回调函数
func (d *Dao) callUserFunc(userFunc UserFunc) {
	userFunc(d)
}

// 重置结构体
func (d *Dao) reset() {
	d.isSQL = false
	_ = d.query.Reset()
}
