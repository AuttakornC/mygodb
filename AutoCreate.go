package main

import (
	"fmt"
	"reflect"
	"strings"

	type_manage "github.com/AuttakornC/mygodb/type"
)

type Table struct {
	Name       string
	Column     []Column
	Link_Table []Table
}

type Column struct {
	Name string
	Type string
	PK   bool
}

func (m *Mygodb) AutoMigrate(template interface{}) {

	// เช็คการมีอยู่ของตาราง
	if is_exist, err := m.check_table(reflect.TypeOf(template).Elem().Name()); err != nil {
		fmt.Println("+-----------------------------------------------------+")
		fmt.Println("| [ MyGoDB ] : Migrate Error                          |")
		fmt.Println("+-----------------------------------------------------+")
		fmt.Println(err)
		return
	} else {
		if is_exist {
			return
		}
	}

	parsed_type := type_manage.ParseType(template)
	var table_template Table = parseTable(parsed_type)
	fmt.Println(table_template)
}

// Function เช็คตาราง
func (m *Mygodb) check_table(table_name string) (bool, error) {
	rows := m.db.QueryRow("SELECT EXISTS (SELECT FROM pg_tables WHERE schemaname=$1 AND tablename=$2);", "public", strings.ToLower(table_name))
	var is_exist bool
	if err := rows.Scan(&is_exist); err != nil {
		return false, err
	} else {
		return is_exist, nil
	}
}

// Function วิเคราะห์ Type ให้เป็น Table
func parseTable(type_ type_manage.TypeTemplate) Table {
	var table Table
	table.Name = strings.ToLower(type_.Name)
	for _, type_in := range type_.Struct {
		if type_in.Type == "struct" {
			table.Link_Table = append(table.Link_Table, parseTable(type_in))
		} else {
			var table_type string
			var pk_status bool = false
			if strings.Contains(type_in.Option, "pk") {
				pk_status = true
			}
			switch {
			case strings.Contains(strings.ToLower(type_in.Type), "int"):
				table_type = "int"
			case type_in.Type == "string":
				table_type = "text"
			}
			table.Column = append(table.Column, Column{Name: strings.ToLower(type_in.Name), Type: table_type, PK: pk_status})
		}
	}

	return table
}
