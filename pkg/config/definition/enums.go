package definition

import (
	"cmx/pkg/util"
	"fmt"
)

// # 枚举描述
// enums_definition:
//    status:
//       - key: normal
//         value: 1
//         desc: normal user
//       - key: vip
//         value: 2
//         desc: vip user
//       - key: admin
//         value: 3
//         zh: 管理员
// 	   	   en: admin
//         desc: admin user

type Enumsdefinition struct {
	Definition map[string][]Enumx `json:"enums_definition" yaml:"enums_definition"`
}

func (e *Enumsdefinition) ParseFile(path string) *Enumsdefinition {
	if err := parseYamlFile(path, e); err != nil {
		fmt.Println(path)
		panic(err)
	}
	return e
}

func (e *Enumsdefinition) Parse(in []byte) *Enumsdefinition {
	util.NoError(parseYaml(in, e))
	return e
}

type Enumx struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
	Desc  string `json:"desc" yaml:"desc"`
	Zh    string `json:"zh" yaml:"zh"`
	En    string `json:"en" yaml:"en"`
}
