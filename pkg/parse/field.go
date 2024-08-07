package parse

type ModelField struct {
	Idx         int               `json:"idx" yaml:"idx"`
	FieldName   string            `json:"field_name" yaml:"field_name"`
	UpFieldName string            `json:"up_field_name" yaml:"up_field_name"`
	SqlSchema   *SqlFieldSchema   `json:"sql_schema" yaml:"sql_schema"`
	GoSchema    *GoFieldSchema    `json:"go_schema" yaml:"go_schema"`
	ProtoSchema *ProtoFieldSchema `json:"proto_schema" yaml:"proto_schema"`
}

func NewModelField() *ModelField {
	return &ModelField{}
}
