package api_model

import (
	"cmx/v1/logic/util"

	"github.com/samber/lo"
)

var ApiCache = map[string][]Api{}

type Api struct {
	Name    string `json:"name" yaml:"name"`
	SubPath string `json:"sub_path" yaml:"sub_path"`
	Http    struct {
		IsPublic  bool   `json:"is_public" yaml:"is_public"`
		IsOpenApi bool   `json:"is_open_api" yaml:"is_open_api"`
		Method    string `json:"method" yaml:"method"`
		Path      string `json:"path" yaml:"path"`
		Body      string `json:"body" yaml:"body"`
		Summary   string `json:"summary" yaml:"summary"`
	} `json:"http" yaml:"http"`
	Request     string    `json:"request" yaml:"request"`
	Response    string    `json:"response" yaml:"response"`
	Description string    `json:"description" yaml:"description"`
	Tags        *[]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type ApiDefinition struct {
	Definition map[string][]Api `json:"api_definition" yaml:"api_definition"`
}

func (e *ApiDefinition) addCache() {
	for groupName, v := range e.Definition {
		if _, ok := ApiCache[groupName]; !ok {
			ApiCache[groupName] = []Api{}
		}
		ary := append(ApiCache[groupName], v...)
		ary = lo.Uniq(ary)
		ApiCache[groupName] = ary
	}
}

func (e *ApiDefinition) GetGroup(group string) []Api {
	if _, ok := ApiCache[group]; !ok {
		return []Api{}
	}
	return ApiCache[group]
}

func (e *ApiDefinition) ParseFile(path string) (*ApiDefinition, error) {
	err := util.ParseYamlFile(path, e)
	if err != nil {
		return nil, err
	}
	e.addCache()
	return e, nil
}

func (e *ApiDefinition) Parse(in []byte) (*ApiDefinition, error) {
	err := util.ParseYaml(in, e)
	if err != nil {
		return nil, err
	}
	e.addCache()
	return e, nil
}
