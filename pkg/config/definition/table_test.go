package definition

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var table_definition = []byte(`
table_definition:
  user:
    - column_name: id
    - column_name: age
    - column_name: name
    - column_name: type
      oneof:
        ref: status
        select: [normal, vip, admin] # 如果为空选中全部
    - column_name: mobile
      fromat: mobile
`)

func TestTableParseFile(t *testing.T) {
	e := new(Tabledefinition)
	_ = e
	test_data := []Field{
		{ColumnName: "id"},
		{ColumnName: "age"},
		{ColumnName: "name"},
		{ColumnName: "type", OneOf: FieldOneOf{Ref: "status", Select: []string{"normal", "vip", "admin"}}},
		{ColumnName: "mobile", Fromat: "mobile"},
	}
	_ = test_data
	tf, err := os.CreateTemp("", "test.yml")
	assert.NoError(t, err)
	defer os.Remove(tf.Name())
	os.WriteFile(tf.Name(), table_definition, 0666)
	em := e.ParseFile(tf.Name())
	assert.NotNil(t, em.Definition["user"])
	_ = em
	for k, v := range em.Definition["user"] {
		assert.Equal(t, test_data[k].ColumnName, v.ColumnName)
		assert.Equal(t, test_data[k].Fromat, v.Fromat)
		assert.Equal(t, test_data[k].OneOf, v.OneOf)
	}
}

func TestTableParse(t *testing.T) {
	e := new(Tabledefinition)
	_ = e
	test_data := []Field{
		{ColumnName: "id"},
		{ColumnName: "age"},
		{ColumnName: "name"},
		{ColumnName: "type", OneOf: FieldOneOf{Ref: "status", Select: []string{"normal", "vip", "admin"}}},
		{ColumnName: "mobile", Fromat: "mobile"},
	}
	_ = test_data
	em := e.Parse(table_definition)
	assert.NotNil(t, em.Definition["user"])
	_ = em
	for k, v := range em.Definition["user"] {
		assert.Equal(t, test_data[k].ColumnName, v.ColumnName)
		assert.Equal(t, test_data[k].Fromat, v.Fromat)
		assert.Equal(t, test_data[k].OneOf, v.OneOf)
	}
}
