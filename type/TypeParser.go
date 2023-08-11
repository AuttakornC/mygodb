package type_manage

import (
	"reflect"
	"strings"
)

// Type Template สำหรับบอกให้ทราบประเภทต่างๆ
type TypeTemplate struct {
	Name   string
	Type   string
	Struct []TypeTemplate
	Option string
}

// Function วิเคราะห์โครงสร้าง
func ParseType(template interface{}) TypeTemplate {

	// Template ต้องเป็น Struct เท่านั้น ไม่งั้น Elem() error
	type_of_template := reflect.TypeOf(template).Elem()

	main_type := parseType(type_of_template)

	return main_type
}

// Function วิเคราห์โครงสร้างในโครงสร้าง
func parseType(template reflect.Type) TypeTemplate {
	var main_type TypeTemplate = TypeTemplate{Name: template.Name(), Type: template.Kind().String(), Struct: []TypeTemplate{}}
	member_in_type := template.NumField()
	for i := 0; i < member_in_type; i++ {
		var new_Type TypeTemplate

		in_Type := template.Field(i)
		if strings.Contains(in_Type.Tag.Get("aut"), "pk") {
			new_Type.Option = "pk"
		}

		// Type Parse
		new_Type.Name = in_Type.Name
		type_type := in_Type.Type
		new_Type.Type = type_type.Kind().String()

		// กรณีเจอ Struct ใน Struct ให้ Recursive ไปเลื่อยๆ
		if new_Type.Type == "struct" {
			new_Type.Struct = append(new_Type.Struct, parseType(type_type).Struct...)
		}
		main_type.Struct = append(main_type.Struct, new_Type)
	}
	return main_type
}
