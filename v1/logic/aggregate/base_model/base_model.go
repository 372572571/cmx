package basemodel

import (
	enum_model "cmx/v1/logic/model/enum"
	message_model "cmx/v1/logic/model/message"
	statement_model "cmx/v1/logic/model/statement"
)


type BaseModel struct {
	StatementDefinition map[string]statement_model.CreateStatement `json:"statement_definition" yaml:"statement_definition"`
	TableDefinition     map[string][]message_model.Field         `json:"table_definition" yaml:"table_definition"`
	MessageDefinition   map[string][]message_model.MessageField  `json:"message_definition" yaml:"message_definition"`
	EnumsDefinition     map[string][]enum_model.Enum         `json:"enums_definition" yaml:"enums_definition"`
}
