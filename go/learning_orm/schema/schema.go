package schema

import (
	"go/ast"
	"learning_orm/dialect"
	"reflect"
)

// Field represents a column of database
type Field struct {
	Name string
	Type string
	Tag  string
}

// Schema represents a table of databse
type Schema struct {
	Model      interface{}
	Name       string
	Fields     []*Field
	FieldNames []string
	fieldMap   map[string]*Field
}

// GetField represents get field func
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// Parse parse sql
func Parse(dast interface{}, d dialect.Dialect) *Schema {
	modelType := reflect.Indirect(reflect.ValueOf(dast)).Type()
	schema := &Schema{
		Model:    dast,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		if p.Anonymous || !ast.IsExported(p.Name) {
			continue
		}

		field := &Field{
			Name: p.Name,
			Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
		}
		if v, ok := p.Tag.Lookup("learning_orm"); ok == true {
			field.Tag = v
		}
		schema.Fields = append(schema.Fields, field)
		schema.FieldNames = append(schema.FieldNames, p.Name)
		schema.fieldMap[p.Name] = field

	}

	return schema
}
