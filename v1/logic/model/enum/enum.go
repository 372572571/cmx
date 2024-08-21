package enum_model

import (
	"cmx/v1/logic/util"
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

type EnumsGroup struct {
	Definition map[string][]Enum `json:"enums_definition" yaml:"enums_definition"`
}

func (e *EnumsGroup) ParseFile(path string) (*EnumsGroup, error) {
	err := util.ParseYamlFile(path, e)
	return e, err
}

func (e *EnumsGroup) Parse(in []byte) (*EnumsGroup, error) {
	err := util.ParseYaml(in, e)
	return e, err
}

type Enum struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
	Desc  string `json:"desc" yaml:"desc"`
	Zh    string `json:"zh" yaml:"zh"`
	En    string `json:"en" yaml:"en"`
}
