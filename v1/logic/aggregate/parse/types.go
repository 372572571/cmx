package parse

type KeyType string

const (
	KeyTypePrimary = KeyType("PRIMARY")
	KeyTypeUnique  = KeyType("UNIQUE")
	KeyTypeIndex   = KeyType("INDEX")
)

type FieldSchemaClass string

const (
	FieldSchemaClassSql   = "SQL"
	FieldSchemaClassGo    = "GO"
	FieldSchemaClassProto = "PROTO"
)

type SqlFieldType string

const (
	SqlFieldTypeInt        = SqlFieldType("int")
	SqlFieldTypeBigInt     = SqlFieldType("bigint")
	SqlFieldTypeTinyInt    = SqlFieldType("tinyint")
	SqlFieldTypeSmallInt   = SqlFieldType("smallint")
	SqlFieldTypeMediumInt  = SqlFieldType("mediumint")
	SqlFieldTypeChar       = SqlFieldType("char")
	SqlFieldTypeVarchar    = SqlFieldType("varchar")
	SqlFieldTypeText       = SqlFieldType("text")
	SqlFieldTypeTinyText   = SqlFieldType("tinytext")
	SqlFieldTypeMediumText = SqlFieldType("mediumtext")
	SqlFieldTypeLongText   = SqlFieldType("longtext")
	SqlFieldTypeDate       = SqlFieldType("date")
	SqlFieldTypeDateTime   = SqlFieldType("datetime")
	SqlFieldTypeTimeStamp  = SqlFieldType("timestamp")
	SqlFieldTypeTime       = SqlFieldType("time")
	SqlFieldTypeEnum       = SqlFieldType("enum")
	SqlFieldTypeDecimal    = SqlFieldType("decimal")
	SqlFieldTypeFloat      = SqlFieldType("float")
	SqlFieldTypeDouble     = SqlFieldType("double")
	SqlFieldTypeBinary     = SqlFieldType("binary")
	SqlFieldTypeJson       = SqlFieldType("json")
)

type GoFieldType string

const (
	GoFieldTypeInt            = GoFieldType("int")
	GoFieldTypeInt8           = GoFieldType("int8")
	GoFieldTypeInt16          = GoFieldType("int16")
	GoFieldTypeInt32          = GoFieldType("int32")
	GoFieldTypeInt64          = GoFieldType("int64")
	GoFieldTypeUint           = GoFieldType("uint")
	GoFieldTypeUint8          = GoFieldType("uint8")
	GoFieldTypeUint16         = GoFieldType("uint16")
	GoFieldTypeUint32         = GoFieldType("uint32")
	GoFieldTypeUint64         = GoFieldType("uint64")
	GoFieldTypeFloat32        = GoFieldType("float32")
	GoFieldTypeFloat64        = GoFieldType("float64")
	GoFieldTypeString         = GoFieldType("string")
	GoFieldTypeBool           = GoFieldType("bool")
	GoFieldTypeTime           = GoFieldType("time.Time")
	GoFieldTypePointerInt     = GoFieldType("*int")
	GoFieldTypePointerInt8    = GoFieldType("*int8")
	GoFieldTypePointerInt16   = GoFieldType("*int16")
	GoFieldTypePointerInt32   = GoFieldType("*int32")
	GoFieldTypePointerInt64   = GoFieldType("*int64")
	GoFieldTypePointerUint    = GoFieldType("*uint")
	GoFieldTypePointerUint8   = GoFieldType("*uint8")
	GoFieldTypePointerUint16  = GoFieldType("*uint16")
	GoFieldTypePointerUint32  = GoFieldType("*uint32")
	GoFieldTypePointerUint64  = GoFieldType("*uint64")
	GoFieldTypePointerFloat32 = GoFieldType("*float32")
	GoFieldTypePointerFloat64 = GoFieldType("*float64")
	GoFieldTypePointerString  = GoFieldType("*string")
	GoFieldTypePointerBool    = GoFieldType("*bool")
	GoFieldTypePointerTime    = GoFieldType("*time.Time")
	GoFieldTypeSoftDelete     = GoFieldType("soft_delete.DeletedAt")
	GoFieldTypeJson           = GoFieldType("json.RawMessage")
)

func (g GoFieldType) IsPointer() bool {
	return g[0] == '*'
}

type ProtoFieldType string

const (
	ProtoFieldTypeInt32   = ProtoFieldType("int32")
	ProtoFieldTypeInt64   = ProtoFieldType("int64")
	ProtoFieldTypeUint32  = ProtoFieldType("uint32")
	ProtoFieldTypeUint64  = ProtoFieldType("uint64")
	ProtoFieldTypeFloat32 = ProtoFieldType("float")
	ProtoFieldTypeFloat64 = ProtoFieldType("double")
	ProtoFieldTypeString  = ProtoFieldType("string")
	ProtoFieldTypeBool    = ProtoFieldType("bool")
	ProtoFieldTypeTime    = ProtoFieldType("google.protobuf.Timestamp")
)
