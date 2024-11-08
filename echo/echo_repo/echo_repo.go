package echo_repo

import (
	"cmp"
	"cmx/pkg/config"
	"cmx/pkg/config/definition"
	"cmx/pkg/parse"
	"cmx/pkg/util"
	_ "embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"

	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	"gopkg.in/yaml.v3"
)

//go:embed echo.tpl
var echo_tpl string
var ECHO_TPL *template.Template

func Generated(referenceString string) (b []byte, err error) {
	sdf, found := config.GetDefaultConfig().GetDefinition().GetStatementField(referenceString)
	if !found {
		panic(fmt.Errorf("not found %s", referenceString))
	}

	mod := parse.ParseSqlToModel(sdf.Statement, parse.Options{})
	mod.Reference = config.NewReferenceInformation(referenceString).Route

	// template func
	var templateFunc = template.FuncMap{
		"getPkgName":         config.GetDefaultConfig().RepoConfig.GetPkgName,
		"getImportPaths":     config.GetDefaultConfig().RepoConfig.GetImportPaths,
		"getDefaultRefs":     config.GetDefaultConfig().RepoConfig.GetDefaultRefs,
		"newModel":           newModel,
		"model":              model,
		"sortColumn":         sortColumn,
		"genreField":         genreField,
		"getOneofDefinition": getOneofDefinition,
		// "getModelToProto":    getModelToProto,
	}

	ECHO_TPL = template.Must(template.New("echo_repo").
		Funcs(templateFunc).
		Parse(echo_tpl))

	sb := strings.Builder{}

	err = ECHO_TPL.Execute(&sb, mod)
	if err != nil {
		return
	}

	// return []byte(sb.String()), nil
	return format.Source([]byte(sb.String()))
}

// sort by idx echo model field
func sortColumn(m *parse.Model) []parse.ModelField {
	return m.SortFields()
}

func newModel(name string, ref bool) (s string) {
	if ref {
		return fmt.Sprintf("*%s%s{}", config.GetDefaultConfig().RepoConfig.ModelNameTpl, name)
	} else {
		return fmt.Sprintf("%s%s{}", config.GetDefaultConfig().RepoConfig.ModelNameTpl, name)
	}
}

func model(name string, ref bool) (s string) {
	if ref {
		return fmt.Sprintf("*%s%s", config.GetDefaultConfig().RepoConfig.ModelNameTpl, name)
	} else {
		return fmt.Sprintf("%s%s", config.GetDefaultConfig().RepoConfig.ModelNameTpl, name)
	}
}

func proto(name string, ref bool) (s string) {
	if ref {
		return fmt.Sprintf("*%s%s", config.GetDefaultConfig().RepoConfig.ProtoNameTpl, name)
	} else {
		return fmt.Sprintf("%s%s", config.GetDefaultConfig().RepoConfig.ProtoNameTpl, name)
	}
}

func genreField(goType parse.GoFieldType) string {
	result := ""
	switch goType {
	case parse.GoFieldTypeSoftDelete:
		return "Field"
	case parse.GoFieldTypeJson:
		return "Bytes"
	case parse.GoFieldTypeTime:
		s := strings.TrimPrefix(string(goType), "*")
		result = util.ToCamelCasing(strings.TrimPrefix(s, "time."))
	default:
		result = util.ToCamelCasing(strings.TrimPrefix(string(goType), "*"))
	}
	return result
}

func getOneofDefinition(m *parse.Model) string {
	sb := strings.Builder{}
	// return m.Reference
	groupDefinition := config.GetDefaultConfig().GetDefinition()
	tp, found := groupDefinition.GetTable("/" + m.Reference)
	if !found {
		panic(fmt.Errorf("not found %s", m.Reference))
	}
	dm, ok := tp.Definition[m.TableName]
	if !ok {
		panic(fmt.Errorf("not found \n%s", util.MustSucc(yaml.Marshal(tp.Definition))))
	}

	result := lo.Filter(dm, func(i definition.Field, idx int) bool {
		return i.OneOf.Ref != ""
	})

	slices.SortFunc(result, func(i definition.Field, j definition.Field) int {
		return cmp.Compare(i.ColumnName, j.ColumnName)
	})

	for _, v := range result {
		enum := groupDefinition.SelectEnumField(v.OneOf)
		sb.WriteString(fmt.Sprintf("// %s %s\n", util.ToCamelCasing(v.ColumnName), v.OneOf.Ref))
		sb.WriteString(fmt.Sprintf("// %s \n", strings.ReplaceAll(groupDefinition.GetEnumComment(v.OneOf), "\\n", " ")))
		upCol := util.ToCamelCasing(v.ColumnName)
		// to int
		for _, ev := range enum {
			upEvKey := util.ToCamelCasing(ev.Key)
			sb.WriteString(fmt.Sprintf("const %s_%s_%s = %s\n", m.UpTableName, upCol, upEvKey, ev.Value))
		}

		for _, ev := range enum {
			upEvKey := util.ToCamelCasing(ev.Key)
			sb.WriteString(fmt.Sprintf(`const Str_%s_%s_%s = "%s"`, m.UpTableName, upCol, upEvKey, ev.Value))
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func getModelToProto(repoName, modelName string, m *parse.Model) string {
	if !config.GetDefaultConfig().RepoConfig.EnableModelToProto {
		return "// not enable model to proto"
	}
	protoName := proto(m.UpTableName, false)
	sb := strings.Builder{}
	_ = sb
	sb.WriteString("func ")
	sb.WriteString(fmt.Sprintf("(repo %s) ", repoName)) // model
	sb.WriteString("ToProto ")
	sb.WriteString(fmt.Sprintf("(m %s) ", modelName)) // model
	sb.WriteString("*" + protoName)                   // return proto
	sb.WriteString("{")
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("proto:=&%s{}", protoName))
	sb.WriteString("\n")
	groupDefinition := config.GetDefaultConfig().GetDefinition()
	dm, found := groupDefinition.GetMessages("/" + m.Reference)
	if !found {
		panic(fmt.Errorf("message not found %s", m.Reference))
	}
	messageField, ok := dm.Definition[m.TableName]
	if !ok {
		panic(fmt.Errorf("message filed not found %s", m.TableName))
	}
	table, found := groupDefinition.GetTable("/" + m.Reference)
	if !found {
		panic(fmt.Errorf("table not found %s", m.Reference))
	}
	tableField, ok := table.Definition[m.TableName]
	if !ok {
		panic(fmt.Errorf("table filed not found %s", m.TableName))
	}
	for _, mf := range messageField {
		if mf.Inhibit == string(definition.Inhibit) {
			continue // skip field
		}
		for _, tf := range tableField {
			if tf.ColumnName != mf.ColumnName {
				continue
			}
			col := util.ToCamelCasing(tf.ColumnName)

			switch tf.Type {
			case "time.Time":
				tmp := ""
				switch mf.Type {
				case "int64", "uint64":
					tmp = fmt.Sprintf(`if !m.%s.IsZero() {proto.%s = m.%s.Unix()}`,
						col, col, col)
				case "string":
					tmp = fmt.Sprintf(`if !m.%s.IsZero() {proto.%s = m.%s.Format(time.RFC3339)}`,
						col, col, col)
				default:
					panic(fmt.Errorf("not support model: %s field: %s time to %s",
						m.Reference, mf.ColumnName, mf.Type))
				}
				sb.WriteString(tmp)
			case "*time.Time":
				tmp := ""
				switch mf.Type {
				case "int64", "uint64":
					tmp = fmt.Sprintf(`if m.%s != nil &&  !m.%s.IsZero() {proto.%s = m.%s.Unix()}`,
						col, col, col, col)
				case "string":
					tmp = fmt.Sprintf(`if m.%s != nil && !m.%s.IsZero() {proto.%s = m.%s.Format(time.RFC3339)}`,
						col, col, col, col)
				default:
					panic(fmt.Errorf("not support model: %s field: %s time to %s",
						m.Reference, mf.ColumnName, mf.Type))
				}
				sb.WriteString(tmp)
			default:
				tfType := strings.ReplaceAll(tf.Type, "*", "")
				if parse.GoFieldTypeSoftDelete == parse.GoFieldType(tfType) {
					tfType = string(parse.GoFieldTypeInt64)
				}
				sb.WriteString(fmt.Sprintf("proto.%s = ", col))
				sb.WriteString(fmt.Sprintf("cast.To%s(m.%s)", util.ToCamelCasing(tfType), col))
			}
			sb.WriteString("\n")
		}
	}
	sb.WriteString("return proto")
	sb.WriteString("}")
	sb.WriteString("\n")
	sb.WriteString("func ")
	sb.WriteString(fmt.Sprintf("(repo %s) ", repoName)) // model
	sb.WriteString("ArrayToProto ")
	sb.WriteString(fmt.Sprintf("(m ...%s) ", modelName)) // model
	sb.WriteString("[]*" + protoName)                    // return proto
	sb.WriteString("{")
	sb.WriteString("\n")
	sb.WriteString("var proto []*" + protoName)
	sb.WriteString("\n")
	sb.WriteString(`for _, v := range m {
	proto = append(proto, repo.ToProto(v))
	}`)
	sb.WriteString("\n")
	sb.WriteString("return proto")
	sb.WriteString("}")
	return sb.String()
}
