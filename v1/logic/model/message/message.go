package message_model

import "cmx/v1/logic/util"

// ref type
type MessageReferenceType string

const (
	RefTypeEmpty        MessageReferenceType = ""              // empty
	// 引用 message
	RefTypeMessage      MessageReferenceType = "message"       // ref message
	// 引用 message field
	RefTypeMessageField MessageReferenceType = "message_field" // ref message field

	RefTypeField        MessageReferenceType = "field"         // ref table field
)

type InhibitType string

const (
	Inhibit  InhibitType = "inhibit"
	Required InhibitType = "required"
	Optional InhibitType = "optional"
)

type MessageReference struct {
	Ref    string               `json:"ref" yaml:"ref"`
	Select []string             `json:"select" yaml:"select"`
	Type   MessageReferenceType `json:"type" yaml:"type"`
}

type FieldOneOf struct {
	Ref    string   `json:"ref" yaml:"ref"`
	Select []string `json:"select" yaml:"select"`
	IsKey  bool     `json:"is_key" yaml:"is_key"`
}

type MessageField struct {
	ColumnName    string           `json:"column_name" yaml:"column_name"`
	Type          string           `json:"type" yaml:"type,omitempty"`
	Array         bool             `json:"array" yaml:"array,omitempty"`
	Optional      bool             `json:"optional" yaml:"optional,omitempty"`
	OneOf         FieldOneOf       `json:"oneof" yaml:"oneof,omitempty"`
	Ref           MessageReference `json:"ref" yaml:"ref,omitempty"`
	Comment       string           `json:"comment" yaml:"comment,omitempty"`
	Validator     string           `json:"validator" yaml:"validator,omitempty"`
	Inhibit       string           `json:"inhibit" yaml:"inhibit,omitempty"`
	Serializer    string           `json:"serializer" yaml:"serializer,omitempty"`
}


type Message struct {
	Definition map[string][]MessageField `json:"message_definition" yaml:"message_definition"`
}

func (e *Message) ParseFile(path string) (*Message,error) {
	err := util.ParseYamlFile(path, e)
	return e,err
}

func (e *Message) Parse(in []byte) (*Message,error) {
	err :=  util.ParseYaml(in,e)
	return e, err
}


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
type Table struct {
	Definition map[string][]Field `json:"table_definition" yaml:"table_definition"`
}

func (e *Table) ParseFile(path string) (*Table,error) {
	err := util.ParseYamlFile(path, e)
	return e,err
}

func (e *Table) Parse(in []byte) (*Table,error) {
	err := util.ParseYaml(in, e)
	return e,err
}

type Field struct {
	ColumnName    string      `json:"column_name" yaml:"column_name"`
	Ref           string      `json:"ref" yaml:"ref,omitempty"`
	Type          string      `json:"type" yaml:"type"`
	Format        string      `json:"format,omitempty" yaml:"format,omitempty"`
	OneOf         FieldOneOf  `json:"oneof" yaml:"oneof,omitempty"`
	Inhibit       InhibitType `json:"inhibit" yaml:"inhibit,omitempty"`
	Comment       string      `json:"comment" yaml:"comment"`
	DetailComment string      `json:"detail_comment" yaml:"detail_comment,omitempty"`
	Validator     string      `json:"validator" yaml:"validator,omitempty"`
}