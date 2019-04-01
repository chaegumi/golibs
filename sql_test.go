package golibs

import (
	"log"
	"testing"
)

func TestInsertSQL(t *testing.T) {
	m := make(map[string]interface{})

	m["customer_name"] = "Customer1"
	m["parent_id"] = 0

	sql, fieldValues := InsertSQL("customers", m)

	log.Print(sql)
	log.Print(fieldValues)
}

func ExampleInsertSQL() {
	m := make(map[string]interface{})

	m["customer_name"] = "Customer1"
	m["parent_id"] = 0

	sql, fieldValues := InsertSQL("customers", m)

	log.Print(sql)
	log.Print(fieldValues)
	// insert into customers(customer_name, parent_id) values(?, ?)
	// beego orm run sql
	// o.Raw(sql, fieldValues...).Exec()
}

func TestUpdateSQL(t *testing.T) {
	m := make(map[string]interface{})

	m["customer_name"] = "Customer2"
	m["parant_id"] = 0

	sql, fieldValues := UpdateSQL("customers", "id", m)

	log.Print(sql)
	log.Print(fieldValues)
}

func ExampleUpdateSQL() {
	m := make(map[string]interface{})

	m["customer_name"] = "Customer2"
	m["parant_id"] = 0

	sql, fieldValues := UpdateSQL("customers", "id", m)

	log.Print(sql)
	log.Print(fieldValues)

	id := 1

	fieldValues = append(fieldValues, id)
	// update customers set customer_name=?, parent_id=? where id=?
	// beego orm run sql
	// o.Raw(sql, fieldValues...).Exec()
}
