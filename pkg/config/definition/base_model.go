package definition

type BaseModel struct {
	StatementDefinition map[string]CreateStatement `json:"statement_definition" yaml:"statement_definition"`
	TableDefinition     map[string][]Field         `json:"table_definition" yaml:"table_definition"`
	MessageDefinition   map[string][]MessageField  `json:"message_definition" yaml:"message_definition"`
	EnumsDefinition     map[string][]Enumx         `json:"enums_definition" yaml:"enums_definition"`
}
