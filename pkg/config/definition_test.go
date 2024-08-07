package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveExt(t *testing.T) {
	df := NewDefinition("")
	assert.Equal(t, df.removeExt("b.c.yaml"), "b.c")
	assert.Equal(t, df.removeExt("b.文件.yaml"), "b.文件")
	assert.Equal(t, df.removeExt("b/c.yaml"), "b/c")
}

func TestCreateRouterFromPath(t *testing.T) {
	df := NewDefinition("")
	testDate := []struct {
		path   string
		expect string
	}{
		{
			path:   "a/b/c/d.yaml",
			expect: "/a/b/c/d",
		},
		{
			path:   "a/b/c/de.yaml",
			expect: "/a/b/c/de",
		},
		{
			path:   "./a/b/c/dxe.yaml",
			expect: "/a/b/c/dxe",
		},
	}

	for _, v := range testDate {
		assert.Equal(t, v.expect, df.createRouterFromPath(v.path))
	}
}

func TestDefinition_GetStatementField(t *testing.T) {
	InitDefaultConfig(func(cfg *Config) {
		cfg.ProjectPath = "../../example/sql/yaml"
	})
	referenceString := "partner.partner"
	ri := NewReferenceInformation(referenceString)
	assert.Equal(t, "partner", ri.Route)
	assert.Equal(t, "partner", ri.Field)
	result, found := GetDefaultConfig().GetDefinition().GetStatementField(referenceString)
	assert.True(t, found)
	assert.NotEqual(t, 0, len(result.Statement))
}

func TestDefinition_GetTableField(t *testing.T) {
	InitDefaultConfig(func(cfg *Config) {
		cfg.ProjectPath = "../../example/sql/yaml"
	})
	referenceString := "partner.id"
	ri := NewReferenceInformation(referenceString)
	assert.Equal(t, "partner", ri.Route)
	assert.Equal(t, "id", ri.Field)
	_, found := GetDefaultConfig().GetDefinition().GetTableField(referenceString)
	assert.True(t, found)
}

func TestDefinition_GetEnumField(t *testing.T) {
	InitDefaultConfig(func(cfg *Config) {
		cfg.ProjectPath = "../../example/sql/yaml"
	})
	referenceString := "default.group.status"
	ri := NewReferenceInformation(referenceString)
	assert.Equal(t, "default/group", ri.Route)
	assert.Equal(t, "status", ri.Field)
	data, found := GetDefaultConfig().GetDefinition().GetEnumField(referenceString)
	assert.True(t, found)
	assert.NotEqual(t, 0, len(data))
}

