package echo_proto

import (
	config "cmx/v1/logic/aggregate/build_config"
	"cmx/v1/logic/aggregate/parse"
	message_model "cmx/v1/logic/model/message"
	_ "embed"
	"fmt"
	"strings"
	"text/template"
)

//go:embed echo_proto.tpl
var echo_proto_tpl string
var ECHO_PROTO_TPL *template.Template

func Generated(m *parse.Model) (b []byte, err error) {
	// 初始化模板

	// proto template func
	var templateFunc = template.FuncMap{
		"getProtoPkgName":         config.GetDefaultConfig().ProtoConfig.GetPkgName,
		"protoGoPkgName":          config.GetDefaultConfig().ProtoConfig.GetGoPkgName,
		"getImportPaths":          config.GetDefaultConfig().ProtoConfig.GetImportPaths,
		"sortColumn":              sortColumn,
		"sortColumnByFieldString": sortColumnByFieldString,
	}

	ECHO_PROTO_TPL = template.Must(template.New("echo_proto").
		Funcs(templateFunc).
		Parse(echo_proto_tpl))
	sb := strings.Builder{}

	err = ECHO_PROTO_TPL.Execute(&sb, m)
	if err != nil {
		return
	}

	return []byte(sb.String()), nil
}

// sort by idx echo model field
func sortColumn(m parse.Model) []parse.ModelField {
	return m.SortFields()
}

// int64 id = 1 [
//
//	(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {
//	  type: [ INTEGER ]
//	  description: "系统编号 "
//	},
//	(google.api.field_behavior) = OUTPUT_ONLY];
func sortColumnByFieldString(mdx parse.Model) string {
	result := []string{}
	fields := sortColumn(mdx)
	// get definition config
	groupDefinition := config.GetDefaultConfig().GetDefinition()
	for _, v := range fields {

		cfg, found := groupDefinition.GetTableField(fmt.Sprintf("%s.%s", mdx.Reference, v.FieldName))
		if found && cfg.Ref != "" {
			cfg, found = groupDefinition.GetTableField(cfg.Ref)
		}
		if found && cfg.Inhibit == message_model.Inhibit {
			continue // if field omit
		}
		selectType := string(v.ProtoSchema.Type)
		if found && cfg.Type != "" {
			selectType = cfg.Type
		}
		item := strings.Builder{}
		item.WriteString(fmt.Sprintf("  %s %s = %d", selectType, v.FieldName, v.Idx+1))
		openApiV2Field := strings.Builder{}
		openApiV2Field.WriteString("  (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = {")
		if strings.Contains(selectType, "int64") {
			openApiV2Field.WriteString(fmt.Sprintf("\n    type: [ %s ]", "INTEGER"))
		}
		if v.ProtoSchema.Comment != "" {
			selectComment := v.ProtoSchema.Comment
			if found && cfg.Comment != "" {
				selectComment = cfg.Comment
			}
			if getEnumComment(cfg.OneOf) != "" {
				selectComment = fmt.Sprintf("%s\\n%s", selectComment, getEnumComment(cfg.OneOf))
			}
			openApiV2Field.WriteString(fmt.Sprintf("\n    description: \"%s\"", selectComment))
		}
		openApiV2Field.WriteString("\n    },")
		item.WriteString(" [\n   ")
		item.WriteString(openApiV2Field.String())
		item.WriteString("\n   (google.api.field_behavior) = OUTPUT_ONLY\n  ];")
		result = append(result, item.String())
	}
	return strings.Join(result, "\n")
}

// getEnumComment
func getEnumComment(of message_model.FieldOneOf) string {
	return config.GetDefaultConfig().GetDefinition().GetEnumComment(of)
}
