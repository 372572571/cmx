package definition

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var message_definition = []byte(`
message_definition:
  id:
    - column_name: id
      array: true
      inhibit: required
`)

func TestMessageParseFile(t *testing.T) {
	e := new(MessageDefinition)
	_ = e
	test_data := []MessageField{
		MessageField{ColumnName: "id", Array: true},
	}

	_ = test_data
	tf, err := os.CreateTemp("", "test.yml")
	assert.NoError(t, err)
	defer os.Remove(tf.Name())
	os.WriteFile(tf.Name(), message_definition, 0666)
	em := e.ParseFile(tf.Name())
	assert.NotNil(t, em.Definition["id"])
	_ = em
	for k, v := range em.Definition["id"] {
		assert.Equal(t, test_data[k].ColumnName, v.ColumnName)
		assert.Equal(t, test_data[k].Array, v.Array)
	}
}

func TestMessageParse(t *testing.T) {
	e := new(MessageDefinition)
	_ = e
	test_data := []MessageField{
		MessageField{ColumnName: "id", Array: true},
	}
	em := e.Parse(message_definition)
	assert.NotNil(t, em.Definition["id"])
	_ = em
	for k, v := range em.Definition["id"] {
		assert.Equal(t, test_data[k].ColumnName, v.ColumnName)
		assert.Equal(t, test_data[k].Array, v.Array)
	}
}
