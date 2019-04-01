package golibs

/*
// 生成数据库语句
*/

import "strings"

// InsertSQL 构建插入语句
//  @param tablename 数据表名称
func InsertSQL(tablename string, m map[string]interface{}) (string, []interface{}) {

	fields := []string{}
	fieldValues := []interface{}{}
	for k, v := range m {
		fields = append(fields, k)
		fieldValues = append(fieldValues, v)
	}

	sql := `insert into ` + tablename + `(` + strings.Join(fields, ",") + `) 
	values(` + strings.TrimRight(strings.Repeat("?,", len(fields)), ",") + `)`
	return sql, fieldValues
}

// UpdateSQL 构建更新语句
func UpdateSQL(tablename string, pkname string, m map[string]interface{}) (string, []interface{}) {
	tsql := ""
	fieldValues := []interface{}{}
	for k, v := range m {
		tsql += k + "=?,"
		fieldValues = append(fieldValues, v)
	}

	sql := `update ` + tablename + ` set ` + strings.TrimRight(tsql, ",") + ` where ` + pkname + `=?`
	return sql, fieldValues
}
