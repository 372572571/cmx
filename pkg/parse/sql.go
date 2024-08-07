package parse

import (
	"cmx/pkg/config"
	"cmx/pkg/util"
	_ "embed"
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"vitess.io/vitess/go/vt/sqlparser"
)

type Options struct {
	EnableProtoBigIntToString bool
}

func ParseSqlToModel(c_sql string, option Options) *Model {
	parser, err := sqlparser.New(sqlparser.Options{})
	if err != nil {
		panic(err)
	}
	stmt, err := parser.Parse(c_sql)
	if err != nil {
		// Do something with the err
		panic(err)
	}

	// fmt.Println(reflect.ValueOf(stmt).String())
	var mod = NewModel()
	switch stmt := stmt.(type) {
	case *sqlparser.CreateTable:
		// get table name
		mod.TableName = stmt.Table.Name.String()
		mod.UpTableName = util.ToCamelCasing(mod.TableName)
		mod.FirstLowerTableName = util.FirstLowerCamelCasing(mod.UpTableName)

		comment := getOptions(stmt.TableSpec.Options, "comment")

		mod.Comment = comment

		for _, information := range stmt.TableSpec.Indexes {
			key_mod := NewModelKey()

			// get key type
			if information.Info.Type == sqlparser.IndexTypePrimary {
				key_mod.KeyType = KeyTypePrimary
			} else if information.Info.Type == sqlparser.IndexTypeUnique {
				key_mod.KeyType = KeyTypeUnique
			} else {
				key_mod.KeyType = KeyTypeIndex
			}

			// get key name
			key_mod.Key = information.Info.Name.CompliantName()

			// marker complex key
			if len(information.Columns) > 1 {
				key_mod.IsComplex = true
			}

			// //  get key field
			for _, column := range information.Columns {
				key_mod.Field = append(key_mod.Field, column.Column.String())
			}

			mod.AddKey(key_mod.KeyType, key_mod)
		}
		// get field information
		for k, column := range stmt.TableSpec.Columns {

			mod_field := NewModelField()
			mod_field.Idx = k
			mod_field.FieldName = column.Name.String()
			mod_field.UpFieldName = util.ToCamelCasing(mod_field.FieldName)
			mod_field.SqlSchema = NewSqlFieldSchema()

			mod_field.SqlSchema.Type = SqlFieldType(column.Type.Type)

			mod_field.SqlSchema.IsUnsigned = bool(column.Type.Unsigned)

			if column.Type.Length != nil {
				mod_field.SqlSchema.Lenght = cast.ToInt(column.Type.Length)
			}

			if column.Type.Scale != nil {
				mod_field.SqlSchema.Scale = cast.ToInt(column.Type.Scale)
			}

			if column.Type.Options != nil && column.Type.Options.Comment != nil {
				mod_field.SqlSchema.Comment = string(column.Type.Options.Comment.Val)
			}

			if column.Type.Options != nil {
				if column.Type.Options.Null != nil {
					mod_field.SqlSchema.IsNotNull = new(bool)
					if !*column.Type.Options.Null {
						mod_field.SqlSchema.IsNotNull = new(bool)
						*mod_field.SqlSchema.IsNotNull = true
					} else {
						mod_field.SqlSchema.IsNotNull = new(bool)
						*mod_field.SqlSchema.IsNotNull = false
					}
				}

				mod_field.SqlSchema.IsAutoincrement = column.Type.Options.Autoincrement

				if len(mod.SearchFieldKey(mod_field.FieldName)[KeyTypePrimary]) == 0 &&
					column.Type.Options.Default != nil {
					// sql field default
					if in, ok := column.Type.Options.Default.(*sqlparser.Literal); ok {
						// 如果字段是主键则忽略
						switch in.Type {
						case sqlparser.IntVal,
							sqlparser.DecimalVal,
							sqlparser.FloatVal:
							mod_field.SqlSchema.DefaultValue = in.Val
						case sqlparser.StrVal:
							// mod_field.SqlSchema.DefaultValue = fmt.Sprintf("'%s'", in.Val)
							mod_field.SqlSchema.DefaultValue = in.Val
						default:
							mod_field.SqlSchema.DefaultValue = in.Val
							fmt.Printf("column: %s  type: %d , value: %s \n", column.Name.String(), in.Type, in.Val)
						}
						// fmt.Printf("%s.%s %+v \n", stmt.Table.Name.String(), column.Name.String(), in.Val)
					}
				} else {
					if column.Type.Options.Default != nil {
						fmt.Printf("table %s field %s key %v ignore primary key default value \n", mod.TableName, mod_field.FieldName, mod.SearchFieldKey(mod_field.FieldName)[KeyTypePrimary][0])
					}
				}

			}

			mod_field.GoSchema = sqlCreateGo(mod, mod_field, option)
			mod_field.ProtoSchema = sqlCreateProto(mod, mod_field, option)

			mod.AddField(mod_field.FieldName, mod_field)

		}

	default:
		panic("not support")
	}

	return mod
}

func sqlCreateGo(m *Model, field *ModelField, option Options) *GoFieldSchema {
	_ = option
	go_field := NewGoFieldSchema()
	_ = go_field

	switch field.SqlSchema.Type {
	case SqlFieldTypeBigInt:
		if field.SqlSchema.IsUnsigned {
			go_field.Type = GoFieldTypeUint64
		} else {
			go_field.Type = GoFieldTypeInt64
		}
	case SqlFieldTypeChar,
		SqlFieldTypeVarchar,
		SqlFieldTypeText:
		go_field.Type = GoFieldTypeString
	case SqlFieldTypeDateTime,
		SqlFieldTypeDate,
		SqlFieldTypeTimeStamp:
		go_field.Type = GoFieldTypeTime
	case SqlFieldTypeDecimal,
		SqlFieldTypeDouble,
		SqlFieldTypeFloat:
		go_field.Type = GoFieldTypeString
	case SqlFieldTypeInt,
		SqlFieldTypeTinyInt,
		SqlFieldTypeSmallInt,
		SqlFieldTypeMediumInt:
		go_field.Type = GoFieldTypeInt
	case SqlFieldTypeJson:
		go_field.Type = GoFieldTypeJson
	default:
		go_field.Type = GoFieldTypeString
		fmt.Printf("not support %v \n", field.SqlSchema.Type)
		// panic(fmt.Errorf("not support %v", field.SqlSchema.Type))
	}

	go_field.Tag = buildGormTag(m, field)

	// nullable type parse to pointer type
	if config.GetDefaultConfig().EnableGoNullPoint {
		if field.SqlSchema.IsNotNull != nil && !*field.SqlSchema.IsNotNull {
			go_field.IsNullable = true
			go_field.Type = parseNullableType(*go_field)
		}
	}

	// gorm soft delete
	if config.GetDefaultConfig().EnableGormSoftDelete {
		if field.UpFieldName == "DeletedAt" {
			go_field.Type = GoFieldTypeSoftDelete
		}
	}

	go_field.Comment = field.SqlSchema.Comment
	go_field.Length = field.SqlSchema.Lenght
	return go_field
}

func sqlCreateProto(m *Model, field *ModelField, option Options) *ProtoFieldSchema {
	_ = m
	proto_field := NewProtoFieldSchema()

	switch field.SqlSchema.Type {
	case SqlFieldTypeBigInt:
		if option.EnableProtoBigIntToString {
			proto_field.Type = ProtoFieldTypeString
		} else {
			if field.SqlSchema.IsUnsigned {
				proto_field.Type = ProtoFieldTypeUint64
			} else {
				proto_field.Type = ProtoFieldTypeInt64
			}
		}
	case SqlFieldTypeChar,
		SqlFieldTypeVarchar,
		SqlFieldTypeText:
		proto_field.Type = ProtoFieldTypeString
	case SqlFieldTypeDateTime,
		SqlFieldTypeDate,
		SqlFieldTypeTimeStamp:
		if option.EnableProtoBigIntToString {
			proto_field.Type = ProtoFieldTypeString
		} else {
			proto_field.Type = ProtoFieldTypeInt64
		}
	case SqlFieldTypeDecimal,
		SqlFieldTypeDouble,
		SqlFieldTypeFloat:
		proto_field.Type = ProtoFieldTypeString
	case SqlFieldTypeInt,
		SqlFieldTypeTinyInt,
		SqlFieldTypeSmallInt,
		SqlFieldTypeMediumInt:
		proto_field.Type = ProtoFieldTypeInt32
	case SqlFieldTypeJson:
		proto_field.Type = ProtoFieldTypeString
	default:
		proto_field.Type = ProtoFieldTypeString
		fmt.Printf("not support %v \n", field.SqlSchema.Type)
	}

	// nullable type parse to pointer type
	if field.SqlSchema.IsNotNull != nil {
		proto_field.IsNullable = !*field.SqlSchema.IsNotNull
	}

	proto_field.Comment = field.SqlSchema.Comment
	proto_field.Length = field.SqlSchema.Lenght

	return proto_field
}

func parseNullableType(schema GoFieldSchema) GoFieldType {
	if schema.IsNullable {
		return GoFieldType(fmt.Sprintf("*%s", schema.Type))
	}
	return schema.Type
}

func buildIndexKeyTag(m *Model,
	field *ModelField,
	sb *strings.Builder) {

	// primary key tag
	// if field.SqlSchema.KeyTyep == KeyTypePrimary {
	// 	if field.SqlSchema.IsAutoincrement {
	// 		sb.WriteString(";primaryKey;autoIncrement")
	// 	} else {
	// 		fmt.Println("primaryKey")
	// 		sb.WriteString(";primaryKey")
	// 	}
	// }

	// index unique tag
	result := m.SearchFieldKey(field.FieldName)
	for k, v := range result {
		switch k {
		case KeyTypePrimary:
			if field.SqlSchema.IsAutoincrement {
				sb.WriteString(";primaryKey;autoIncrement")
			} else {
				sb.WriteString(";primaryKey")
			}
		case KeyTypeUnique:
			for _, v1 := range v {
				sb.WriteString(fmt.Sprintf(";uniqueIndex:%s", v1.Key))
			}
		case KeyTypeIndex:
			for _, v1 := range v {
				sb.WriteString(fmt.Sprintf(";index:%s", v1.Key))
			}
		}
	}
}

// buildGormTag build gorm tag
func buildGormTag(m *Model, field *ModelField) string {
	if config.GetDefaultConfig().EnableGormTag {
		sb := strings.Builder{}
		sb.WriteString(fmt.Sprintf("gorm:\"column:%s", field.FieldName))

		ts := string(field.SqlSchema.Type)
		if field.SqlSchema.Lenght != 0 {
			if field.SqlSchema.Scale == 0 {
				ts = fmt.Sprintf("%s(%d)", ts, field.SqlSchema.Lenght)
			} else {
				ts = fmt.Sprintf("%s(%d,%d)", ts, field.SqlSchema.Lenght, field.SqlSchema.Scale)
			}
		}
		if field.SqlSchema.IsUnsigned {
			ts = fmt.Sprintf("%s unsigned", ts)
		}
		sb.WriteString(fmt.Sprintf(";type:%s", ts))

		buildIndexKeyTag(m, field, &sb)

		// not null
		if field.SqlSchema.IsNotNull != nil && *field.SqlSchema.IsNotNull {
			sb.WriteString(";not null")
		}

		// default value
		if field.SqlSchema.DefaultValue != "" {
			sb.WriteString(fmt.Sprintf(";default:%s", field.SqlSchema.DefaultValue))
		}

		// comment
		if field.SqlSchema.Comment != "" {
			sb.WriteString(fmt.Sprintf(";comment:%s", field.SqlSchema.Comment))
		}

		// end gorm tag
		sb.WriteString(`" `)

		// json tag
		sb.WriteString(`json:"` + field.FieldName + `,omitempty" `)

		// yaml tag
		sb.WriteString(`yaml:"` + field.FieldName + `,omitempty"`)

		return sb.String()
	}
	return ""
}

func getOptions(options sqlparser.TableOptions, key string) string {
	uk := strings.ToUpper(key)
	for _, option := range options {
		if option.Name == uk {
			return option.String
		}
	}
	return ""
}
