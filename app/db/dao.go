package db

import (
	"database/sql"
	"errors"
	"github.com/go-touch/regin/app/db/model"
	"github.com/go-touch/regin/app/db/query"
	"regexp"
	"strings"
)

// 用户自定义函数
type UserFunc func(*Dao)

// 数据对象Dao
type Dao struct {
	table *model.Table
	query query.BaseQuery
	isSQL bool
	isTx  bool
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

// 执行SQL - 查询一条数据
func (d *Dao) QueryRow(sql string, args ...interface{}) *AnyValue {
	anyValue := d.Query(sql, args...)
	if err := anyValue.ToError(); err != nil {
		return Eval(err)
	}
	if stringStringSlice := anyValue.ToStringMapSlice(); len(stringStringSlice) == 0 {
		return Eval(map[string]string{})
	} else {
		return Eval(stringStringSlice[0])
	}
}

// 执行SQL - 增删改查
func (d *Dao) Query(sql string, args ...interface{}) *AnyValue {
	defer d.reset()
	sqlArray := strings.Split(sql, " ")
	switch strings.ToUpper(sqlArray[0]) {
	case "SELECT":
		sqlRows, err := d.query.QueryAll(sql, args...)
		if err != nil {
			return Eval(err)
		}
		return d.parserRows(sqlRows)
	case "INSERT", "UPDATE", "DELETE":
		sqlResult, err := d.query.Exec(sql, args...)
		if err != nil {
			return Eval(err)
		}
		return Eval(sqlResult)
	}
	return Eval(errors.New("this sql is illegal, Please check it"))
}

// 获取查询构造器
func (d *Dao) GetQuery() query.BaseQuery {
	return d.query
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
	expr := make([]string, 0)
	field = strings.TrimSpace(field)
	field = regexp.MustCompile(`\s+`+"").ReplaceAllString(field, " ")

	// Where expr.
	expr = append(expr, strings.Split(field, " ")...)
	if regexp.MustCompile(`(<|<=|=|>|>=|!=|like|not like)`+"").FindString(field) != "" {
		expr = append(expr, "?")
	} else if regexp.MustCompile("in").FindString(field) != "" {
		newValue := make([]interface{}, 0)
		if val, ok := value.(string); ok {
			stringSlice := strings.Split(val, ",")
			for _, val := range stringSlice {
				newValue = append(newValue, val)
			}
		} else if val, ok := value.([]interface{}); ok {
			for _, v := range val {
				newValue = append(newValue, v)
			}
		} else if val, ok := value.([]int); ok {
			for _, v := range val {
				newValue = append(newValue, v)
			}
		} else if val, ok := value.([]string); ok {
			for _, v := range val {
				newValue = append(newValue, v)
			}
		}
		inValue := strings.Repeat("?,", len(newValue))
		inValue = strings.Trim(inValue, ",")
		inValue = "(" + inValue + ")"
		expr = append(expr, inValue)
		value = newValue
	} else {
		expr = append(expr, "=", "?")
	}
	d.query.Where(strings.Join(expr, " "), value, linkSymbol...)
	return d
}

// Where组条件
func (d *Dao) WhereMap(fieldMap map[string]interface{}, linkSymbol ...string) *Dao {
	for field, value := range fieldMap {
		d.Where(field, value, linkSymbol...)
	}
	return d
}

// Bind SQL VALUES for INSERT or UPDATE SQL.
func (d *Dao) Values(valueMap map[string]interface{}) *Dao {
	if d.query.GetSqlType() == "INSERT" {
		d.query.Values(valueMap)
	} else {
		d.query.Set(valueMap)
	}
	return d
}

// Batch bind VALUES for INSERT SQL.
func (d *Dao) BatchValues(anyMapSlice []map[string]interface{}) *Dao {
	if len(anyMapSlice) > 20 {
		anyMapSlice = anyMapSlice[0:20]
	}
	for _, anyMap := range anyMapSlice {
		d.query.Values(anyMap)
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

// 是否打印SQL
func (d *Dao) Sql() *Dao {
	d.isSQL = true
	return d
}

// 查询一条记录
func (d *Dao) FetchRow(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	_ = d.query.Limit(1)
	_ = d.fetch(userFunc...)
	if d.isSQL {
		return Eval(d.query.GetSql())
	}
	// Get row.
	sqlRow := d.query.FetchRow()
	return d.parserRow(sqlRow)
}

// 查询多条记录
func (d *Dao) FetchAll(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	_ = d.fetch(userFunc...)
	if d.isSQL {
		return Eval(d.query.GetSql())
	}
	// Get rows.
	sqlRows, err := d.query.FetchAll()
	if err != nil {
		return Eval(err)
	}
	return d.parserRows(sqlRows)
}

// 插入记录
func (d *Dao) Insert(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	return d.modify("INSERT", userFunc...)
}

// 更新记录
func (d *Dao) Update(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	return d.modify("UPDATE", userFunc...)
}

// 删除记录
func (d *Dao) Delete(userFunc ...UserFunc) *AnyValue {
	defer d.reset()
	return d.modify("DELETE", userFunc...)
}

// 事务接管
func (d *Dao) Tx(dao *Dao) *Dao {
	if tx := dao.GetQuery().GetTx(); tx != nil {
		d.isTx = true
		d.GetQuery().SetTx(tx)
	}
	return d
}

// 开启事务
func (d *Dao) Begin() {
	d.query.Begin()
}

// 提交事务
func (d *Dao) Commit() {
	d.query.Commit()
}

// 回滚事务
func (d *Dao) Rollback() {
	d.query.Rollback()
}

// Common SELECT part.
func (d *Dao) fetch(userFunc ...UserFunc) *AnyValue {
	_ = d.query.SetSqlType("SELECT")
	// 执行过程
	if userFunc != nil {
		userFunc[0](d)
	}
	// 字段处理
	if d.query.GetField().GetExpr() == "" {
		d.query.Field(d.table.GetTableFields())
	}
	_ = d.query.SetSql() // SQL处理
	return nil
}

// 执行过程
func (d *Dao) modify(sType string, userFunc ...UserFunc) *AnyValue {
	_ = d.query.SetSqlType(sType)
	// 执行过程
	if userFunc != nil {
		userFunc[0](d)
	}
	// SQL处理
	_ = d.query.SetSql()
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

// 解析单行记录
func (d *Dao) parserRow(sqlRow *sql.Row) *AnyValue {
	column := d.query.GetField().GetNameArray()
	// 接收参数
	args := make([]interface{}, len(column))
	for k := range args {
		args[k] = &args[k]
	}
	// 接收查询结果
	row := make(map[string]interface{})
	err := sqlRow.Scan(args...)
	if err != nil {
		if regexp.MustCompile("no rows in result set"+"").FindString(err.Error()) != "" {
			return Eval(row)
		}
		return Eval(err)
	}
	// 结果处理
	for i := 0; i < len(column); i++ {
		row[column[i]] = args[i]
	}
	return Eval(row)
}

// 解析多行记录
func (d *Dao) parserRows(sqlRows *sql.Rows) *AnyValue {
	defer func() { _ = sqlRows.Close() }()
	// 获取字段
	columns, err2 := sqlRows.Columns()
	if err2 != nil {
		return Eval(err2)
	}
	// 迭代后者的 Next() 方法，然后使用 Scan() 方法给对应类型变量赋值，以便取出结果，最后再把结果集关闭（释放连接）
	list := make([]map[string]interface{}, 0)
	length := len(columns) // 字段数组长度
	for sqlRows.Next() {
		// 接收参数
		args := make([]interface{}, length)
		for k := range args {
			args[k] = &args[k]
		}
		// 数据行接收
		if err := sqlRows.Scan(args...); err != nil {
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

// 重置结构体
func (d *Dao) reset() {
	d.isSQL = false
	_ = d.query.Reset()
	if d.isTx == true {
		d.isTx = false
		d.query.UnsetTx()
	}
}
