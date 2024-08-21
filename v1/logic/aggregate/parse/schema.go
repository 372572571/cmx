package parse

/*
 # 数据库中表示的的类型主常用的类型
   char               = go string
   varchar            = go string
   text               = go string
   set                = go string
   date               = go time.Time
   datetime           = go time.Time
   timestamp          = go time.Time
   time               = go time.Time
   tinyint            = go int8
   smallint           = go int16
   int                = go int32
*/
/*
	# 在proto中的类型
	double   = go float64
	float    = go float32
	int32    = go int32
	int64    = go int64
	uint32   = go uint32
	uint64   = go uint64
	sint32   = go int32    // 这些比普通的int32s更有效地编码负数。
	sint64   = go int64
	fixed32  = go uint32   // 总是4个字节
	fixed64  = go uint64   // 总是8个字节
	sfixed32 = go int32    // 总是4个字节
	sfixed64 = go int64    // 总是8个字节
	bool     = go bool
	string   = go string
	bytes    = go []byte
*/

type SqlFieldSchema struct {
	Type  SqlFieldType     `json:"type" yaml:"type"`
	Class FieldSchemaClass `json:"class" yaml:"class"`

	IsUnsigned      bool  `json:"is_unsigned" yaml:"is_unsigned"`
	IsNotNull       *bool `json:"is_not_null" yaml:"is_not_null"`
	IsAutoincrement bool  `json:"is_autoincrement" yaml:"is_autoincrement"`
	IsZerofill      bool  `json:"is_zerofill" yaml:"is_zerofill"`

	Comment      string `json:"comment" yaml:"comment"`
	DefaultValue string `json:"default_value" yaml:"default_value"`

	Scale   int    `json:"scale" yaml:"scale"`
	Charset string `json:"charset" yaml:"charset"`
	Collate string `json:"collate" yaml:"collate"`

	EnumValues []string `json:"enum_values" yaml:"enum_values"`

	// KeyTyep KeyType `json:"key_type" yaml:"key_type"`

	Format string `json:"format" yaml:"format"`
	Lenght int    `json:"lenght" yaml:"lenght"`
}

func NewSqlFieldSchema() *SqlFieldSchema {
	return &SqlFieldSchema{
		Class: FieldSchemaClassSql,
	}
}

type GoFieldSchema struct {
	Type  GoFieldType      `json:"type" yaml:"type"`
	Class FieldSchemaClass `json:"class" yaml:"class"`

	KeyType KeyType `json:"key_type" yaml:"key_type"`

	IsNullable bool `json:"is_nullable" yaml:"is_nullable"`

	Tag     string   `json:"tag" yaml:"tag"`
	Comment string   `json:"comment" yaml:"comment"`
	Length  int      `json:"length" yaml:"length"`
	Format  string   `json:"format" yaml:"format"`
	OneOf   []string `json:"one_of" yaml:"one_of"`
}

func NewGoFieldSchema() *GoFieldSchema {
	return &GoFieldSchema{
		Class: FieldSchemaClassGo,
	}
}

type ProtoFieldSchema struct {
	Type  ProtoFieldType   `json:"type" yaml:"type"`
	Class FieldSchemaClass `json:"class" yaml:"class"`

	IsNullable bool `json:"is_nullable" yaml:"is_nullable"`

	Tag     string   `json:"tag" yaml:"tag"`
	Comment string   `json:"comment" yaml:"comment"`
	Length  int      `json:"length" yaml:"length"`
	Format  string   `json:"format" yaml:"format"`
	OneOf   []string `json:"one_of" yaml:"one_of"`
}

func NewProtoFieldSchema() *ProtoFieldSchema {
	return &ProtoFieldSchema{
		Class: FieldSchemaClassProto,
	}
}
