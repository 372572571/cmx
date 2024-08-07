package echo_type

import (
	"cmx/pkg/config"
	"cmx/pkg/util"
	_ "embed"
	"fmt"
	"go/format"
	"strings"
	"text/template"
)

//go:embed echo.tpl
var echo_tpl string
var ECHO_TPL *template.Template

func Generated(cfg config.Config) ([]byte, error) {
	result := ""
	definition := cfg.GetDefinition()
	routes := definition.GetEnumRoutes()
	for _, route := range routes {
		gen, err := generatedItem(route)
		if err != nil {
			panic(err)
		}
		result = fmt.Sprintf("%s \n%s", result, gen)
	}

	// template func
	var templateFunc = template.FuncMap{
		"getPkgName":     config.GetDefaultConfig().TypeConfig.GetGoPkgName,
		"getImportPaths": config.GetDefaultConfig().TypeConfig.GetImportPaths,
		"getDefaultRefs": config.GetDefaultConfig().TypeConfig.GetDefaultRefs,
	}

	ECHO_TPL = template.Must(template.New("echo_type").
		Funcs(templateFunc).
		Parse(echo_tpl))
	sb := strings.Builder{}

	err := ECHO_TPL.Execute(&sb, result)
	if err != nil {
		panic(err)
	}
	// return []byte(sb.String()), nil
	return format.Source([]byte(sb.String()))
}

func generatedItem(route string) (string, error) {
	definition := config.GetDefaultConfig().GetDefinition()
	sb := strings.Builder{}
	if content, found := definition.GetEnum(route); found {
		for k, v := range content.Definition {
			name := fmt.Sprintf("%s_%s", route, k)
			low := util.FirstLowerCamelCasing(name)
			up := util.ToCamelCasing(name)
			sb.WriteString(fmt.Sprintf("type %s struct{}", low))
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("var _%s = %s{}", low, low))
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("type E%s int32", up))
			sb.WriteString("\n")
			for _, vv := range v {
				if vv.Desc != "" {
					sb.WriteString(fmt.Sprintf(`// %s`, vv.Desc))
				} else {
					sb.WriteString(fmt.Sprintf(`// %s`, vv.Zh))
				}
				sb.WriteString("\n")
				sb.WriteString(fmt.Sprintf("const %s_%s E%s = %s", up, util.ToCamelCasing(vv.Key), up, vv.Value))
				sb.WriteString("\n")
			}
			sb.WriteString("\n")
			// desc
			// sb.WriteString("/*")
			sb.WriteString(fmt.Sprintf("func (e E%s) desc() map[int32]string {", up))
			sb.WriteString("\n")
			sb.WriteString("return map[int32]string{")
			sb.WriteString("\n")
			for _, vv := range v {
				desc := vv.Desc
				if vv.Desc == "" {
					desc = vv.Zh
				}
				sb.WriteString(fmt.Sprintf("%s: \"%s\",", vv.Value, desc))
				sb.WriteString("\n")
			}
			sb.WriteString("}")
			sb.WriteString("\n")
			sb.WriteString("}")
			sb.WriteString("\n")
			// sb.WriteString("*/")
			// ekeys
			sb.WriteString(fmt.Sprintf("func (e E%s) keys() map[int32]string {", up))
			sb.WriteString("\n")
			sb.WriteString("return map[int32]string{")
			sb.WriteString("\n")
			for _, vv := range v {
				sb.WriteString(fmt.Sprintf("%s: \"%s\",", vv.Value, vv.Key))
				sb.WriteString("\n")
			}
			sb.WriteString("}")
			sb.WriteString("\n")
			sb.WriteString("}")
			sb.WriteString("\n")

			sb.WriteString(fmt.Sprintf("func (e E%s) ZH() string {", up))
			sb.WriteString("return e.desc()[e.I3()]")
			sb.WriteString("}")
			sb.WriteString("\n")

			sb.WriteString(fmt.Sprintf("func (e E%s) I3() int32 {", up))
			sb.WriteString("return int32(e)")
			sb.WriteString("}")
			sb.WriteString("\n")

			sb.WriteString(fmt.Sprintf("func (e E%s) UI3() uint32 {", up))
			sb.WriteString("return uint32(e)")
			sb.WriteString("}")
			sb.WriteString("\n")

			sb.WriteString(fmt.Sprintf("func (e E%s) I() int {", up))
			sb.WriteString("return int(e)")
			sb.WriteString("}")
			sb.WriteString("\n")

			sb.WriteString(fmt.Sprintf("func (e E%s) S() string {", up))
			sb.WriteString("return strconv.Itoa(e.I())")
			sb.WriteString("}")
			sb.WriteString("\n")

			sb.WriteString(fmt.Sprintf("func (e E%s) K() string {", up))
			sb.WriteString("return e.keys()[e.I3()]")
			sb.WriteString("}")
			sb.WriteString("\n")

			sb.WriteString(fmt.Sprintf("var %s func() %s = func() %s {", up, low, low))
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("return _%s", low))
			sb.WriteString("\n")
			sb.WriteString("}")
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("func (e %s) nameToValueMap() map[string]int32 {", low))
			sb.WriteString("\n")
			sb.WriteString("return map[string]int32{")
			sb.WriteString("\n")
			for _, vv := range v {
				sb.WriteString(fmt.Sprintf("\"%s\": %s,", vv.Key, vv.Value))
				sb.WriteString("\n")
			}
			sb.WriteString("}")
			sb.WriteString("\n")
			sb.WriteString("}")
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf("func (e %s) valueToNameMap() map[int32]string {", low))
			sb.WriteString("\n")
			sb.WriteString("return map[int32]string{")
			sb.WriteString("\n")
			for _, vv := range v {
				sb.WriteString(fmt.Sprintf("%s: \"%s\",", vv.Value, vv.Key))
				sb.WriteString("\n")
			}
			sb.WriteString("}")
			sb.WriteString("\n")
			sb.WriteString("}")
			sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf(`
			func (e %s) Value(key string) int32 {
				return e.nameToValueMap()[key]
			}`, low))
			sb.WriteString("\n")
			// sb.WriteString(fmt.Sprintf(`
			// func (e %s) Desc(value int32) string {
			// 	return e.desc()[value]
			// }`, low))
			// sb.WriteString("\n")
			sb.WriteString(fmt.Sprintf(`
			func (e %s) Key(value int32) string {
				return e.valueToNameMap()[value]
			}`, low))
			sb.WriteString("\n")
			for _, vv := range v {
				if vv.Desc != "" {
					sb.WriteString(fmt.Sprintf(`// %s`, vv.Desc))
				}
				if vv.Zh != "" {
					sb.WriteString(fmt.Sprintf(`// %s`, vv.Zh))
				}
				sb.WriteString(fmt.Sprintf(`
				func (e %s) %sKey() string {
					return "%s"
				}`, low, util.ToCamelCasing(vv.Key), vv.Key))
				sb.WriteString("\n")
			}
			for _, vv := range v {
				if vv.Desc != "" {
					sb.WriteString(fmt.Sprintf(`// %s`, vv.Desc))
				}
				if vv.Zh != "" {
					sb.WriteString(fmt.Sprintf(`// %s`, vv.Zh))
				}
				sb.WriteString(fmt.Sprintf(`
				func (e %s) %sValue() int32 {
					return %s
				}`, low, util.ToCamelCasing(vv.Key), vv.Value))
				sb.WriteString("\n")
				sb.WriteString(fmt.Sprintf(`
				func (e %s) %sValString() string {
					return "%s"
				}`, low, util.ToCamelCasing(vv.Key), vv.Value))
				sb.WriteString("\n")
			}
		}
	}
	return sb.String(), nil
}

func createEtype(up, low string) {

}
