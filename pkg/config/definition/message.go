package definition

import "cmx/pkg/util"

// message_definition:
//
//	id:
//	  - column_name: id
//	    array: true
//	    tag:
//	      gorm: primary_key,auto_increment
//	    inhibit: required

type MessageDefinition struct {
	Definition map[string][]MessageField `json:"message_definition" yaml:"message_definition"`
}

func (e *MessageDefinition) ParseFile(path string) *MessageDefinition {
	util.NoError(parseYamlFile(path, e))
	return e
}

func (e *MessageDefinition) Parse(in []byte) *MessageDefinition {
	util.NoError(parseYaml(in, e))
	return e
}

type MessageField struct {
	ColumnName    string           `josn:"column_name" yaml:"column_name"`
	Type          string           `json:"type" yaml:"type,omitempty"`
	Array         bool             `json:"array" yaml:"array,omitempty"`
	Optional      bool             `json:"optional" yaml:"optional,omitempty"`
	OneOf         FieldOneOf       `json:"oneof" yaml:"oneof,omitempty"`
	Ref           MessageReference `json:"ref" yaml:"ref,omitempty"`
	Comment       string           `json:"comment" yaml:"comment,omitempty"`
	DetailComment string           `json:"detail_comment" yaml:"detail_comment,omitempty"`
	Validator     string           `json:"validator" yaml:"validator,omitempty"`
	Inhibit       string           `json:"inhibit" yaml:"inhibit,omitempty"`
	Serializer    string           `json:"serializer" yaml:"serializer,omitempty"`
}

// ref type
type MessageReferenceType string

const (
	RefTypeEmpty        MessageReferenceType = ""              // empty
	RefTypeMessage      MessageReferenceType = "message"       // ref message
	RefTypeField        MessageReferenceType = "field"         // ref table field
	RefTypeMessageField MessageReferenceType = "message_field" // ref message field
)

type MessageReference struct {
	Type   MessageReferenceType `json:"type" yaml:"type"`
	Ref    string               `json:"ref" yaml:"ref"`
	Select []string             `json:"select" yaml:"select"`
}
