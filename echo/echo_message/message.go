package echo_message

import (
	"cmx/pkg/config"
	"cmx/pkg/config/definition"
	"cmx/pkg/util"
	"fmt"
	"strings"
)

type Message struct {
	father              *Message // father message
	Reference           string
	Name                string
	UpName              string
	MessageType         definition.MessageReferenceType // message type
	Fields              []definition.MessageField       // message field
	child               map[string]*Message             // child message
	Definition          config.Definition               // definition config
	writerMessageRecord map[string]bool                 // 已经写入的记录
	disableIncorporate  bool                            // 如果为真 message 类型子项不输出原始定义，而是引用定义
	incorporatePkName   string                          // 引用定义的包名
	quote               bool                            // 是否不包含外部message
}

type Option func(*Message)

func SetDisableIncorporate(disableIncorporate bool) func(*Message) {
	return func(m *Message) {
		m.disableIncorporate = disableIncorporate
	}
}

func SetIncorporatePkName(incorporatePkName string) func(*Message) {
	return func(m *Message) {
		m.incorporatePkName = incorporatePkName
	}
}

func NewMessage(df config.Definition, fn ...Option) *Message {
	msg := &Message{
		Fields:              make([]definition.MessageField, 0),
		child:               map[string]*Message{},
		Definition:          df,
		writerMessageRecord: map[string]bool{},
	}
	for _, f := range fn {
		f(msg)
	}
	return msg
}

func (m *Message) InitNode(reference string) {
	router := config.NewReferenceInformation(reference)
	var found bool
	if m.Fields, found = config.GetDefaultConfig().GetDefinition().GetMessageField(reference); !found {
		panic(fmt.Errorf("not found message %s", reference))
	}
	m.Reference = reference
	m.Name = router.Field
	m.UpName = util.ToCamelCasing(m.Name)
	m.MessageType = definition.RefTypeMessage
	for _, v := range m.Fields {
		if v.Ref.Type == definition.RefTypeMessage {
			child := NewMessage(
				m.Definition,
				SetDisableIncorporate(m.disableIncorporate),
			)
			child.InitNode(v.Ref.Ref)
			child.father = m
			m.child[v.ColumnName] = child
		}
	}
}

func (m *Message) WriterMessage(rw *strings.Builder) {
	if m.MessageType != definition.RefTypeMessage {
		panic(fmt.Errorf("not message type %s", m.MessageType))
	}
	if len(m.Fields) > 1 {
		rw.WriteString("/* \n")
		for _, v := range m.Fields {
			rw.WriteString(fmt.Sprintf(" %s $? \n",
				util.ToCamelCasing(v.ColumnName)))
		}
		rw.WriteString("*/ \n")

	}
	rw.WriteString(fmt.Sprintf("message %s {", m.UpName))
	// rw.WriteString("\n")
	rw.WriteString(fmt.Sprintf(`option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = {json_schema: {title: "%s" description: "%s"}};`, m.Name, m.Reference))
	rw.WriteString("\n")
	var idx int
	for _, v := range m.Fields {
		idx++
		// rw.WriteString(fmt.Sprintf("/* %s */\n", v.ColumnName))
		switch v.Ref.Type {
		case definition.RefTypeField:
			rw.WriteString(m.writerTypeField(v, idx))
		case definition.RefTypeEmpty:
			rw.WriteString(m.writerTypeSelf(v, idx))
		case definition.RefTypeMessage:
			rw.WriteString(m.writerTypeMessage(v, idx))
		case definition.RefTypeMessageField:
			rw.WriteString(m.writerTypeMessageField(v, idx))
		default:
			panic(fmt.Errorf("not support type %s", v.Ref.Type))
		}
	}
	rw.WriteString("}")
	rw.WriteString("\n")
}

// writer type message
func (m *Message) writerTypeMessage(v definition.MessageField, idx int) string {
	rw := strings.Builder{}

	if child, found := m.child[v.ColumnName]; found {
		if !m.writerMessageRecord[child.Reference] && !m.disableIncorporate {
			child.WriterMessage(&rw) // recursive writer message
			m.writerMessageRecord[child.Reference] = true
		}
		commentTag := strings.ReplaceAll(v.Comment, `"`, " ") + v.DetailComment + m.Definition.GetEnumComment(v.OneOf)
		rw.WriteString(m.Tag(v.ColumnName, v.Validator, v.OneOf, commentTag, v.Serializer))
		rw.WriteString("\n")
		messageUpName := child.UpName
		if m.disableIncorporate && m.incorporatePkName != "" {
			messageUpName = fmt.Sprintf("%s.%s", m.incorporatePkName, child.UpName)
			m.quote = true
		}

		rf := fmt.Sprintf("%s %s = %d", messageUpName, v.ColumnName, idx)
		if v.Array {
			rf = fmt.Sprintf("repeated %s %s = %d", messageUpName, v.ColumnName, idx)
		}
		if v.Optional {
			rf = fmt.Sprintf("optional %s %s = %d", messageUpName, v.ColumnName, idx)
		}
		rw.WriteString(rf)
		// rw.WriteString("\n")
		rw.WriteString("[(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field)={")
		// rw.WriteString("\n")
		rw.WriteString(fmt.Sprintf(`title: "%s"`, strings.ReplaceAll(v.Comment, `"`, " ")))
		// rw.WriteString("\n")
		rw.WriteString(" ")
		rw.WriteString(fmt.Sprintf(`description: "%s"`, v.DetailComment+m.Definition.GetEnumComment(v.OneOf)))
		// rw.WriteString("\n")
		rw.WriteString("}")
		m.fieldBehavior(&rw, v.Inhibit)
		rw.WriteString("]")
		rw.WriteString(";")
		rw.WriteString("\n")
	} else {
		panic(fmt.Errorf("not found child %s", v.ColumnName))
	}
	return rw.String()
}

// writer type self
func (m Message) writerTypeSelf(v definition.MessageField, idx int) string {
	rw := strings.Builder{}

	commentTag := strings.ReplaceAll(v.Comment, `"`, " ") + v.DetailComment + m.Definition.GetEnumComment(v.OneOf)
	rw.WriteString(m.Tag(v.ColumnName, v.Validator, v.OneOf, commentTag, v.Serializer))
	rw.WriteString("\n")

	rf := fmt.Sprintf("	%s %s = %d", v.Type, v.ColumnName, idx)
	if v.Array {
		rf = fmt.Sprintf("repeated %s %s = %d", v.Type, v.ColumnName, idx)
	}
	if v.Optional {
		rf = fmt.Sprintf("optional %s %s = %d", v.Type, v.ColumnName, idx)
	}
	rw.WriteString(rf)
	// rw.WriteString("\n")
	rw.WriteString("[(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field)={")
	rw.WriteString(fmt.Sprintf(`title: "%s"`, strings.ReplaceAll(v.Comment, `"`, " ")))
	rw.WriteString(" ")
	rw.WriteString(fmt.Sprintf(`description: "%s"`, v.DetailComment+m.Definition.GetEnumComment(v.OneOf)))
	// rw.WriteString("\n")
	rw.WriteString("}")
	m.fieldBehavior(&rw, v.Inhibit)
	rw.WriteString("];")
	rw.WriteString("\n")
	return rw.String()
}

// writer type field
func (m Message) writerTypeField(v definition.MessageField, idx int) string {
	rw := strings.Builder{}
	if tf, found := m.Definition.GetTableField(v.Ref.Ref); found {
		commentTag := fmt.Sprintf("%s %s %s", strings.ReplaceAll(tf.Comment, `"`, " "),
			v.DetailComment,
			m.Definition.GetEnumComment(tf.OneOf))
		rw.WriteString(m.Tag(v.ColumnName, tf.Validator, tf.OneOf, commentTag, v.Serializer))
		rw.WriteString("\n")
		rf := fmt.Sprintf("%s %s = %d", tf.Type, v.ColumnName, idx)
		if v.Array {
			rf = fmt.Sprintf("repeated %s %s = %d", tf.Type, v.ColumnName, idx)
		}
		if v.Optional {
			rf = fmt.Sprintf("optional %s %s = %d", tf.Type, v.ColumnName, idx)
		}
		rw.WriteString(rf)
		rw.WriteString("[(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field)={")
		// rw.WriteString("\n")
		rw.WriteString(fmt.Sprintf(`title: "%s"`, strings.ReplaceAll(tf.Comment, `"`, " ")))
		// rw.WriteString("\n")
		rw.WriteString(" ")
		rw.WriteString(fmt.Sprintf(`description: "%s"`, tf.DetailComment+
			m.Definition.GetEnumComment(tf.OneOf)))
		// rw.WriteString("\n")
		rw.WriteString("}")
		m.fieldBehavior(&rw, v.Inhibit)
		rw.WriteString("];")
		rw.WriteString("\n")
	} else {
		panic(fmt.Errorf("not found table field %s", v.Ref.Ref))
	}
	return rw.String()
}

// writer type message field
func (m Message) writerTypeMessageField(v definition.MessageField, idx int) string {
	rw := strings.Builder{}
	if len(v.Ref.Select) != 1 {
		panic(fmt.Errorf("not support select %v", v.Ref.Select))
	}
	slv := v.Ref.Select[0]
	if fields, found := m.Definition.GetMessageField(v.Ref.Ref); found {
		tf := definition.MessageField{}
		for _, v1 := range fields {
			if v1.ColumnName == slv {
				tf = v1
				break
			}
		}
		if slv != tf.ColumnName {
			panic(fmt.Errorf("not found select %s", slv))
		}
		commentTag := fmt.Sprintf("%s %s %s",
			strings.ReplaceAll(tf.Comment, `"`, " "),
			v.DetailComment,
			m.Definition.GetEnumComment(tf.OneOf))
		if v.Validator != "" {
			rw.WriteString(m.Tag(v.ColumnName, v.Validator, tf.OneOf, commentTag, v.Serializer))
		} else {
			rw.WriteString(m.Tag(v.ColumnName, tf.Validator, tf.OneOf, commentTag, v.Serializer))
		}
		rw.WriteString("\n")
		rf := fmt.Sprintf("%s %s = %d", tf.Type, v.ColumnName, idx)
		if v.Array {
			rf = fmt.Sprintf("repeated %s %s = %d", tf.Type, v.ColumnName, idx)
		}
		if v.Optional {
			rf = fmt.Sprintf("optional %s %s = %d", tf.Type, v.ColumnName, idx)
		}
		// if v.Ref.Ref == "model.commissions.commissions" && v.ColumnName == "status" {
		// 	fmt.Println("debug")
		// 	fmt.Println(rf)
		// }
		rw.WriteString(rf)
		rw.WriteString("[(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field)={")
		rw.WriteString(fmt.Sprintf(`title: "%s"`, strings.ReplaceAll(tf.Comment, `"`, " ")))
		rw.WriteString(" ")
		rw.WriteString(fmt.Sprintf(`description: "%s"`, tf.DetailComment+
			m.Definition.GetEnumComment(tf.OneOf)))
		rw.WriteString("}")
		m.fieldBehavior(&rw, v.Inhibit)
		rw.WriteString("];")
		rw.WriteString("\n")
	} else {
		panic(fmt.Errorf("not found table field %s", v.Ref.Ref))
	}
	return rw.String()
}

func (m Message) parserValidator(validator string, foo definition.FieldOneOf) (result string, found bool) {

	if validator == "" {
		return "omitempty", true
	}
	result = validator
	if strings.Contains(validator, "${oneof}") {
		enums := m.Definition.SelectEnumField(foo)
		if len(enums) == 0 {
			return validator, false
		}
		values := []string{}
		if foo.IsKey {
			for _, v := range enums {
				values = append(values, v.Key)
			}
		} else {
			for _, v := range enums {
				values = append(values, v.Value)
			}
		}
		result = strings.ReplaceAll(validator, "${oneof}", fmt.Sprintf("oneof=%s", strings.Join(values, " ")))
		found = true
	}

	return result, true
}

func (m Message) Tag(col, validator string, foo definition.FieldOneOf, comment, serializer string) string {
	rw := strings.Builder{}
	if validator, found := m.parserValidator(validator, foo); found {
		rw.WriteString(fmt.Sprintf(`// @gotags: binding:"%s" form:"%s" comment:"%s"`, validator, col, comment))
		if serializer != "" {
			// https://gorm.io/docs/serializer.html
			// gorm:"serializer:unixtime"
			rw.WriteString(fmt.Sprintf(` gorm:"serializer:%s"`, serializer))
		}
	} else {
		panic(fmt.Sprintf("build gotags error %s tag %s", col, validator))
	}
	return rw.String()
}

func (m Message) fieldBehavior(rw *strings.Builder, inhibit string) {
	if inhibit == "required" {
		rw.WriteString(",")
		rw.WriteString(`(google.api.field_behavior) = REQUIRED`)
	} else {
		rw.WriteString(",")
		rw.WriteString(`(google.api.field_behavior) = OPTIONAL`)
	}
}

func (m *Message) IsWriterSubMessage() (result bool) {
	if m.quote {
		return true
	}
	for _, v := range m.child {
		if v.IsWriterSubMessage() {
			return true
		}
	}
	return
}
