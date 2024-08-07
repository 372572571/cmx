package definition

import (
	"embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed test.yml
var test_yaml embed.FS

var enums_definition = []byte(`
enums_definition:
  status:
    - key: normal
      value: 1
      desc: normal user
    - key: vip
      value: 2
      desc: vip user
    - key: admin
      value: 3
      zh: 管理员
      desc: admin user
`)

func TestEnumsParseFile(t *testing.T) {
	e := new(Enumsdefinition)
	_ = e
	test_data := []Enumx{
		{Key: "normal", Value: "1", Desc: "normal user"},
		{Key: "vip", Value: "2", Desc: "vip user"},
		{Key: "admin", Value: "3", Zh: "管理员", Desc: "admin user"},
	}
	tf, err := os.CreateTemp("", "test.yml")
	assert.NoError(t, err)
	defer os.Remove(tf.Name())
	os.WriteFile(tf.Name(), enums_definition, 0666)
	em := e.ParseFile(tf.Name())
	assert.NotNil(t, em.Definition["status"])
	_ = em
	for k, v := range em.Definition["status"] {
		assert.Equal(t, test_data[k].Key, v.Key)
		assert.Equal(t, test_data[k].Value, v.Value)
		assert.Equal(t, test_data[k].Desc, v.Desc)
	}
}

func TestEnumsParse(t *testing.T) {
	e := new(Enumsdefinition)
	_ = e
	test_data := []Enumx{
		{Key: "normal", Value: "1", Desc: "normal user"},
		{Key: "vip", Value: "2", Desc: "vip user"},
		{Key: "admin", Value: "3", Zh: "管理员", Desc: "admin user"},
	}
	em := e.Parse(enums_definition)
	assert.NotNil(t, em.Definition["status"])
	_ = em
	for k, v := range em.Definition["status"] {
		assert.Equal(t, test_data[k].Key, v.Key)
		assert.Equal(t, test_data[k].Value, v.Value)
		assert.Equal(t, test_data[k].Desc, v.Desc)
	}
}
