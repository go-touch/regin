package parts

type Join struct {
	members []*SubJoin
}

type SubJoin struct {
	connector string // 连接符, for example: INNER JOIN、LEFT JOIN、RIGHT JOIN
	tableName string // 表名, for example: admin as a
	on        string // 关联条件
}
