package config

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestConfig(t *testing.T) {
	str, _ := yaml.Marshal(defaultConfig)
	t.Log(string(str))
}
