package definition

import (
	"cmx/pkg/util"

	"github.com/samber/lo"
)

/*

api_definition:
  user: # 方法组名称
    - name: create
      http:
        method: POST
        path: /v1/user/create
        summary: "创建用户"
      request: default.group.id
      response: default.group.id
    - name: update
      http:
        method: POST
        path: /v1/user/update
        summary: "更新用户"
      request: default.group.id
      response: default.group.id

api_definition:
  user: # 方法组名称
    - name: remove
      http:
        method: POST
        path: /v1/user/remove
        summary: "user remove"
      request: api.user.remove.remove_request
      response: default.group.id

message_definition:
  remove_request:
    - column_name: id
      ref:
        type: field
        ref: partner.id

*/

// ["group_name"][]Api
var ApiCache = map[string][]Api{}

type Apidefinition struct {
	Definition map[string][]Api `json:"api_definition" yaml:"api_definition"`
}

type SignType string

const (
	SignTypeDefault SignType = ""
	// SignTypeSkip             SignType = "skip"
	SignTypeHmacSha256Secret SignType = "hmac-sha256-secret"
)

type Api struct {
	Name string `json:"name" yaml:"name"`
	Http struct {
		IsOpenApi bool   `json:"is_open_api" yaml:"is_open_api"`
		Method    string `json:"method" yaml:"method"`
		Path      string `json:"path" yaml:"path"`
		Body      string `json:"body" yaml:"body"`
		Summary   string `json:"summary" yaml:"summary"`
	} `json:"http" yaml:"http"`
	Request     string   `json:"request" yaml:"request"`
	Response    string   `json:"response" yaml:"response"`
	SignType    SignType `json:"sign_type" yaml:"sign_type"`
	Description string   `json:"description" yaml:"description"`
}

func (e *Apidefinition) addCache() {
	for groupName, v := range e.Definition {
		if _, ok := ApiCache[groupName]; !ok {
			ApiCache[groupName] = []Api{}
		}
		ary := append(ApiCache[groupName], v...)
		ary = lo.Uniq(ary)
		ApiCache[groupName] = ary
	}
}

func (e *Apidefinition) GetGroup(group string) []Api {
	if _, ok := ApiCache[group]; !ok {
		return []Api{}
	}
	return ApiCache[group]
}

func (e *Apidefinition) ParseFile(path string) *Apidefinition {
	util.NoError(parseYamlFile(path, e))
	e.addCache()
	return e
}

func (e *Apidefinition) Parse(in []byte) *Apidefinition {
	util.NoError(parseYaml(in, e))
	e.addCache()
	return e
}
