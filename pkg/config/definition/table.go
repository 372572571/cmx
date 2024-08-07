package definition

// import "formworker/util"

// # 字段定义 描述
// table_definition:
//
//	user:
//	  - column_name: id
//	  - column_name: age
//	  - column_name: name
//	  - column_name: type
//	    ref: xxx 标识这个字段引用其他文件里的字段
//	    oneof:
//	      ref: status
//	      select: [normal, vip, admin] # 如果为空选中全部
//	  - column_name: mobile
//	    fromat: mobile
type Tabledefinition struct {
	Definition map[string][]Field `json:"table_definition" yaml:"table_definition"`
}

func (e *Tabledefinition) ParseFile(path string) *Tabledefinition {
	err := parseYamlFile(path, e)
	if err != nil {
		panic(err)
	}
	return e
}

func (e *Tabledefinition) Parse(in []byte) *Tabledefinition {
	err := parseYaml(in, e)
	if err != nil {
		panic(err)
	}
	return e
}

type Field struct {
	ColumnName    string      `json:"column_name" yaml:"column_name"`
	Ref           string      `json:"ref" yaml:"ref,omitempty"`
	Type          string      `json:"type" yaml:"type"`
	Fromat        string      `json:"fromat,omitempty" yaml:"fromat,omitempty"`
	OneOf         FieldOneOf  `json:"oneof" yaml:"oneof,omitempty"`
	Inhibit       InhibitType `json:"inhibit" yaml:"inhibit,omitempty"`
	Comment       string      `json:"comment" yaml:"comment"`
	DetailComment string      `json:"detail_comment" yaml:"detail_comment,omitempty"`
	Validator     string      `json:"validator" yaml:"validator,omitempty"`
	// MaxLength     uint64      `json:"max_length" yaml:"max_length"`
	// MinLength     uint64      `json:"min_length" yaml:"min_length"`
}

type FieldOneOf struct {
	Ref    string   `json:"ref" yaml:"ref"`
	Select []string `json:"select" yaml:"select"`
	IsKey  bool     `json:"is_key" yaml:"is_key"`
}
