package query

type MysqlQuery struct {
	Combine
	sqlExpr  string
	sqlParam map[string]string
}

// 获取SQL语句
func (mq *MysqlQuery) Copy() BaseQuery {
	return &MysqlQuery{
		Combine:  mq.Combine,
		sqlExpr:  "",
		sqlParam: make(map[string]string),
	}
}

// 获取SQL语句
func (mq *MysqlQuery) Sql() string {
	return ""
}

// 插入一条记录
func (mq *MysqlQuery) Insert() *MysqlQuery {
	return mq
}

// 更新一条记录
func (mq *MysqlQuery) Update() *MysqlQuery {
	return mq
}

// 删除一条记录
func (mq *MysqlQuery) Delete() *MysqlQuery {
	return mq
}

// 查询条件
func (mq *MysqlQuery) Where() *MysqlQuery {
	return mq
}

// 获取一条记录
func (mq *MysqlQuery) FetchRow() {

}

// 获取多条记录
func (mq *MysqlQuery) FetchAll() {

}
